package controller

import (
	"github.com/mitene/terrafire/internal/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestController(t *testing.T) {
	type call struct {
		action    string
		project   string
		workspace string
	}
	type response struct {
		err       error
		tryCancel bool
	}

	tests := []struct {
		messages  []*api.GetActionResponse
		responses map[call]response
		wantCalls []call
	}{
		{
			messages: []*api.GetActionResponse{
				{
					Type: api.GetActionResponse_NONE,
				},
				{
					Type:      api.GetActionResponse_SUBMIT,
					Project:   "pj1",
					Workspace: "ws1",
				},
				{
					Type:      api.GetActionResponse_APPROVE,
					Project:   "pj1",
					Workspace: "ws1",
				},
			},
			responses: map[call]response{
				{action: "Apply", project: "pj1", workspace: "ws1"}: {
					tryCancel: true,
				},
			},
			wantCalls: []call{
				{action: "Plan", project: "pj1", workspace: "ws1"},
				{action: "Apply", project: "pj1", workspace: "ws1"},
				{action: "Cancel", project: "pj1", workspace: "ws1"},
			},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			client := api.NewSchedulerClientMock()
			executor := NewExecutorMock()
			ctrl := New(client, executor, 1)
			var calls []call

			wg := sync.WaitGroup{}
			wg.Add(len(tt.messages))

			// mock GetAction and GetActionControl
			ms := tt.messages
			getActionCall := client.On("GetAction", mock.Anything, mock.Anything, mock.Anything)
			getActionCall.Run(func(args mock.Arguments) {
				if len(ms) == 0 {
					wg.Wait()
					getActionCall.Return(&api.GetActionResponse{Type: api.GetActionResponse_NONE}, nil)
					return
				}
				m := ms[0]
				ms = ms[1:]
				if m.GetType() == api.GetActionResponse_NONE {
					wg.Done()
				}
				getActionCall.Return(m, nil)
			})
			msCtrl := make(chan *api.GetActionControlResponse)
			getActionControlCall := client.On("GetActionControl", mock.Anything, mock.Anything, mock.Anything)
			getActionControlCall.Run(func(args mock.Arguments) {
				m := <-msCtrl
				getActionControlCall.Return(m, nil)
			})

			// mock Plan and Apply
			newMockProcess := func(cc call) *mockProcess {
				m := &mockProcess{}
				ch := make(chan struct{})
				m.On("cancel").Return(nil).Run(func(mock.Arguments) {
					calls = append(calls, call{action: "Cancel", project: cc.project, workspace: cc.workspace})
					wg.Done()
					close(ch)
				})
				m.On("wait").Return(nil).Run(func(mock.Arguments) {
					if resp, ok := tt.responses[cc]; ok {
						if resp.tryCancel {
							wg.Add(1)
							msCtrl <- &api.GetActionControlResponse{
								Type:      api.GetActionControlResponse_CANCEL,
								Project:   cc.project,
								Workspace: cc.workspace,
							}
							<-ch
						}
					}
					wg.Done()
				})
				return m
			}
			mockAction := func(meth string) {
				c := executor.On(meth, mock.Anything, mock.Anything)
				c.Run(func(args mock.Arguments) {
					project, workspace := args.String(0), args.String(1)
					cc := call{action: meth, project: project, workspace: workspace}

					var err error
					if resp, ok := tt.responses[cc]; ok {
						err = resp.err
					}

					calls = append(calls, cc)
					c.Return(newMockProcess(cc), err)
				})
			}
			mockAction("Plan")
			mockAction("Apply")

			// start/stop controller
			done := make(chan struct{})
			go func() {
				assert.NoError(t, ctrl.Start())
				close(done)
			}()
			go func() {
				wg.Wait()
				assert.NoError(t, ctrl.Stop())
			}()
			select {
			case <-done:
				break
			case <-time.After(10 * time.Second):
				assert.FailNow(t, "controller did not stop within 10 seconds")
			}

			assert.Equal(t, tt.wantCalls, calls)
		})
	}
}

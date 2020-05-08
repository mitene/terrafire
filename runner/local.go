package runner

import (
	"bytes"
	"github.com/mitene/terrafire/core"
	"hash/maphash"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

type LocalRunner struct {
	chs []chan *payload
	tf  Terraform
	dir string
}

func NewLocalRunner(workerNum int, tf Terraform) *LocalRunner {
	if workerNum < 1 {
		workerNum = 1
	}

	chs := make([]chan *payload, workerNum)
	for i := range chs {
		chs[i] = make(chan *payload, 4)
	}

	return &LocalRunner{
		chs: chs,
		tf:  tf,
		dir: "", // initialized when started
	}
}

func (r *LocalRunner) Plan(project string, workspace *core.Workspace) error {
	i := r.selectChannel(project, workspace.Name)
	r.chs[i] <- &payload{
		action:    actionPlan,
		project:   project,
		workspace: workspace,
	}
	return nil
}

func (r *LocalRunner) Apply(project string, workspace *core.Workspace) error {
	i := r.selectChannel(project, workspace.Name)
	r.chs[i] <- &payload{
		action:    actionApply,
		project:   project,
		workspace: workspace,
	}
	return nil
}

func (r *LocalRunner) selectChannel(project string, workspace string) int {
	h := maphash.Hash{}
	_, err := h.WriteString(project + ":" + workspace)
	if err != nil {
		return 0
	}
	return int(h.Sum64() % uint64(len(r.chs)))
}

func (r *LocalRunner) Start(h core.ServiceProvider) error {
	dir, err := ioutil.TempDir("", "terrafire-local-runner")
	if err != nil {
		return err
	}
	r.dir = dir

	for _, ch := range r.chs {
		go r.startWorker(ch, h)
	}

	return nil
}

func (r *LocalRunner) Clean() error {
	if r.dir != "" {
		return os.RemoveAll(r.dir)
	}
	return nil
}

func (r *LocalRunner) startWorker(ch chan *payload, h core.ServiceProvider) {
	for payload := range ch {
		switch payload.action {
		case actionPlan:
			r.doPlan(payload, h)
		case actionApply:
			r.doApply(payload, h)
		default:
			log.Printf("invalid action type: %d\n", payload.action)
		}
	}
}

func (r *LocalRunner) doPlan(payload *payload, h core.ServiceProvider) {
	dir, err := r.initWorkDir(payload)
	if err != nil {
		logError(h.UpdateJobStatusPlanFailed(payload.project, payload.workspace.Name, err))
		return
	}

	logError(h.UpdateJobStatusPlanInProgress(payload.project, payload.workspace.Name))

	output := bytes.NewBuffer(nil)
	finish := false
	go func() {
		for !finish {
			logError(h.SavePlanLog(payload.project, payload.workspace.Name, output.String()))
			time.Sleep(5 * time.Second)
		}
		logError(h.SavePlanLog(payload.project, payload.workspace.Name, output.String()))
	}()
	defer func() {
		finish = true
	}()

	result, err := r.tf.Plan(dir, payload.workspace, output)
	if err != nil {
		logError(h.UpdateJobStatusPlanFailed(payload.project, payload.workspace.Name, err))
		return
	}

	logError(h.UpdateJobStatusReviewRequired(payload.project, payload.workspace.Name, result))
}

func (r *LocalRunner) doApply(payload *payload, h core.ServiceProvider) {
	dir := r.getWorkDir(payload)

	//logError(h.UpdateJobStatusApplyInProgress(payload.project, payload.workspace.Name))

	output := bytes.NewBuffer(nil)
	finish := false
	go func() {
		for !finish {
			logError(h.SaveApplyLog(payload.project, payload.workspace.Name, output.String()))
			time.Sleep(1 * time.Second)
		}
		logError(h.SaveApplyLog(payload.project, payload.workspace.Name, output.String()))
	}()
	defer func() {
		finish = true
	}()

	err := r.tf.Apply(dir, output)
	if err != nil {
		logError(h.UpdateJobStatusApplyFailed(payload.project, payload.workspace.Name, err))
		return
	}

	logError(h.UpdateJobStatusSucceeded(payload.project, payload.workspace.Name))
}

func (r *LocalRunner) getWorkDir(payload *payload) string {
	return filepath.Join(r.dir, payload.project, payload.workspace.Name)
}

func (r *LocalRunner) initWorkDir(payload *payload) (string, error) {
	d := r.getWorkDir(payload)

	err := os.RemoveAll(d)
	if err != nil {
		return "", err
	}

	err = os.MkdirAll(d, 0755)
	if err != nil {
		return "", err
	}

	return d, nil
}

func logError(err error) {
	if err != nil {
		log.Printf("ERROR: %s", err.Error())
	}
}

func saveOutput(payload *payload, svc core.ServiceProvider) (*bytes.Buffer, func()) {
	output := bytes.NewBuffer(nil)
	finish := false
	go func() {
		for finish {
			logError(svc.SaveApplyLog(payload.project, payload.workspace.Name, output.String()))
			time.Sleep(1 * time.Second)
		}
		logError(svc.SaveApplyLog(payload.project, payload.workspace.Name, output.String()))
	}()

	return output, func() { finish = true }
}

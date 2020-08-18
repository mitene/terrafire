package controller

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
)

type ExecutorEcs struct {
	config *ExecutorEcsConfig
}

type ExecutorEcsConfig struct {
	Cluster          string
	TaskDefinition   string
	ContainerName    string
	CapacityProvider string
	Subnets          []string
	SecurityGroups   []string
	AssignPublicIp   bool
}

func NewEcsExecutor(config *ExecutorEcsConfig) Executor {
	return &ExecutorEcs{
		config: config,
	}
}

func (r *ExecutorEcs) Plan(project string, workspace string) (Process, error) {
	return r.run("plan", project, workspace)
}

func (r *ExecutorEcs) Apply(project string, workspace string) (Process, error) {
	return r.run("apply", project, workspace)
}

func (r *ExecutorEcs) run(phase string, project string, workspace string) (Process, error) {
	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}
	svc := ecs.New(sess)

	resp, err := svc.RunTask(&ecs.RunTaskInput{
		Cluster:        aws.String(r.config.Cluster),
		TaskDefinition: aws.String(r.config.TaskDefinition),
		Overrides: &ecs.TaskOverride{
			ContainerOverrides: []*ecs.ContainerOverride{
				{
					Name:    aws.String(r.config.ContainerName),
					Command: aws.StringSlice([]string{"run", phase, project, workspace}),
				},
			},
		},
		CapacityProviderStrategy: []*ecs.CapacityProviderStrategyItem{
			{
				CapacityProvider: aws.String(r.config.CapacityProvider),
			},
		},
		NetworkConfiguration: &ecs.NetworkConfiguration{
			AwsvpcConfiguration: &ecs.AwsVpcConfiguration{
				Subnets:        aws.StringSlice(r.config.Subnets),
				SecurityGroups: aws.StringSlice(r.config.SecurityGroups),
				AssignPublicIp: func() *string {
					if r.config.AssignPublicIp {
						return aws.String("ENABLED")
					} else {
						return aws.String("DISABLED")
					}
				}(),
			},
		},
	})
	if err != nil {
		return nil, err
	}
	if len(resp.Failures) > 0 {
		return nil, fmt.Errorf(aws.StringValue(resp.Failures[0].Reason))
	}
	task := resp.Tasks[0]

	return &ecsProcess{
		svc:     svc,
		cluster: r.config.Cluster,
		taskArn: aws.StringValue(task.TaskArn),
	}, nil
}

type ecsProcess struct {
	svc     *ecs.ECS
	cluster string
	taskArn string
}

func (p *ecsProcess) wait() error {
	return p.svc.WaitUntilTasksStopped(&ecs.DescribeTasksInput{
		Cluster: aws.String(p.cluster),
		Tasks:   aws.StringSlice([]string{p.taskArn}),
	})
}

func (p *ecsProcess) cancel() error {
	_, err := p.svc.StopTask(&ecs.StopTaskInput{
		Cluster: aws.String(p.cluster),
		Task:    aws.String(p.taskArn),
	})
	return err
}

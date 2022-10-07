package job

import (
	"sync"
	"time"

	"github.com/docker/docker/api/types"
)

type Job struct {
	Image    string   `json:"image"`
	Cmd      []string `json:"cmd"`
	Duration int      `json:"duration"`
}

type BaseJob struct {
	StartTime    time.Time
	TotalRunTime time.Duration
	Duration     time.Duration
}

type DockerJob struct {
	BaseJob
	types.Container
}

type SharedJobsArray struct {
	mu   sync.Mutex
	jobs []Job
}

type SharedDockerJobsMap struct {
	mu   sync.Mutex
	jobs map[string]DockerJob
}

func (j *BaseJob) UpdateTotalRunTime(time.Time) error {
	j.TotalRunTime += time.Since(j.StartTime)
	return nil
}

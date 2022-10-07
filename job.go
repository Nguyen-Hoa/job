package worker

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

func (j *BaseJob) UpdateTotalRunTime(time.Time) error {
	j.TotalRunTime += time.Since(j.StartTime)
	return nil
}

type SharedJobsArray struct {
	mu   sync.Mutex
	jobs []Job
}

func (j *SharedJobsArray) Get(index int) Job {
	var job Job
	j.mu.Lock()
	job = j.jobs[index]
	j.mu.Unlock()
	return job
}

func (j *SharedJobsArray) Pop() Job {
	var job Job
	j.mu.Lock()
	job, j.jobs = j.jobs[0], j.jobs[1:]
	j.mu.Unlock()
	return job
}

func (j *SharedJobsArray) Append(job Job) {
	j.mu.Lock()
	j.jobs = append(j.jobs, job)
	j.mu.Unlock()
}

type SharedDockerJobsMap struct {
	mu   sync.Mutex
	jobs map[string]DockerJob
}

func (j *SharedDockerJobsMap) Get(key string) DockerJob {
	var job DockerJob
	j.mu.Lock()
	job = j.jobs[key]
	j.mu.Unlock()
	return job
}

func (j *SharedDockerJobsMap) Delete(key string) {
	j.mu.Lock()
	delete(j.jobs, key)
	j.mu.Unlock()
}

func (j *SharedDockerJobsMap) Add(key string, job DockerJob) {
	j.mu.Lock()
	j.jobs[key] = job
	j.mu.Unlock()
}

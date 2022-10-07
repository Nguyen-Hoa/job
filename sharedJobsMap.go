package job

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

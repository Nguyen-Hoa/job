package job

func (j *SharedDockerJobsMap) Init() {
	j.mu.Lock()
	j.jobs = make(map[string]DockerJob)
	j.mu.Unlock()
}

func (j *SharedDockerJobsMap) InitFromMap(jobs map[string]DockerJob) {
	j.mu.Lock()
	j.jobs = jobs
	j.mu.Unlock()
}

func (j *SharedDockerJobsMap) Get(key string) (DockerJob, bool) {
	var job DockerJob
	var exists bool
	j.mu.Lock()
	job, exists = j.jobs[key]
	j.mu.Unlock()
	return job, exists
}

func (j *SharedDockerJobsMap) Delete(key string) {
	j.mu.Lock()
	delete(j.jobs, key)
	j.mu.Unlock()
}

func (j *SharedDockerJobsMap) Update(key string, job DockerJob) {
	j.mu.Lock()
	j.jobs[key] = job
	j.mu.Unlock()
}

func (j *SharedDockerJobsMap) Length() int {
	return len(j.jobs)
}

func (j *SharedDockerJobsMap) Exists(key string) bool {
	_, exists := j.jobs[key]
	return exists
}

func (j *SharedDockerJobsMap) Refresh(keys []string) {
	key_map := make(map[string]string, len(keys))
	for _, k := range keys {
		key_map[k] = k
	}

	for _, job := range j.jobs {
		_, exists := key_map[job.ID]
		if !exists {
			j.Delete(job.ID)
		}
	}
}

func (j *SharedDockerJobsMap) Keys() []string {
	keys := make([]string, 0, j.Length())
	for k := range j.jobs {
		keys = append(keys, k)
	}
	return keys
}

func (j *SharedDockerJobsMap) Snap() map[string]DockerJob {
	return j.jobs
}

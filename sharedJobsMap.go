package job

func (j *SharedDockerJobsMap) Init() {
	j.jobs = make(map[string]DockerJob)
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
	for job := range j.jobs {
		if !contains(keys, job) {
			j.Delete(job)
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

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

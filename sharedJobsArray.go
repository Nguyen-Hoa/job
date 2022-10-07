package job

func (j *SharedJobsArray) Get(index int) Job {
	var job Job
	j.mu.Lock()
	job = j.jobs[index]
	j.mu.Unlock()
	return job
}

func (j *SharedJobsArray) Length() int {
	return len(j.jobs)
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

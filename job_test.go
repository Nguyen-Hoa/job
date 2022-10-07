package job

import (
	"log"
	"sync"
	"testing"
)

func TestSharedJobsArrayAccess(t *testing.T) {
	log.Print("testing SharedJobsArray concurrect access")
	jobs := SharedJobsArray{}

	count := []int{1, 1, 1}
	var wg sync.WaitGroup
	for _, val := range count {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			jobs.Append(Job{Image: "someImage", Cmd: []string{"someCmd"}, Duration: val})
			log.Println(jobs.Length())
		}(val)
	}
	wg.Wait()

	for _, val := range count {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			jobs.Pop()
			log.Println(jobs.Length())
		}(val)
	}
	wg.Wait()
}

func TestSharedDockerJobsMapAccess(t *testing.T) {
	log.Print("testing SharedDockerJobsMap concurrect access")

	jobs := SharedDockerJobsMap{}
	jobs.Init()

	// Try to retrived a key that doesn't exist
	// `exist` should be false
	_, exists := jobs.Get("non-existant job")
	if exists {
		log.Print("error when accessing non-existant job")
	}

	// Using goroutines,
	// Add a few jobs, then delete them
	count := []string{"one", "two", "three"}
	var wg sync.WaitGroup
	for _, val := range count {
		wg.Add(1)
		go func(val string) {
			defer wg.Done()
			jobs.Update(val, DockerJob{})
			log.Print(jobs.Length())
		}(val)
	}
	wg.Wait()

	for _, val := range count {
		wg.Add(1)
		go func(val string) {
			defer wg.Done()
			jobs.Delete(val)
			log.Print(jobs.Length())
		}(val)
	}
	wg.Wait()
}

package job

import (
	"log"
	"sync"
	"testing"
)

func TestSharedAccess(t *testing.T) {
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

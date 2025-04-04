package cron

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Job represents a scheduled job
type Job struct {
	Schedule string
	Task     func()
	nextRun  time.Time
	index    int // needed for heap interface
}

// JobHeap is a priority queue of Jobs sorted by nextRun
// Implements heap.Interface
type JobHeap []*Job

func (h JobHeap) Len() int           { return len(h) }
func (h JobHeap) Less(i, j int) bool { return h[i].nextRun.Before(h[j].nextRun) }
func (h JobHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *JobHeap) Push(x any) {
	job := x.(*Job)
	job.index = len(*h)
	*h = append(*h, job)
}

func (h *JobHeap) Pop() any {
	old := *h
	n := len(old)
	job := old[n-1]
	job.index = -1
	*h = old[0 : n-1]
	return job
}

// Scheduler holds jobs and runs them at the right time
type Scheduler struct {
	mu   sync.Mutex
	h    JobHeap
	done chan struct{}
}

func NewScheduler() *Scheduler {
	h := make(JobHeap, 0)
	heap.Init(&h)
	return &Scheduler{
		h:    h,
		done: make(chan struct{}),
	}
}

func (s *Scheduler) AddJob(cronExpr string, task func()) {
	next, err := getNextFromCron(cronExpr, time.Now())
	if err != nil {
		fmt.Println("Invalid cron expression:", err)
		return
	}
	job := &Job{
		Schedule: cronExpr,
		Task:     task,
		nextRun:  next,
	}
	s.mu.Lock()
	heap.Push(&s.h, job)
	s.mu.Unlock()
}

func (s *Scheduler) Start() {
	go func() {
		for {
			s.mu.Lock()
			if s.h.Len() == 0 {
				s.mu.Unlock()
				time.Sleep(1 * time.Second)
				continue
			}

			nextJob := s.h[0]
			now := time.Now()
			sleepDuration := nextJob.nextRun.Sub(now)
			if sleepDuration > 0 {
				s.mu.Unlock()
				time.Sleep(sleepDuration)
				continue
			}

			// Pop and run the job
		heap.Pop(&s.h)
		s.mu.Unlock()

			go nextJob.Task()

			// Reschedule using last run time + 1s to avoid looping
		nextJob.nextRun, _ = getNextFromCron(nextJob.Schedule, nextJob.nextRun.Add(1*time.Second))
		s.mu.Lock()
		heap.Push(&s.h, nextJob)
		s.mu.Unlock()
		}
	}()
}

// getNextFromCron parses a cron string (sec min hour * * *) and calculates next run
func getNextFromCron(expr string, from time.Time) (time.Time, error) {
	parts := strings.Split(expr, " ")
	if len(parts) != 6 {
		return time.Time{}, fmt.Errorf("expected 6 fields in cron expression")
	}

	secField, minField, hourField := parts[0], parts[1], parts[2]

	for i := 0; i < 86400; i++ { // Check next 24 hours
		next := from.Add(time.Duration(i) * time.Second)
		if matchesCronField(secField, next.Second()) &&
			matchesCronField(minField, next.Minute()) &&
			matchesCronField(hourField, next.Hour()) {
			return next, nil
		}
	}

	return time.Time{}, fmt.Errorf("no valid time found for cron expression")
}

func matchesCronField(field string, value int) bool {
	if field == "*" {
		return true
	}
	if strings.HasPrefix(field, "*/") {
		n, err := strconv.Atoi(strings.TrimPrefix(field, "*/"))
		return err == nil && value%n == 0
	}
	n, err := strconv.Atoi(field)
	return err == nil && value == n
}
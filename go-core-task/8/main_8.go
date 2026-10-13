package main

import (
	"context"

	"golang.org/x/sync/semaphore"
)

type SemaphoreWaitGroup struct {
	weight    int64
	semaphore *semaphore.Weighted
}

func (swg *SemaphoreWaitGroup) Add() {
	swg.semaphore.Acquire(context.TODO(), 1)
}

func (swg *SemaphoreWaitGroup) Done() {
	swg.semaphore.Release(1)
}

func (swg *SemaphoreWaitGroup) Wait() {
	defer swg.semaphore.Release(swg.weight)
	swg.semaphore.Acquire(context.TODO(), swg.weight)
}

func New(weight int64) *SemaphoreWaitGroup {
	return &SemaphoreWaitGroup{
		weight:    weight,
		semaphore: semaphore.NewWeighted(weight),
	}
}

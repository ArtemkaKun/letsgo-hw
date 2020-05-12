package main

import (
	"fmt"
	"sync"
)

type DirectorCache struct {
	Directors map[int]Director
	sync.Mutex
}

func (cache *DirectorCache) GetDirector(id int) (Director, error) {
	cache.Lock()
	defer cache.Unlock()

	if v, ok := cache.Directors[id]; ok {
		return v, nil
	}

	return Director{}, fmt.Errorf("Cannot find director with that ID\n")
}

type WorkersCache struct {
	Workers map[int]Employee
	sync.Mutex
}

func (cache *WorkersCache) GetCache(id int) (Employee, error) {
	cache.Lock()
	defer cache.Unlock()

	if v, ok := cache.Workers[id]; ok {
		return v, nil
	}

	return Employee{}, fmt.Errorf("Cannot find employee with that ID\n")
}

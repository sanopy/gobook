package main

import (
	"os"
	"sync"
)

type Cwd string

var mu sync.Mutex

func (c *Cwd) cd(dir string) error {
	mu.Lock()
	defer mu.Unlock()

	if err := os.Chdir(string(*c)); err != nil {
		return err
	}

	if err := os.Chdir(dir); err != nil {
		return err
	}

	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	*c = Cwd(wd)

	return nil
}

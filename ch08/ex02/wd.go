package main

import (
	"os"
	"sync"
)

type Wd string

var mu sync.Mutex

func (c *Wd) cd(dir string) error {
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
	*c = Wd(wd)

	return nil
}

func (c Wd) pwd() string {
	return string(c)
}

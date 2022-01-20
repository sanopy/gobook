package main

import (
	"os"
	"os/exec"
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

func (c Wd) ls(path string) ([]byte, error) {
	mu.Lock()
	defer mu.Unlock()

	if err := os.Chdir(string(c)); err != nil {
		return nil, err
	}

	return exec.Command("ls", "-l", path).Output()
}

func (c Wd) mkdir(path string) error {
	mu.Lock()
	defer mu.Unlock()

	if err := os.Chdir(string(c)); err != nil {
		return err
	}

	return os.MkdirAll(path, 0755)
}

func (c Wd) create(path string) (*os.File, error) {
	mu.Lock()
	defer mu.Unlock()

	if err := os.Chdir(string(c)); err != nil {
		return nil, err
	}

	return os.Create(path)
}

func (c Wd) open(path string) (*os.File, error) {
	mu.Lock()
	defer mu.Unlock()

	if err := os.Chdir(string(c)); err != nil {
		return nil, err
	}

	return os.Open(path)
}

func (c Wd) remove(path string) error {
	mu.Lock()
	defer mu.Unlock()

	if err := os.Chdir(string(c)); err != nil {
		return nil
	}

	return os.Remove(path)
}

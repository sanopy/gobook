package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

const targetModule = "github.com/sanopy/gobook/..."

type Package struct {
	ImportPath string
	Deps       []string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: $ %s [packages]...\n", os.Args[0])
		os.Exit(1)
	}

	modulePkgs, err := runGoList(targetModule)
	if err != nil {
		fmt.Fprintf(os.Stderr, "runGoList: %v\n", err)
		os.Exit(1)
	}
	targetPkgs, err := runGoList(os.Args[1:]...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "runGoList: %v\n", err)
		os.Exit(1)
	}

	printDependency(modulePkgs, targetPkgs)
}

func runGoList(args ...string) ([]*Package, error) {
	cmdArgs := append([]string{"list", "-json"}, args...)
	cmd := exec.Command("go", cmdArgs...)
	r, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	defer r.Close()

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(r)
	pkgs := []*Package{}
	for {
		pkg := new(Package)
		err := decoder.Decode(&pkg)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		pkgs = append(pkgs, pkg)
	}

	return pkgs, nil
}

func printDependency(pkgs, targets []*Package) {
	for _, pkg := range pkgs {
		if isDependentPackage(pkg, targets) {
			printPackageInfo(pkg)
		}
	}
}

func isDependentPackage(pkg *Package, targets []*Package) bool {
	cnt := 0
	for _, dep := range pkg.Deps {
		for _, target := range targets {
			if dep == target.ImportPath {
				cnt++
			}
		}
	}
	return cnt == len(targets)
}

func printPackageInfo(pkg *Package) {
	fmt.Printf("%s\n", pkg.ImportPath)
	fmt.Printf("  dependency:\n")
	fmt.Printf("    %s\n", strings.Join(pkg.Deps, "\n    "))
}

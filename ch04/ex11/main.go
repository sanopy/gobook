// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/sanopy/gobook/ch04/ex11/github"
)

func main() {
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getRepo := getCmd.String("repo", "", "target repository")
	getNum := getCmd.Int("num", 1, "issue number")

	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	createRepo := createCmd.String("repo", "", "target repository")
	createTitle := createCmd.String("title", "", "issue title")
	createBody := createCmd.String("body", "", "issue body")

	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	updateRepo := updateCmd.String("repo", "", "target repository")
	updateTitle := updateCmd.String("title", "", "issue title")
	updateBody := updateCmd.String("body", "", "issue body")
	updateNum := updateCmd.Int("num", 1, "issue number")

	closeCmd := flag.NewFlagSet("close", flag.ExitOnError)
	closeRepo := closeCmd.String("repo", "", "target repository")
	closeNum := closeCmd.Int("num", 1, "issue number")

	if len(os.Args) < 2 {
		fmt.Println("expected 'get', 'create', 'update' or 'close' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		getCmd.Parse(os.Args[2:])
		get(*getRepo, *getNum)
	case "create":
		createCmd.Parse(os.Args[2:])
		create(*createRepo, *createTitle, *createBody)
	case "update":
		updateCmd.Parse(os.Args[2:])
		update(*updateRepo, *updateTitle, *updateBody, *updateNum)
	case "close":
		closeCmd.Parse(os.Args[2:])
		close(*closeRepo, *closeNum)
	default:
		fmt.Println("expected 'get', 'create', 'update' or 'close' subcommands")
		os.Exit(1)
	}
}

func get(repo string, num int) {
	result, err := github.GetIssue(repo, num)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("#%d %s @%s\n\n",
		result.Number, result.CreatedAt.Format("2006-01-02 15:04:05"), result.User.Login)
	fmt.Printf("%s\n\n", result.Title)
	fmt.Printf("%s\n", result.Body)
}

func create(repo, title, body string) {
	if title == "" {
		title = inputFromEditor("${input issue title here}")
	}
	if body == "" {
		body = inputFromEditor("${input issue body here}")
	}

	post := &github.PostIssue{Title: title, Body: body}
	result, err := github.CreateIssue(repo, post)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("An Issue was successfully created.\n\n")
	fmt.Printf("%s\n", result.HTMLURL)
}

func update(repo, title, body string, num int) {
	issue, err := github.GetIssue(repo, num)
	if err != nil {
		log.Fatal(err)
	}

	if title == "" {
		title = inputFromEditor(issue.Title)
	}
	if body == "" {
		body = inputFromEditor(issue.Body)
	}

	post := &github.PostIssue{Title: title, Body: body}
	result, err := github.UpdateIssue(repo, post, num)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("An Issue was successfully updated.\n\n")
	fmt.Printf("%s\n", result.HTMLURL)
}

func close(repo string, num int) {
	result, err := github.CloseIssue(repo, num)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("An Issue was successfully closed.\n\n")
	fmt.Printf("%s\n", result.HTMLURL)
}

func inputFromEditor(placeholder string) string {
	path := "/tmp/edit.txt"
	w, err := os.Create(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "main: %q", err)
		os.Exit(1)
	}
	defer w.Close()

	_, err = fmt.Fprintln(w, placeholder)
	if err != nil {
		fmt.Fprintf(os.Stderr, "main: %q", err)
		os.Exit(1)
	}

	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}

	cmd := exec.Command(editor, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	r, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "main: %q", err)
		os.Exit(1)
	}
	defer r.Close()

	b, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "main: %q", err)
		os.Exit(1)
	}
	return string(b)
}

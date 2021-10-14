// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sanopy/gobook/ch04/ex10/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	moreThanYear := []*github.Issue{}
	lessThanYear := []*github.Issue{}
	lessThanMonth := []*github.Issue{}

	now := time.Now()
	aMonthAgo := now.AddDate(0, -1, 0)
	aYearAgo := now.AddDate(-1, 0, 0)

	for _, item := range result.Items {
		if item.CreatedAt.Before(aYearAgo) {
			moreThanYear = append(moreThanYear, item)
		} else if item.CreatedAt.Before(aMonthAgo) {
			lessThanYear = append(lessThanYear, item)
		} else {
			lessThanMonth = append(lessThanMonth, item)
		}
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	fmt.Println("-------------------- less than a month --------------------")
	for _, item := range lessThanMonth {
		fmt.Printf("#%-5d %s %9.9s %.55s\n",
			item.Number, item.CreatedAt.Format("2006-01-02 15:04:05"), item.User.Login, item.Title)
	}
	fmt.Println("-------------------- less than a year --------------------")
	for _, item := range lessThanYear {
		fmt.Printf("#%-5d %s %9.9s %.55s\n",
			item.Number, item.CreatedAt.Format("2006-01-02 15:04:05"), item.User.Login, item.Title)
	}
	fmt.Println("-------------------- more than year --------------------")
	for _, item := range moreThanYear {
		fmt.Printf("#%-5d %s %9.9s %.55s\n",
			item.Number, item.CreatedAt.Format("2006-01-02 15:04:05"), item.User.Login, item.Title)
	}
}

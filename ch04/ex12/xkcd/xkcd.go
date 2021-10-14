package xkcd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

const ApiURL = "https://xkcd.com/%d/info.0.json"

type XkcdItems struct {
	Item *XkcdItem
	Num  int
	Ok   bool
}

type XkcdItem struct {
	Day        string
	Month      string
	Year       string
	Num        int
	Link       string
	Title      string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
}

var Wg sync.WaitGroup

func CallInfoAPI(num int, ch chan<- *XkcdItems) {
	defer Wg.Done()
	url := fmt.Sprintf(ApiURL, num)
	resp, err := http.Get(url)
	if err != nil {
		ch <- &XkcdItems{nil, num, false}
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		ch <- &XkcdItems{nil, num, false}
		return
	}

	var result XkcdItem
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		ch <- &XkcdItems{nil, num, false}
		return
	}
	ch <- &XkcdItems{&result, num, true}
	return
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"text/template"
)

func main() {
	db := database{data: map[string]dollars{"shoes": 50, "socks": 5}}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database struct {
	data map[string]dollars
	sync.Mutex
}

type ItemDatabase struct {
	Items []Item
}

type Item struct {
	Item  string
	Price dollars
}

var itemList = template.Must(template.New("itemlist").Parse(`
<h1>{{.Items | len}} items</h1>
<table>
<tr style='text-align: left'>
	<th>item</th>
	<th>price</th>
</tr>
{{range .Items}}
<tr>
	<td>{{.Item}}</td>
	<td>{{.Price}}</td>
</tr>
{{end}}
</table>
`))

func (db *database) list(w http.ResponseWriter, req *http.Request) {
	db.Lock()
	defer db.Unlock()
	itemDb := ItemDatabase{[]Item{}}
	for item, price := range db.data {
		itemDb.Items = append(itemDb.Items, Item{item, price})
	}
	if err := itemList.Execute(w, itemDb); err != nil {
		fmt.Fprintf(w, "%q", err)
	}
}

func (db *database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	db.Lock()
	defer db.Unlock()
	if price, ok := db.data[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db *database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	db.Lock()
	defer db.Unlock()
	if _, ok := db.data[item]; ok {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "already exist item: %q\n", item)
	} else {
		p, err := strconv.ParseFloat(price, 32)
		if err != nil {
			fmt.Fprintf(w, "price parse failed: %v\n", err)
			return
		}
		db.data[item] = dollars(p)
		fmt.Fprintf(w, "%q created\n", item)
	}
}

func (db *database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	db.Lock()
	defer db.Unlock()
	if _, ok := db.data[item]; ok {
		p, err := strconv.ParseFloat(price, 32)
		if err != nil {
			fmt.Fprintf(w, "price parse failed: %v\n", err)
			return
		}
		db.data[item] = dollars(p)
		fmt.Fprintf(w, "%q updated\n", item)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db *database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	db.Lock()
	defer db.Unlock()
	if _, ok := db.data[item]; ok {
		delete(db.data, item)
		fmt.Fprintf(w, "%q deleted\n", item)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

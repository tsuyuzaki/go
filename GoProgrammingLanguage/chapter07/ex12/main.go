/**
 * テキストではなく HTML の表として出力を印字するように /list に対するハンドラを変更しなさい。html/template (4.6節) が役立つでしょう。
 */
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	db.doList(w)
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item, ok := getItem(w, req)
	if !ok {
		return
	}
	if _, ok = db[item]; ok {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, "[%s] already exists.\n", item)
		return
	}
	price, ok := getPrice(w, req)
	if !ok {
		return
	}
	db[item] = price
	w.WriteHeader(http.StatusCreated)
	db.doList(w)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item, ok := getItem(w, req)
	if !ok {
		return
	}
	if _, ok = db[item]; !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "[%s] not found.\n", item)
		return
	}
	price, ok := getPrice(w, req)
	if !ok {
		return
	}
	db[item] = price
	w.WriteHeader(http.StatusOK)
	db.doList(w)
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item, ok := getItem(w, req)
	if !ok {
		return
	}
	if _, ok = db[item]; !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "[%s] not found.\n", item)
		return
	}
	delete(db, item)
	w.WriteHeader(http.StatusOK)
	db.doList(w)
}

func (db database) doList(w http.ResponseWriter) {
	list := template.Must(template.New("table").Parse(`<!DOCTYPE html>
<html lang=ja>
    <head>
        <meta charset="utf-8">
    </head>
    <body>
        <h1>Chapter07 Ex12</h1>
        <table>
        	<tr style="'text-align: left'">
        		<th>Item</th>
        		<th>Price</th>
        	</tr>
        	{{range $item, $price := .}}
        	<tr>
        		<td>{{$item}}</td>
        		<td>{{$price}}</td>
        	</tr>
        	{{end}}
        </table>
    </body>
</html>`))
	if err := list.Execute(w, db); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "list.Execute error. err:%v\n", err)
		return
	}
}

func getItem(w http.ResponseWriter, req *http.Request) (string, bool) {
	q := req.URL.Query()
	v, ok := q["item"]
	if !ok || len(v) != 1 || v[0] == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Can not identify a item. query[%s]\n", req.URL.RawQuery)
		return "", false
	}
	return v[0], true
}

func getPrice(w http.ResponseWriter, req *http.Request) (dollars, bool) {
	q := req.URL.Query()
	v, ok := q["price"]
	if !ok || len(v) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Can not identify a price. query[%s]\n", req.URL.RawQuery)
		return 0, false
	}
	f, err := strconv.ParseFloat(v[0], 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Price parse error:[%v] query[%s]\n", err, req.URL.RawQuery)
		return 0, false
	}
	if f < 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid price %f\n", f)
		return 0, false
	}
	return dollars(f), true
}

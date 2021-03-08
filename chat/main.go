package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

type templateHandler struct {
	once     sync.Once
	templ    *template.Template
	filename string
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})

	t.templ.Execute(w, req)
}

func main() {
	addr := flag.String("addr", ":8080", "the addr of the application")
	flag.Parse()
	r := newRoom()
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)

	go r.run()

	log.Println("Starting web server on ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}

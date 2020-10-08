package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)
// The type []byte means "a byte slice". (See Slices: usage and internals for more on slices.)
// The Body element is a []byte rather than string because that is the type expected by the io libraries we will use, as you'll see below.
type Page struct {
	Title string
	Body []byte
}

/*
The Page struct describes how page data will be stored in memory.
But what about persistent storage?
We can address that by creating a save method on Page:
*/
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
        return nil, err
    }
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main(){
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
// func main() {
//     p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
//     p1.save()
//     p2, _ := loadPage("TestPage")
//     fmt.Println(string(p2.Body))
// }

// package main

// import (
//     "fmt"
//     "log"
//     "net/http"
// )

// func handler(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
// }

// func main() {
//     http.HandleFunc("/", handler)
//     log.Fatal(http.ListenAndServe(":8080", nil))
// }
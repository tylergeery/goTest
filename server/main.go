package gofun

import (
    // "os"
    // "path/filepath"
    // "fmt"
    "net/http"
    // "strings"
    "log"
    "html/template"
    "encoding/json"
)

type Test struct {
    Name string
}



func test(w http.ResponseWriter, r *http.Request) {
    test := Test{Name: r.URL.Path}
    obj, _ := json.Marshal(test)

    w.Header().Set("Content-Type", "application/json")
    w.Write(obj)
}

// root site access, respond with a template
func showSite(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")

    l := MainLayout()
    t := template.New("main template") // Create a template.
	t, _ = t.Parse(l)  // Parse template file.

	t.Execute(w, nil)
}

func main() {
    http.HandleFunc("/", showSite) // set router
    err := http.ListenAndServe(":8080", nil) // set listen port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

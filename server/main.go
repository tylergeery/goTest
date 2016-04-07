package main

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

var mainTemplate = `
<DOCTYPE html>
  <head>
    <title>GoFun Web Chat App</title>
    <link href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-1q8mTJOASx8j1Au+a5WDVnPi2lkFfwwEAa8hDDdjZlpLegxhjVME1fgjWPGmkzs7" crossorigin="anonymous">
  </head>
  <body>
    <div id="root">
      <!-- React takes over here -->
    </div>

    <script type="text/javascript" src="/js/main.bundle.js"></script>
  </body>
</html>
`

func test(w http.ResponseWriter, r *http.Request) {
    test := Test{Name: r.URL.Path}
    obj, _ := json.Marshal(test)

    w.Header().Set("Content-Type", "application/json")
    w.Write(obj)
}

// root site access, respond with a template
func showSite(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")

	t := template.New("main template") // Create a template.
	t, _ = t.Parse(mainTemplate)  // Parse template file.

	t.Execute(w, nil)
}

func main() {
    http.HandleFunc("/", showSite) // set router
    err := http.ListenAndServe(":8080", nil) // set listen port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

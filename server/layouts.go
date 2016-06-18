package gofun

var mainTemplate = `
<DOCTYPE html>
  <head>
    <title>GoTest Web Chat App</title>
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

/*
The main app layout, should be considered default
*/
func MainLayout() string {
    return mainTemplate
}

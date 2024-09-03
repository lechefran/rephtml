# rephtml
A simple HTML file creator written in Go! rephtml supports both string and struct parameters.
### Simple Example
```
html := NewHtmlFile()
html.Style(`h1, h2, h3, h4, h5, h6, p {
		font-family: Arial;
		text-align: center;
	}`)
html.Style(`th, td {
		font-family: Arial;
		padding: 10px;
		text-align: center;
	}`)
html.Style(`table {
	margin-left: auto;
	margin-right: auto;
	width: 80%;
	}`)
html.H1("Test H1 Header")
html.Table([]string{"header1", "header2", "header3"},
          [][]string{{"record1", "record2", "record3"},
          {"record4", "record5", "record6"}})
html.P("Test paragraph for testing purposes")
html.Prepare()
html.WriteToFile("report.html") // create a file named report.html in the current directory
```

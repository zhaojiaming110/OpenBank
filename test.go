package main

import (
	"fmt"
	"html/template"
	"os"
)

type Movie struct {
	Title	string
	Year	int	`json:"released"`
	Actors	[]string
}

type MoviesResult struct {
	Number	int
	Items	[]Movie
}

//!+template
//const templ = `MovieCount: {{.Number}}
//{{range .Items}}
//	Title:	{{.Title}}
//	Year:	{{.Year}}
//	Actors:	{{.Actors}}
//-----------------------------
//{{end}}`
//!-template
//<h1>{{.TotalCount}} issues</h1>
//<table>
//<tr style='text-align: left'>
//<th>#</th>
//<th>State</th>
//<th>User</th>
//<th>Title</th>
//</tr>
//{{range .Items}}
//<tr>
//<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
//<td>{{.State}}</td>
//<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
//<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
//</tr>
//{{end}}
//</table>
const temp2 = `
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<h1>英叔电影总数： {{.Number}}</h1>
<table>
<tr style='text-align: left' >
	<th>电影名称</th>
	<th>上映年份</th>
	<th>演员表</th>
</tr>
{{range .Items}}
<tr>
	<td>{{.Title}}</td>
	<td>{{.Year}}</td>
	<td>{{.Actors}}</td>
</tr>
{{end}}
</table>
</head>
`

// !+exec
var report = template.Must(template.New("movielist").Parse(temp2))

func main() {
	var mr MoviesResult = MoviesResult{
		Number: 2,
		Items:  []Movie{
			{Title:"僵尸道长",Year:1942,Actors:[]string{"林正英", "秋生"}},
			{Title:"僵尸叔叔", Year:1945,Actors:[]string{"林正英", "文才"}},
		},
	}
	if err := report.Execute(os.Stdout, mr); err != nil {
		fmt.Println(err)
	}
}








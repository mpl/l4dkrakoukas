package godle

import (
	"html/template"
)

var funcMap = template.FuncMap{ "toString": toString, }

var rootTemplate = template.Must(template.New("root").Funcs(funcMap).Parse(rootHTML))

const rootHTML = `
<!DOCTYPE html>
<html>
<head>
	<title>l4dkrakoukas</title>
</head>
<body>
Nope
</body>
</html>
`
var weekTemplate = template.Must(template.New("week").Funcs(funcMap).Parse(weekHTML))

// TODO: autogen?
const weekHTML = `
<!DOCTYPE html>
<html>
<head>
	<title>l4dkrakoukas</title>
</head>
<body>
Planning pour semaine {{.W.Date}}
<form name="weekinput" action="/week/{{.W.Date}}" method="post">
<table>
	{{range $player, $schedule := .Foo}}
	<tr>
		<td> {{$player}} </td>
		{{range $day, $avail := $schedule}}
		<td>
			<input type="checkbox" name={{$player}}day value={{$day | toString}}
			{{if $avail}} checked="true" {{end}}
			>
		</td>
		{{end}}
	</tr>
	{{end}}
</table>
<input type="submit" value="Save">
</form>
</body>
</html>
`

/*
	{{end}}
	<tr>
		<th>Ptits Joueurs</th>
		<th>Vendredi</th>
		<th>Samedi</th>
		<th>Dimanche</th>
	</tr>
		<td><input type="checkbox" name="fooday" value={{.}}></td>
*/

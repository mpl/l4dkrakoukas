package godle

import (
	"html/template"
)

var rootTemplate = template.Must(template.New("root").Parse(rootHTML))

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
var weekTemplate = template.Must(template.New("week").Parse(weekHTML))

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
	{{with $dts := .DTS}}
	{{range $player, $schedule := .Foo}}
	<tr>
		<td> {{$player}} </td>
		{{range $day, $avail := $schedule}}
		<td>
			<input type="checkbox" name={{$player}}day value={{$day}}
			{{if $avail}} checked="true" {{end}}
			>
		</td>
		{{end}}
	</tr>
	{{end}}
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

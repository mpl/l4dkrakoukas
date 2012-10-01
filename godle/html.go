package godle

import (
	"html/template"
)

var funcMap = template.FuncMap{ "prettyDate": prettyDate, }

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
Planning pour semaine {{.Date | prettyDate }}
<form name="weekinput" action="/week/{{.Date}}" method="post">
<table>
	<tr>
		<th>Ptits Joueurs</th>
		<th>Lundi</th>
		<th>Mardi</th>
		<th>Mercredi</th>
		<th>Jeudi</th>
		<th>Vendredi</th>
		<th>Samedi</th>
		<th>Dimanche</th>
	</tr>
	{{range $player, $days := .Schedule}}
	<tr>
		<td> {{$player}} </td>
		{{range $day, $avail := $days}}
		<td>
			<input type="checkbox" name={{$player}}days value={{$day}}
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

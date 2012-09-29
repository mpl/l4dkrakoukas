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
Planning pour semaine {{.Date}}
<form name="weekinput" action="week" method="get">
<table>
	<tr>
		<th>Ptits Joueurs</th>
		<th>Vendredi</th>
		<th>Samedi</th>
		<th>Dimanche</th>
	</tr>
	<tr>
		<td>Asticot</td>
		<td><input type="checkbox" name="asticotday" value="friday"></td>
		<td><input type="checkbox" name="asticotday" value="saturday"></td>
		<td><input type="checkbox" name="asticotday" value="sunday"></td>
	</tr>
	<tr>
		<td>ChuckMaurice</td>
		<td><input type="checkbox" name="chuckmauriceday" value="friday"></td>
		<td><input type="checkbox" name="chuckmauriceday" value="saturday"></td>
		<td><input type="checkbox" name="chuckmauriceday" value="sunday"></td>
	</tr>
	<tr>
		<td>Posi</td>
	</tr>
	<tr>
		<td>Lagoule</td>
	</tr>
</table>
<input type="submit" value="Save">
</form>
</body>
</html>
`

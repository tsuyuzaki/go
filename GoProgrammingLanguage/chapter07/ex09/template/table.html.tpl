<!DOCTYPE html>
<html lang=ja>
	<head>
		<meta charset="utf-8">
	</head>
	<body>
		<h1>Chapter07 Ex09</h1>
		<table>
			<tr style="'text-align: left'">
				<th><a href="?sortKey=Title">Title</a></th>
				<th><a href="?sortKey=Artist">Artist</a></th>
				<th><a href="?sortKey=Album">Album</a></th>
				<th><a href="?sortKey=Year">Year</a></th>
				<th><a href="?sortKey=Length">Length</a></th>
			</tr>
			{{range .}}
			<tr>
				<td>{{.Title}}</td>
				<td>{{.Artist}}</td>
				<td>{{.Album}}</td>
				<td>{{.Year}}</td>
				<td>{{.Length}}</td>
			</tr>
			{{end}}
		</table>
	</body>
</html>

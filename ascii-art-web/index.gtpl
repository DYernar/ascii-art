<html>
    <head>
    <title></title>
    </head>
    <body>
        <form action="/getText" method="post">
            Text:<input type="text" name="text">
            <input type="submit" value="GetAscii">
        </form>
    {{.AsciiArt}}
    </body>
</html>

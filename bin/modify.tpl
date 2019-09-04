<html>
<head>
<title></title>
<meta charset="UTF-8"/>
<meta name="viewport" content= "width=device-width, initial-scale=1.0,maximum-scale=1.0,user-scalable=0" />
</head>
<body>
<form action="http://123.206.55.31:8888/modify" method="post">
<pre>
学豆:	<input type="text" name="xd" value={{.Num}} />
密码:	<input type="password" name="root" />  原因: <input type="text" name="reason" value="" /> <input type="submit" value="修改" />
</p>

<input type="hidden" name="lastxd" value={{.LASTXD}} />
</form>
</body>
</html>

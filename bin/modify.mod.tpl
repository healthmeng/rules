<html>
<head>
<title></title>
<meta charset="UTF-8"/>
<meta name="viewport" content= "width=device-width, initial-scale=1.0,maximum-scale=1.0,user-scalable=0" />
</head>
<body>
<form action="http://123.206.55.31:8888/modify" method="post">
<pre>
学豆:	<input type="text" name="xd" value={{.Num}} readonly=true/>
修改:	<input type="text" name="modify" value=0 /><input type="button" value="奖" onclick="onxdplus()" /><input type="button" value="罚" onclick="onxdminus()">
原因:	<input type="text" name="reason" />

生活豆:	<input type="text" name="shd" value={{.Shd}} />
修改:	<input type="text" name="modify" value=0 /><input type="button" value="奖" onclick="onplusshd()" /><input type="button" value="罚" onclick="onminusshd()" />
原因:	<input type="text" name="reason" />

密码:	<input type="password" name="root" /> <input type="submit" value="修改" />	
</pre>
<input type="submit" value="奖罚记录" />
</p>
<input type="hidden" name="logs" value={{.Logs}} />
</form>
</body>
</html>

<html>
<head>
<meta charset="UTF-8"/>
<meta name="viewport" content= "width=device-width, initial-scale=1.0,maximum-scale=1.0,user-scalable=0" />
<title>
学豆商城
</title>
<link href="http://cdn-cnc.17zuoye.cn/public/skin/project/studentlottery/images/turntable-beans-4.png" rel="shortcut icon">
<script type="text/javascript" src="/static/count.js"></script>
</head>
<body>
<form name="xdinfo" action="/xdinfo" method="post">

<img src="http://cdn-cnc.17zuoye.cn/public/skin/reward/imagesV1/w-gold-icon.png" /> 学豆: <input type="text" name="xd" readonly="true" style="width:70px" value={{.NUM}} /> 
<hr />
<img src="http://cdn-cnc.17zuoye.cn/public/skin/project/studentlottery/images/turntable-beans-6.png" weight="20" height="20" /> 兑换奖励: <br />
{{range $k,$v:=.Goods}}{{$v.Name}}	<input type="button"  value="{{$v.Price}}" onclick="ondhval('{{$v.Name}}',value)" /><br />{{end}} 
<hr />
<input type="hidden" name="spendxd" />
<input type="hidden" name="obj" />
<input type="hidden" name="log" value={{.LOG}} />
<input type="hidden" name="modlog" value={{.MODLOG}} />
<input type="hidden" name="helpxd" value={{.HELPXD}} />
<style type="text/css">
a:link,a:visited{
 text-decoration:none;  
}
a:hover{
 text-decoration:underline; 
}
</style>
<pre>
<img src="https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1505734464797&di=3e1708e94bdf2fd76b63d9901c6b94fc&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F01021b55b4f0c66ac725ca50891bd1.jpg" height="24" weight="24" /> <a href="javascript:;" onclick="onmodlog()"> 奖罚记录 </a>    <img src="http://uus-img7.android.d.cn/content_pic/201509/behpic/icon/60/2-63060/icon-1442390018925.png" height="20" weight="20" /> <a href="javascript:;" onclick="onlog()"> 兑换记录 </a>

<img src="https://ss2.bdstatic.com/70cFvnSh_Q1YnxGkpoWK1HF6hhy/it/u=4286309319,2835405153&fm=15&gp=0.jpg" height="25" weight="25" /> <a href="javascript:;" onclick="onhelpxd()"> 学豆奖励帮助 </a>
</pre>
</form>
</body>
</html>

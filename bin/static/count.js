function ondhval(name,val)
{
 var left=document.forms["xdinfo"].xd.value;
 if (parseInt(left)<parseInt(val))
 {
	alert("学豆不够，继续努力!");
	return;
 }
 var ret=confirm("确认使用"+val+"学豆兑换"+name+"?");
 if (ret==true)
 {
	document.forms["xdinfo"].spendxd.value=val;
	document.forms["xdinfo"].obj.value=name;
	document.forms["xdinfo"].submit();
 }
}

function onlog()
{
 alert(document.forms["xdinfo"].log.value);
}

function onmodlog()
{
 alert(document.forms["xdinfo"].modlog.value);
}


function onhelpxd()
{
 alert(document.forms["xdinfo"].helpxd.value);

}

function onhelpshd()
{
 alert(document.forms["xdinfo"].helpshd.value);
}



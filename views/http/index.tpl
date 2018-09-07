<link rel="stylesheet" href="/static/css/jquery.jsonview.css" />
<script type="text/javascript" src="/static/js/jquery.jsonview.js"></script>
<div style="text-align:center">
	<form action="" method="post">
	Url:<input size="100" type="text" name="url" value="http://yun.zjer.cn/index.php?r=api/school/GetContentlistByschoolid" /><br/>
	<br/><br/>
	Param:<textarea rows="10" cols="100" name="data" id="data" >
	{"isclassbrand":"1","schoolid":"3433000319","page":"1","pagesize":"3","version":"V2.4.0"}
	</textarea><br/>
	<input type="button" id="sub" value="获取"/>
	</form>
</div>
<br/><br/>
<hr/>
<div id="datas"></div>
<div id="json"></div>
<script>
function getCurrentTime()
{
	var date = new Date();
	y = date.getFullYear(); //获取年
	m = date.getMonth()+1;//获取月
	d = date.getDate(); //获取当日
	h = date.getHours();//获取小时
	s = date.getMinutes();//获取分钟
	i = date.getSeconds();//获取秒
	//d.getMilliseconds()//获取毫秒
	return y+"-"+m+"-"+d+" "+h+":"+s+":"+i;
}
$(document).ready(function(){
	alert(getCurrentTime());
	$("#sub").click(function(){
		var url = $("input[name=url]").val();
		var data =$("#data").val();
		if(url.trim()=="" || data.trim()=="")
		{
			alert("请输入");
		}else{
				alert(url+"\n"+data);
				$.ajax({
				type:"POST",
	            url:{{urlfor "HttpController.Post"}},
	            data:'data='+data+'&url='+url,
	            dataType:"json",
	            timeout:5000,
	            success:function(json){                  
	                //$('#datas').html(json);
					$("#json").JSONView(json);
	            },
			});
		}
		
	});

});
</script>
<div class="container">
		 <!--请求得到Json数据-->
	      <div style="width:100%;height:50px;">
          <button onclick="getjson()" class="btn btn-primary">得到Json</button>
          <label id="txtjson"></label>
          </div>
          <!--请求得到Xml数据-->
        <div style="width:100%;height:50px;">
            <button onclick="getxml()" class="btn btn-primary">得到Xml</button>
            <label id="txtxml"></label>
        </div>
        <!--请求得到Jsonp数据-->
        <div style="width:100%;height:50px;">
          <button onclick="getjsonp()" class="btn btn-primary">得到Jsonp</button>
          <label id="txtjsonp"></label>
        </div>
        <!--请求得到字典数据-->
        <div style="width:100%;height:50px;">
            <button onclick="getdictionary()" class="btn btn-primary">得到字典</button>
            <label id="txtdictionary"></label>
        </div>
 
        <!--请求得到字典数据-->
        <div style="width:100%;height:50px;">
            <input type="text" id="search" placeholder="请输入参数"/>
            <button onclick="getparams()" class="btn btn-primary">得到参数</button>
            <label id="txtparams"></label>
        </div>
       </div>
</div>

<script type="text/javascript">
	$(document).ready(function(){
		//alert('jquery生效');
	});
	function getjson(){
		//alert("得到Json");
		$.ajax({
               type:'get',
               url:'{{urlfor "ApiJsonController.Get"}}',
               dataType:'json',//此处的是json数据的格式
               data:{},
               success:function(result){
                   console.log('获取json的数据')
                   console.log(result.uid)
                   console.log(result.info)
                 $("#txtjson").html("json的结果："+result.uid+"|"+result.info);
               }
            })
	}
	//得到Xml
          function getxml(){
            $.ajax({
               type:'get',
               url:'{{urlfor "ApiXMLController.Get"}}',
               dataType:'xml',//此处的是xml数据的格式
               data:{},
               success:function(result){
                   console.log('获取xml的数据')
                   console.log(result)
                    
                 $("#txtxml").html("xml的结果："+$(result).text());
               }
            })
          }
           //得到jsonp
           function getjsonp(){
            $.ajax({
               type:'get',
               url:'{{urlfor "ApiJsonpController.Get"}}',              
               dataType:'jsonp',//此处的是jsonp数据的格式
               data:{},
               success:function(result){
                   console.log('获取jsonp的数据')
                   console.log(result)
                 $("#txtjsonp").html("jsonp的结果:"+result);
               }
            })
          }
          //得到字典
          function getdictionary(){
            $.ajax({
               type:'get',
               url:'{{urlfor "ApiDictionaryController.Get"}}',//此处的是json数据的格式
               data:{},
               success:function(result){
                   console.log('获取字典的数据')
                   console.log(result)
                 $("#txtdictionary").html("字典的结果:"+result.name+","+result.rows+","+result.flag);
               }
            })
          }
          //得到参数
          function getparams(){
            $.ajax({
               type:'get',
               url:'{{urlfor "ApiParamsController.Get"}}',//此处的是json数据的格式
               data:{
                 "name":$("#search").val()
               },
               success:function(result){
                   console.log('获取参数的数据')
                   console.log(result.json)
                 $("#txtparams").html("获取参数结果:"+result.name+","+result.rows+","+result.flag);
               }
            })
          }
</script>
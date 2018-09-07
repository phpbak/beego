<!--创建text文件-->
<div style="width:100%;height:50px;">
        <button onclick="createtxt()" class="btn btn-primary">创建text文件</button>
        <label id="txtname"></label>

     </div>
      <!--写入文件-->
      <div style="width:100%;height:300px;margin-top:20px;">
          <label>输入内容:</label>
          <textarea id="writeinfo" style="width:50%;height:200px;" class="form-control"></textarea>
          <button value="写入内容" onclick="txtwrite()" class="btn btn-primary" style="margin-top:20px;">写入内容</button>
      </div>
      
      <!--读取文件内容部分-->
      <div style="width:100%;height:300px;">
          <button value="读取内容" onclick="txtread()" class="btn btn-primary" style="margin-bottom:20px;">读取内容</button>           
          <textarea id="readinfo" style="width:50%;height:200px;" class="form-control" ></textarea>
      </div>
          
      <button onclick="deletetxt()" class="btn btn-primary">删除text文件</button>

      <div id="showinfo"></div>
      <!--JS部分-->
        <script type="text/javascript">
         
          //创建text文件
          function createtxt(){
            $.ajax({
               type:'post',
               url:'{{urlfor "FileController.Create"}}',
               data:{},
               success:function(result){
                   console.log('获取的数据')
                   console.log(result)
                 $("#txtname").html(result.data);
               }
            })
 
          }
          //写入文件的内容
          function txtwrite(){
            $.ajax({
               type:'post',
               url:'{{urlfor "FileController.FileWrite"}}',
               data:{
                   "info":$("#writeinfo").val(),
                   "path":$("#txtname").html()
               },
               success:function(result){
                   console.log('获取的数据')
                   console.log(result)
                 $("#showinfo").html(result.data);
               }
            })
          }
          //读取文件的内容
          function txtread(){
            $.ajax({
               type:'post',
               url:'{{urlfor "FileController.FileRead"}}',
               data:{
                 "path":$("#txtname").html()
               },
               success:function(result){
                  console.log('获取的数据')
                  console.log(result)
                  $("#readinfo").html(result.data);
               }
            })
          }
           
           //删除text文件
           function deletetxt(){
            $.ajax({
               type:'post',
               url:'{{urlfor "FileController.FileDelete"}}',
               data:{
                   "path": $("#txtname").html()
               },
               success:function(result){
                   console.log('获取的数据')
                   console.log(result)
                   if(result.data){
                      $("#txtname").html('');
                      alert("删除成功");
                   }else{
                      alert("删除失败");
                   }                
               }
            })
 
          }
        </script>
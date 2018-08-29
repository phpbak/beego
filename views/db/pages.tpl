<style>
          .am-cf{
            height: 50px;         
            margin-top: 30px;
            line-height: 50px;
            text-align: center;
            vertical-align: middle;
            margin-left: 40%;
          }
          .am-fr{
              float: left;
              line-height: 50px;
              text-align: center;
              vertical-align: middle;
              height: 50px;
              margin-top: -15px;
          }
          .am-pagination{
            list-style:none;
            height: 50px;   
            line-height: 50px;
            text-align: center;
            vertical-align: middle;
          }
          .am-pagination li{
            float:left;
            margin-left: 10px;
          }
          .am-pagination li a{
            text-decoration:none;
          }
          .am-jl{
              float: left;
              margin-left: 20px;
          }
          .am-active{
              color: #f00;
          }
        </style>          
<table class="table table-striped table-hover table-bordered ">
<thead>
<th style="text-align: center">ID</th>
<th style="text-align: center">名称</th>
<th style="text-align: center">昵称</th>
<th style="text-align: center">Email</th>

<th style="text-align: center">手机号</th>
<th style="text-align: center">操作</th>
</thead>

<tbody id="sortable">
{{range $index, $item := .datas}}
<tr class="sort-item"  id="module_{{$item.Id}}" value="{{$item.id}}">
  <td style="text-align: center;width: 150px;"><span class="label label-default" >{{$item.id}}</span></td>
  <td style="text-align: center;width: 240px;" ><strong>{{$item.login_name}}</strong></td>
  <td style="text-align: center;width: 240px;" ><strong>{{$item.real_name}}</strong></td>
  <td style="text-align: center;width: 240px;" ><strong>{{$item.email}}</strong></td>
  <td style="text-align: center;width: 240px;" ><strong>{{$item.phone}}</strong></td>
  <td style="text-align: center;width: 150px;">
    <a href="{{urlfor "DbController.Edit"}}?Id={{$item.id}}" class="label label-info" title="修改" >修改</a>
    <a href="{{urlfor "DbController.Del"}}?Id={{$item.id}}" class="label label-info" title="删除">删除</a>
  </td>
{{end}}           
</tbody>
</table>
<!--分页开始--> 
<div class="am-cf">    
<div class="am-fr">
  <ul class="am-pagination">
  <li class=""><a href="{{urlfor "DbController.Pages"}}?page={{.paginator.firstpage}}">«</a></li>
  {{range $index,$page := .paginator.pages}}
    <li  {{if eq $.paginator.currpage $page }}class="am-active"{{end}}>
        <a href="{{urlfor "DbController.Pages"}}?page={{$page}}">{{$page}}</a></li>
  {{end}}
    <li><a href="{{urlfor "DbController.Pages"}}?page={{.paginator.lastpage}}">»</a></li>
  </ul>
</div>
<div class="am-jl">
共{{.totals}}条记录 共记{{.paginator.totalpages}} 页 当前{{.paginator.currpage}}页
</div>   
</div>
<!--分页结束--> 
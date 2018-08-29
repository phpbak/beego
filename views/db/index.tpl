<a href="{{urlfor "DbController.Add"}}">添加</a><br/>
<hr>
共{{.count}}条记录<br/>
<table border="1">
{{range $k, $v := .list}}
<tr>
	<td>{{$k}}</td>
<td>{{$v.id}}</td>
<td>{{$v.login_name}}</td>
<td>{{$v.real_name}}</td>
<td>{{$v.phone}}</td>
<td>{{$v.email}}</td>
<td>{{$v.role_ids}}</td>
<td><a href="{{urlfor "DbController.Edit"}}?id={{$v.id}}">修改</a></td>
<td><a href="{{urlfor "DbController.Del"}}?id={{$v.id}}">删除</a></td>
</tr>
{{end}}
</table>
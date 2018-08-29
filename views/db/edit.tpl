<form class="layui-form" action="{{urlfor "DbController.EditSave"}}" method="post" >
登录账号:<input name="login_name" value="{{.admin.login_name}}" /><br/>
真实姓名:<input name="real_name" value="{{.admin.real_name}}"/><br/>
手机号码:<input name="phone" value="{{.admin.phone}}"/><br/>
电子邮箱:<input name="email" value="{{.admin.email}}"/><br/>
选择角色:<input name="roleids" value="{{.admin.role_ids}}"/><br/>
ID:<input name="id" value="{{.admin.id}}"/><br/>
<input type="submit" name="提交"/>
</form>
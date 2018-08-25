<form method="post" action="{{urlfor "LoginController.Login"}}" >
username:<input name="uid" value="{{.uid}}" /><br/>
password:<input name="pwd"  value="{{.pwd}}" /><br/>
<input type="submit" value="submit" />
</form>
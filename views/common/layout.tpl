<!DOCTYPE html>
<html>
<head>
    <title>布局页面--layout.tpl</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <link rel="stylesheet" href="/static/bootstrap/css/bootstrap.min.css"/>
    <link rel="stylesheet" href="/static/bootstrap/css/bootstrap-theme.min.css"/>
    <script type="text/javascript" src="/static/js/jquery-2.1.1.min.js"></script> 
    <script type="text/javascript" src="/static/bootstrap/js/bootstrap.min.js"></script> 
    {{.HtmlHead}}
</head>
<body> 
    {{template "common/common_link.tpl" .}}
    <div class="container">
        {{.LayoutContent}}
    </div>
    <div class="sidebar">
        {{.SideBar}}
    </div>
    {{.Scripts}}

</body>
</html>
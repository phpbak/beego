<!DOCTYPE html>
<html>
<head>
    <title>{{.pageTitle}}--使用了布局页面(view/common/layouts.tpl)</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <script type="text/javascript" src="http://lib.sinaapp.com/js/jquery/2.0.2/jquery-2.0.2.min.js"></script> 
</head>
<body> 
    {{template "common/common_link.tpl" .}}
    <div class="container">
        {{.LayoutContent}}
    </div>

</body>
</html>
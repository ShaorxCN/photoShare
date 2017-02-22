<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="_xsrf" content="{{.xsrf_token}}" />
<style>
</style>
{{template "prepare/head.tpl"}}
 <script src="static/js/b.js" type="text/javascript"></script>
</head>
<body>
<div class="row">
  <div class="col-md-6">
    <div class="panel panel-default">
      <div class="panel-heading">注册</div>
      <div class="panel-body">
        <form   id="register-form" role="form">
        {{.xsrfdata}}
          <div class="form-group">
            <label for="username">用户名</label>
            <input type="text" id="username" name="username" class="form-control" placeholder="用户名" autofocus>
          </div>
          <div class="form-group">
            <label for="password">密码</label>
            <input type="password" id="password" name="password" class="form-control" placeholder="密码">
          </div>

          <div class="form-group">
          <label for="confirmpwd">确认密码</label>
          <input type="password" id="confirmpwd" name="confirmpwd" class="form-control" placeholder="确认密码">
          </div>
          <input type="submit" class="btn btn-sm btn-default" value="注册"> <a href="/login">去登录</a>
        </form>
      </div>
    </div>
  </div>
</div>
{{template "prepare/foot.tpl"}}
</body>
</html>
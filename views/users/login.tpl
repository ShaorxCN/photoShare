<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{config "String" "globaltitle" ""}}</title>
<style>
.form-signin .help-block{color:#a94442;}
</style>
{{template "prepare/head.tpl"}}
</head>
<body >
<div class="container">
  <form  id="login-form" style="width:30%;height:50%">
    <div >
      <h1>photoShare</h1>
      </div>
    <div>
      userName: <input type="text"   name="username" placeholder="请填写用户名,默认:shao" autofocus></br></br>
      password:&nbsp;&nbsp;<input type="password"   name="password" placeholder="请填写密码,默认:123456">
      </br>
      <button  type="submit" > 登录</button>
    </div>
  </form>
</div>
{{template "prepare/foot.tpl"}}
</body>
</html>

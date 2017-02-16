<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<style>
</style>
{{template "prepare/head.tpl"}}
</head>
<body style="background:yellow">

 
  
<div class="container" style="margin-top:5%">
 <div class="row col-md-offset-5 " >
      <h1>photoShare</h1>
  </div>
  <div class="row col-md-6 col-md-offset-3" >
  <form  id="login-form" role="form" style="width:50%;height:50%;margin:5% auto;">
    
    <div class="form-group">
      <label for="username">username</label> 
      <input type="text" class="form-control" id="username" name="username"  placeholder="请填写用户名，默认:shao" autofocus>
    </div>
    <div class="form-group">
      <label for="password">password</label>
      <input class="form-control" type="password" id="password" name="password" placeholder="请填写密码,默认:123456">
    </div>    
      <button class="btn btn-default col-md-offset-2" type="submit" >登录</button>
      <a class="btn btn-default" href="/register" style="margin-left:20%">注册</a>
  </form>
  </div>
</div>
{{template "prepare/foot.tpl"}}
</body>
</html>

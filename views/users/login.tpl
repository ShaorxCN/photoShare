<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<style>
</style>
{{template "prepare/head.tpl"}}
 <script src="static/js/b.js" type="text/javascript"></script>
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
      <input type="text" class="form-control" id="username" name="username"  placeholder="请填写用户名，默认:shao">
    </div>
    <div class="form-group">
      <label for="password">password</label>
      <input class="form-control" type="password" id="password" name="password" placeholder="请填写密码,默认:123456">
    </div>    
      <button class="btn btn-default col-md-offset-2" type="submit" >登录</button>

      <button class="btn btn-default col-md-offset-2" >注册</button>
  </form>
  </div>
</div>
{{template "prepare/foot.tpl"}}
</body>
</html>

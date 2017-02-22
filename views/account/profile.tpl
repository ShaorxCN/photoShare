<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="_xsrf" content="{{.xsrf_token}}">
<style type="text/css"></style>
{{template "prepare/head.tpl"}}
    <title>用户信息</title>
</head>
<body>

<div class="row">
    <div class="col-md-5">
    <div class="panel panel-default">
        <div class="panel-heading">用户信息</div>
        <div class="panel-body">
            <form id="profile-form" role="form">
                <input type="text" name="userId" id="userId" value="{{.response.profile.Id}}" style="display:none">
                <div class="form-group">
                    <label for="realname">真实姓名<font color="red">*</font></label>
                    <input type="text" name="realname" id="realname" class="form-control" placeholder="真实姓名" autofocus>
                </div>

                <div class="form-group">
                    <label for="sex">性别<font color="red">*</font></label><br>
                     
                    <input type="radio" name="sex" id="sex" value="1">男</label>

                    
                    <input type="radio" name="sex" id="sex2" value="0">女</label>
                </div>

                <div class="form-group input-append date">
                <label for="birth">出生日期<font color="red">*</font></label>
                <input type="text" name="birth" id="birth" class="form-date form-control">
                </div>

                <div class="form-group">
                    <label for="email">电子邮箱<font color="red">*</font></label>
                    <input type="text" name="email" id="email" class="form-control" placeholder="123@xx.com">
                </div>


                <div class="form-group">
                    <label for="phone">手机号码<font color="red">*</font></label>
                    <input type="text" name="phone" id="phone" class="form-control">
                </div>

                <div class="form-group">
                <label for="address">地址</label>
                <input type="text" name="address" ip="address" class="form-control">
                </div>

                <input type="submit" name="btn btn-sm btn-default" value="提交">
            </form>
        </div>
    </div>
    </div>
</div>
{{template "prepare/foot.tpl"}}
</body>
</html>
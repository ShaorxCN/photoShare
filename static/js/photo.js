$(function(){
    
    $('#login-form').validate({
        ignore:'',
        rules : {
            username:{ required: true},
            password:{required: true}
        },
        
        submitHandler:function(form) {
            var url = '/login';
            //var param = $("#login-form").serialize();
            var pwd = hex_md5(document.getElementById("password").value);
            var username = document.getElementById("username").value;
            var param = {"username":username,"password":pwd};

            
            $.ajax({
                url:url,
                type:'POST',
                dataType:'json',
                data:param,
                timeout:1000,
                success:function(data) {
                    dialogInfo(data.message)
                    var redirect = data.redirect
                    if (data.code){
                      setTimeout(function(){window.location.replace(redirect)}, 2000);
                    } else {
                       setTimeout(function(){ $('#dialogInfo').modal('hide'); }, 1000);
                    }
                },
                error:function(XMLHttpRequest, textStatus, errorThrown){
                     dialogInfo("服务器正忙，请稍后尝试")
                    setTimeout(function() {
                         $('#dialogInfo').modal('hide'); 
                    }, 1000);
                }
            });
        }
    });


    $('#register-form').validate({
        ignore:'',
        rules: {
            username:{required:true,maxlength:15},
            password:{required:true,maxlength:40,minlength:8},
            confirmpwd:{required:true,maxlength:40,minlength:8,equalTo:"#password"}
        },
        
        submitHandler:function(form){
            var url = '/register';
            var pwd = hex_md5(document.getElementById("password").value);
            var username = document.getElementById("username").value;
            var param = {"username":username,"password":pwd};
            $.ajax({
                url:url,
                type:'POST',
                dataType:'json',
                data:param,
                timeout:1000,
                success:function(data) {
                    dialogInfo(data.message)
                    var redirect = data.redirect
                    if (data.code) {
                      setTimeout(function(){window.location.replace(redirect)}, 2000);
                    } else {
                       setTimeout(function(){ $('#dialogInfo').modal('hide'); }, 1000);
                    }
                },
                error:function(XMLHttpRequest, textStatus, errorThrown){
                     dialogInfo("服务器正忙，请稍后尝试")
                    setTimeout(function() {
                         $('#dialogInfo').modal('hide'); 
                    }, 1000);
                }
            });
        }
    });
});


//引用日期控件
$('.form-date').datepicker({
    language:'zh-CN',
    weekStart:1,
    todayHighlight:true,
    todayBtn:"linked",
    format:"yyyy-mm-dd",
    minView:"month",
    autoclose:true
});

function dialogInfo(msg) {
    $('#dialogInfo').remove();
    var html = '';
    html = '<div class="modal fade" id="dialogInfo" tabindex="-1" role="dialog" aria-labelledby="dialogInfoTitle">';
    html += '<div class="modal-dialog" role="document">';
    html += '<div class="modal-content">';
    html += '<div class="modal-header">';
    html += '<button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>';
    html += '<h4 class="modal-title" id="dialogInfoTitle">photoShare提示</h4>';
    html += ' </div>';
    html += '<div class="modal-body">';
    html += '<p>'+msg+'</p>';
    html += '</div>';
    //html += '<div class="modal-footer">';
    //html += ' <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>';
    //html += ' <button type="button" class="btn btn-primary">Send message</button>';
    //html += '</div>';
    html += '</div>';
    html += '</div>';
    html += '</div>';
    $('body').append(html);
    $('#dialogInfo').modal('show')  
}

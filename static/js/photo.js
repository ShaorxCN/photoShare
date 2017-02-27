$(function(){
    
    var ajax = $.ajax;
    $.extend({
        ajax: function(url, options) {
            if (typeof url === 'object') {
                options = url;
                url = undefined;
            }
            options = options || {};
            url = options.url;
            var xsrftoken = $('meta[name=_xsrf]').attr('content');
            var headers = options.headers || {};
            var domain = document.domain.replace(/\./ig, '\\.');
            if (!/^(http:|https:).*/.test(url) || eval('/^(http:|https:)\\/\\/(.+\\.)*' + domain + '.*/').test(url)) {
                headers = $.extend(headers, {'X-Xsrftoken':xsrftoken});
            }
            options.headers = headers;
            return ajax(url, options);
        }
    });

    
    $('#login-form').validate({
        ignore:'',
        rules : {
            username:{ required: true},
            password:{required: true}
        },
        highlight : function(element) { 
                $(element).closest('.form-group').addClass('has-error'); 
            }, 
 
            success : function(label) { 
                label.closest('.form-group').removeClass('has-error'); 
                label.remove(); 
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
        highlight : function(element) { 
                $(element).closest('.form-group').addClass('has-error'); 
            }, 
 
            success : function(label) { 
                label.closest('.form-group').removeClass('has-error'); 
                label.remove(); 
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


    //客户信息
    $('#profile-form').validate({
        ignore:'',
        rules: {
            userId:{required:true},
            realname:{required:true,maxlength:15},
            sex:{required:true},
            birth:{required:true},
            email:{required:true,email:true},
            phone:{required:true,isMobile:true}
        },

        errorPlacement:function(error,element){
            $( element )
        .closest( "form" )
          .find( "label[for='" + element.attr( "id" ) + "']" )
            .append( error );
  },
        highlight : function(element) { 
                $(element).closest('.form-group').addClass('has-error'); 
            }, 
 
            success : function(label) { 
                label.closest('.form-group').removeClass('has-error'); 
                label.remove(); 
            },  

        submitHandler:function(form){
            var userId = document.getElementById("userId").value;
            var url = '/profile/'+userId;
            var url2 = '/user/'+userId;
            var param = $("#profile-form").serialize();
            $.ajax({
                url:url,
                type:'POST',
                dataType:'json',
                data:param,
                timeout:3000,
                success:function(data) {
                    
                    dialogInfo(data.message);
                    if (data.code) {
                      setTimeout(function(){window.location.replace(url2)}, 2000);
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


//todo
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


function reset(){
  document.getElementById("res").click();
}
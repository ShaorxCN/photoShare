jQuery.extend(jQuery.validator.messages,{
  required: "该字段不能为空",
  remote: "请修正该字段",
  email: "请输入正确格式的电子邮件",
  url: "请输入合法的网址",
  date: "请输入合法的日期",
  dateISO: "请输入合法的日期 (ISO).",
  number: "请输入合法的数字",
  digits: "只能输入整数",
  creditcard: "请输入合法的信用卡号",
  equalTo: "请再次输入相同的值",
  accept: "请输入拥有合法后缀名的字符串",
  maxlength: jQuery.validator.format("长度不能大于 {0}"),
  minlength: jQuery.validator.format("长度不能小于 {0} "),
  rangelength: jQuery.validator.format("请输入 一个长度介于 {0} 和 {1} 之间的字符串"),
  range: jQuery.validator.format("请输入一个介于 {0} 和 {1} 之间的值"),
  max: jQuery.validator.format("请输入一个最大为{0} 的值"),
  min: jQuery.validator.format("请输入一个最小为{0} 的值")
});

//自定义手机校验
jQuery.validator.addMethod("isMobile",function(value,element){
  var length = value.length;
  var mobile = /^(13[0-9]{9})|(18[0-9]{9})|(17[0-9]{9})|(15[0-9]{9})|(14[0-9]{9})$/;
  return this.optional(element) || (length==11 &&mobile.test(value));
},"请输入正确的手机号,如有疑问请联系管理员");
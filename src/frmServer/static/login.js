// 生成随机数
var ran = "dfaedefd";
var makeRan = function(){
	var d = new Date();
	ran = $.phpjs.sha1(d.getTime());
};
makeRan()
// 每100秒执行一次makeRan函数
window.setInterval(makeRan,100000);

$(document).ready(function(){
	$("#log-form .submit").click(function(){
		var username = $("#log-form .username").val()
		var password = $("#log-form .password").val()
		var check = ran
		password = $.phpjs.sha1(password)
		password = password + check
		password = $.phpjs.sha1(password)
		$.post("login", {username: username, password: password, check: check}).done(function(data){
			if (data == "ok") {
				window.location.href='/'
			}else{
				alert(data)
			}
		})
	});
});

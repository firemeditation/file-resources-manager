var iNowUserInfo = {
	UPowerLevel : {
		0 : "没有权限，Level 0" ,
		1 : "普通读权限，Level 1",
		2 : "普通写权限，Level 2",
		3 : "增强权限，可以访问受保护的资源，Level 3",
		4 : "管理权限，具有本机构的管理权限，Level 4",
		5 : "全局权限，具有全系统的管理权限，Level 5"
	},
	LoadShowInfo : function(){
		$("#main-box #now-user-info #iNowUserInfo .user_title_name span").text(login_user.Name);
		$("#main-box #now-user-info #iNowUserInfo .user_all_info .belong-unit").text(login_user.UnitName);
		$("#main-box #now-user-info #iNowUserInfo .user_all_info .belong-group").text(login_user.GroupName);
		if (login_user.UPower.user.user){
			$("#main-box #now-user-info #iNowUserInfo .user-power-box .power-user").text(this.UPowerLevel[login_user.UPower.user.user]);
		}else{
			$("#main-box #now-user-info #iNowUserInfo .user-power-box .power-user").text(this.UPowerLevel[0]);
		}
		if (login_user.UPower.user.unit){
			$("#main-box #now-user-info #iNowUserInfo .user-power-box .power-units").text(this.UPowerLevel[login_user.UPower.user.unit]);
		}else{
			$("#main-box #now-user-info #iNowUserInfo .user-power-box .power-units").text(this.UPowerLevel[0]);
		}
		if (login_user.UPower.user.group){
			$("#main-box #now-user-info #iNowUserInfo .user-power-box .power-usergroup").text(this.UPowerLevel[login_user.UPower.user.group]);
		}else{
			$("#main-box #now-user-info #iNowUserInfo .user-power-box .power-usergroup").text(this.UPowerLevel[0]);
		}
		if (login_user.UPower.resource.origin){
			$("#main-box #now-user-info #iNowUserInfo .user-power-box .power-resource").text(this.UPowerLevel[login_user.UPower.resource.origin]);
		}else{
			$("#main-box #now-user-info #iNowUserInfo .user-power-box .power-resource").text(this.UPowerLevel[0]);
		}
	},
	doSubmit : function(){
		var ckArray = [0,0,0];
		ckArray[0] = $.RequestProcess.Password('#change-password .old_password',0,1);
		ckArray[1] = $.RequestProcess.PasswordOne('#change-password .new_password',0,6);
        ckArray[2] = $.RequestProcess.PasswordTwo('#change-password .new_password','#change-password .new_password2',0);
        if($.RequestProcess.ckAllOne(ckArray)==0){ alert("输入错误") ;}
        oldpassword = $('#change-password .old_password input').val();
        oldpassword = $.phpjs.sha1(oldpassword);
        newpassword = $('#change-password .new_password input').val();
        newpassword = $.phpjs.sha1(newpassword);
        $("#nowloadbox").fadeIn(200);
        $.post("webInterface?type=change-password",{old: oldpassword, news: newpassword}).done(function(data){
			theJson = $.parseJSON(data);
			if(theJson.err){
				alert(theJson.err); processServerError(theJson.err);
			}else{
				alert("修改密码成功");
			};
			$("#nowloadbox").fadeOut(200);
		});
	},
};

$('#change-password .old_password input').keyup(function(){ $.RequestProcess.Password('#change-password .old_password',0,1); });
$('#change-password .new_password input').keyup(function(){ $.RequestProcess.PasswordOne('#change-password .new_password',0,6); });
$('#change-password .new_password2 input').keyup(function(){ $.RequestProcess.PasswordTwo('#change-password .new_password','#change-password .new_password2',0); });
$("#change-password .submit input").click(function(){iNowUserInfo.doSubmit()});

iNowUserInfo.LoadShowInfo();

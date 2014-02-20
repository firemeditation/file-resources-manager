var local_client_port = 9876;

var search_word = ""; //正在搜索的关键词
var search_type = "rg";  //搜索类型，可选rg、rf、rt

var login_user;  //登录的用户信息

// 隐藏一切需要隐藏的
var hideAll = function(){
	$("#main-box #help-box").hide();
	$("#main-box #resource-list").hide();
	$("#main-box #resource-add-box").hide();
	$("#main-box #now-user-info").hide();
	$("html, body").animate({scrollTop:0}, 'slow')
};

var processServerError = function(err){
	if (err == "用户超时" || err == "用户不存在" || err == "不是正确的接口请求"){
		window.location.href='login'
	}
}

// begin 显示基本的查看资源的资源列表
var showTheBasicResourceList = function(){
	$("#nowloadbox").fadeIn(200);
	$("#main-box #resource-list").load("static/iResourceList.htm",function(){
		$.getScript("static/iResourceList.js");
	});
	$("#main-box #resource-list").fadeIn();
}
// end 显示基本的查看资源的资源列表

// 清除搜索带来的影响
var searchClean = function(){
	search_word = "";
	$("#top-kongzhi .soutext").val("");
}

// doSearch 执行搜索框的搜索业务
var doSearch = function(){
	var search_change_port_reg = new RegExp("^client.port:([0-9]+)$");
	var search_text = $("#top-kongzhi .soutext").val();
	search_text = cTrim(search_text,0)
	if (mbStringLength(search_text) == 0){
		return
	}
	search_text = inputSafe.CleanAll(search_text)
	if (search_text.match(search_change_port_reg)){
		var exec = search_change_port_reg.exec(search_text);
		var newport = exec[1];
		if (newport < 1 || newport > 65535){
			alert("请输入正确的端口号。");
		}else{
			local_client_port = newport;
			$("#top-status-area .local_client_port").text(local_client_port);
			alert("设置本地客户端端口成功。");
			$("#top-kongzhi .soutext").val("");
			$.cookie('local_client_port',local_client_port,{expires:365, path:'/'});
		};
	}else{
		//执行其他搜索
		search_word = search_text;
		hideAll();
		showTheBasicResourceList();
	};
};

// checkClientStatus 检查客户端状态
var checkClientStatus = function(){
	$.getJSON("http://127.0.0.1:"+local_client_port+"/checkLink?callback=?",function(r){
		if(r['client'] == 'all'){
			$("#top-status-area .client_status").text("客户端连接正常");
		}else if ((r['client'] == 'less'){
			$("#top-status-area .client_status").text("客户端无法访问服务器");
		}
	}).fail(function(){
		$("#top-status-area .client_status").text("客户端没有启动");
	});
}
// 每30秒执行一次checkClientStatus函数
window.setInterval(checkClientStatus,30000);


// 维持登录状态心跳
var updateLive = function(){
	$.get("updateLive",function(data){
		theJson = $.parseJSON(data);
		if(theJson.err){
			alert(theJson.err);
			processServerError(theJson.err);
		};
	});
}
// 每10分钟执行一次updateLive函数
window.setInterval(updateLive,600000);


//载入后直接获取用户信息
var onloadGetUserinfo = function(){
	$.get("webInterface?type=get-basic-user-info",function(data){
		//alert(data);
		var json = $.parseJSON(data);
		login_user = json;
		//alert(login_user.UPower.resource.origin);
		if(json.err){
			$("#top-kongzhi .show_hide_userinfo .user-info-self").text("错误：" + json.err)
		}else{
			$("#top-kongzhi .show_hide_userinfo .user-info-self").html("用户名：" + json.Name + "&nbsp;&nbsp;所属机构：" + json.UnitName + "&nbsp;&nbsp;所属组：" + json.GroupName)
		}
	});
};

$(document).ready(function(){
	updateLive();
	
	onloadGetUserinfo();
	//检查端口cookie
	var local_client_port_cookie = $.cookie('local_client_port');
	if (local_client_port_cookie != undefined) {
		local_client_port = local_client_port_cookie;
		$("#top-status-area .local_client_port").text(local_client_port);
	};
	
	checkClientStatus();  //检查客户端状态
	$("#top-status-area .local_client_port").text(local_client_port);
	$("#top-kongzhi .userinfo").mouseover(function(){
		$("#top-kongzhi .show_hide_userinfo").show(100);
		$("#top-kongzhi .sousuo").hide(100);
	}).mouseout(function(){
		$("#top-kongzhi .show_hide_userinfo").hide(100);
		$("#top-kongzhi .sousuo").show(100);
	});
		
	//begin 点击搜索动作
	$("#top-kongzhi .soubutton").click(function(){
		doSearch();
	});
	$("#top-kongzhi .soutext").keydown(function(e){
		if(e.keyCode == 13){
			doSearch();
		};
	});
	//end 点击搜索动作
	
	//begin 点击客户端状态
	$("#top-status-area .client_status").click(function(){
		$("#allwhite").show();
		$("#showBackupRecord").show();
		$("#showBackupRecord .inside-box").html("");
		$.getJSON("http://127.0.0.1:"+local_client_port+"/getBackupRecord?userid="+login_user.HashId+"&callback=?",function(r){
			if(r["err"]){
				$("#showBackupRecord .inside-box").prepend(r["err"])
				return
			}
			$.each(r, function(key,value){
				var theTime = key * 1000;
				var timedate = new Date(theTime);
				var theTime = timedate.formatDate("yyyy年MM月dd日 hh:mm:ss");
				$.each(value, function(lkey, lval){
					$("#showBackupRecord .inside-box").prepend("<p>" + theTime + "：" + lval + "</p>")
				})
			})
		})
	})
	$("#showBackupRecord .close-box").click(function(){
		$("#allwhite").hide();
		$("#showBackupRecord").hide();
	});
	//end 点击客户端状态
	
	// begin 点击查看当前用户的信息
	$("#top-kongzhi .userinfo").click(function(){
		hideAll();
		$("#nowloadbox").fadeIn(200);
		searchClean();
		$("#main-box #now-user-info").load("static/iNowUserInfo.htm", function(){
			$.getScript("static/iNowUserInfo.js").done(function(){
				$("#main-box #now-user-info").fadeIn();
				$("#nowloadbox").fadeOut(200);
			});
		});
	});
	// end 点击查看当前用户的信息
	
	//begin 点击查看资源
	$("#top-kongzhi .chakan").click(function(){
		if(login_user.UPower.resource.origin < 1){ return }
		hideAll();
		searchClean();
		showTheBasicResourceList();
	});
	//end 点击查看资源
	
	$("#top-kongzhi .usehelp").click(function(){
		hideAll();
		$("#top-kongzhi .soutext").val("");
		$("#main-box #help-box").fadeIn();
	});
	
	//begin 点击新建资源
	$("#top-kongzhi .xinjian").click(function(){
		if(login_user.UPower.resource.origin < 2){ return }
		hideAll();
		$("#nowloadbox").fadeIn(200);
		searchClean();
		$("#main-box #resource-add-box-true").load("static/iAddResource.htm", function(){
			$.getScript("static/iAddResource.js").done(function(){
				$("#main-box #resource-add-box").fadeIn();
				$("#nowloadbox").fadeOut(200);
			});
		});
	});
	//end 点击新建资源
});

var logoutSystem = function(){
	$.get("logout",function(){
		window.location.href='login';
	});
};

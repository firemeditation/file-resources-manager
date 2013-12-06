var local_client_port = 9876;

// 隐藏一切需要隐藏的
var hideAll = function(){
	$("#main-box #help-box").hide();
	$("#main-box #resource-list").hide();
};

// doSearch 执行搜索框的搜索业务
var doSearch = function(){
	var search_change_port_reg = new RegExp("^client.port:([0-9]+)$");
	var search_text = $("#top-kongzhi .soutext").val();
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
	};
};

// checkClientStatus 检查客户端状态
var checkClientStatus = function(){
	$.getJSON("http://127.0.0.1/frm/index.php?checkClient=yes&callback=?",function(r){
		if(r['client'] == 'yes'){
			$("#top-status-area .client_status").text("连接正常");
		}else{
			$("#top-status-area .client_status").text("连接不正常");
		}
	}).fail(function(){
		$("#top-status-area .client_status").text("没有连接");
	});
}
// 每30秒执行一次checkClientStatus函数
window.setInterval(checkClientStatus,30000);


// 维持登录状态心跳
var updateLive = function(){
	$.get("updateLive")
}
// 每10分钟执行一次updateLive函数
window.setInterval(updateLive,600000);

$(document).ready(function(){
	
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
	
	$("#top-kongzhi .userinfo").click(function(){
		hideAll();
	});
	$("#top-kongzhi .chakan").click(function(){
		hideAll();
		$("#main-box #resource-list").fadeIn();
	});
	$("#top-kongzhi .usehelp").click(function(){
		hideAll();
		$("#main-box #help-box").fadeIn();
	});
	
	//begin 图书列表的点击动作
	var closeOneBookAll = function(theone,type){
		if(type != "resource-all-info"){
			theone.children(".resource-all-info").hide(100);
			theone.children(".resource-all-info").attr('showit',"no");
		}
		if(type != "resource-all-file"){
			theone.children(".resource-all-file").hide(100);
			theone.children(".resource-all-file").attr('showit',"no");
		}
	}
	$("#resource-main-list .the-resource-name").click(function(){
		closeOneBookAll($(this).parent().parent(),"resource-all-info");
		var theone = $(this).parent().parent().children(".resource-all-info");
		if(theone.attr('showit') == "no"){
			theone.show(100);
			theone.attr('showit',"yes");
		}else{
			theone.hide(100);
			theone.attr('showit',"no");
		}
	});
	$("#resource-main-list .liulan").click(function(){
		closeOneBookAll($(this).parent().parent().parent(),"resource-all-file");
		var theone = $(this).parents(".one-resource-main").children(".resource-all-file");
		if(theone.attr('showit') == "no"){
			theone.show(100);
			theone.attr('showit',"yes");
		}else{
			theone.hide(100);
			theone.attr('showit',"no");
		}
	});
	//end 图书列表的点击动作
});
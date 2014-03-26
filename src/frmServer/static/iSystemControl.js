var iSystemControl = {
	
	//修改系统管理的高度
	changeSystemControlHeight : function(){
		var windowHeight = $(window).height();
		var fullboxHeight = windowHeight - 60;
		var upponHeight = fullboxHeight - 31 - 40 - 5;
		$("#system-control-true").height(fullboxHeight);
		$("#system-control-true .body-right").height(upponHeight);
	},
	
	//关闭管理界面
	Close : function(){
		$("#main-box #system-control").hide();
		$("#allwhite").hide();
	},
	
	//处理body-left的isnow=true的显示
	blIsNow : function(){
		$("#system-control-true .body .body-left .menu-list").each(function(index){
			if($(this).attr("isnow") == "true"){
				$(this).children("span").css("border-bottom","dashed 1px #ff0000");
				//$(this).children("span").addClass("is_now_true");
			}else{
				//$(this).children("span").addClass("is_now_false");
				$(this).children("span").css("border","0");
			}
		});
	},
};

iSystemControl.changeSystemControlHeight(); 
$(window).resize(function(){ iSystemControl.changeSystemControlHeight(); });
iSystemControl.blIsNow();

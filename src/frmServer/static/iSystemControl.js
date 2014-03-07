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
};

iSystemControl.changeSystemControlHeight(); 
$(window).resize(function(){ iSystemControl.changeSystemControlHeight(); });

var iSystemControl = {
	
	//修改系统管理的高度
	changeSystemControlHeight : function(){
		var windowHeight = $(window).height();
		var fullboxHeight = windowHeight - 120;
		var upponHeight = fullboxHeight - 60;
		$("#system-control-true").height(fullboxHeight);
		//$("#system-control-true .uppon-info-show").height(upponHeight);
	},
};

iSystemControl.changeSystemControlHeight(); 
$(window).resize(function(){ iSystemControl.changeSystemControlHeight(); });

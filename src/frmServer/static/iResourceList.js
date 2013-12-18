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
var resourceNameClick = function(self){
	closeOneBookAll($(self).parent().parent(),"resource-all-info");
	var theone = $(self).parent().parent().children(".resource-all-info");
	if(theone.attr('showit') == "no"){
		theone.show(100);
		theone.attr('showit',"yes");
	}else{
		theone.hide(100);
		theone.attr('showit',"no");
	}
}

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

var resourceLiulanClick = function(self){
	closeOneBookAll($(self).parent().parent().parent(),"resource-all-file");
	var theone = $(self).parents(".one-resource-main").children(".resource-all-file");
	if(theone.attr('showit') == "no"){
		theone.show(100);
		theone.attr('showit',"yes");
	}else{
		theone.hide(100);
		theone.attr('showit',"no");
	}
}
//end 图书列表的点击动作

$.get("webInterface?type=resource-list",function(data){
	var json = $.parseJSON(data)
	$("#resource-list .allListBookCount").text(json.Count)
	var i = 0
	$(json.List).each(function(){
		//var li = $.parseJSON(this.MetaData)
		var onebook = '<div class="one-resource-main" hashid="'+this.HashId+'">\
		<div class="one-resource-total-info">\
			<div class="the-resource-name" onclick=resourceNameClick(this)>'+this.Name+'</div>\
			<div class="the-resource-edits">\
				<div class="shanchu resource-edits-1">删</div>\
				<div class="shangchuan resource-edits-1">上</div>\
				<div class="bianji resource-edits-1">编</div>\
				<div class="liulan resource-edits-1" onclick=resourceLiulanClick(this)>浏</div>\
				<div class="xiazai resource-edits-1">下</div>\
			</div>\
		</div>\
		<div class="resource-all-info" showit="no">'+this.Info+'<br>'+json.Meta[i].Author+'</div>\
		<div class="resource-all-file" showit="no">\
			<div class="now-dir">··/所在文件夹（如果上级没有则斜线后没内容）</div>\
			<ul class="file-list">\
				<li hashid="dfa34edfgserasgdd45yuhgf" filetype="d/f/t"><span class="file-list-type">D</span><span class="file-list-name">文件名或文件夹名</span><span class="xiazai2 file-list-opt">下</span><span class="bianji2 file-list-opt">编</span><span class="shanchu2 file-list-opt">删</span></li>\
				<li hashid="dfa34edfgserasgdd45yuhgf" filetype="d/f/t"><span class="file-list-type">F</span><span class="file-list-name">文件名或文件夹名</span><span class="xiazai2 file-list-opt">下</span><span class="bianji2 file-list-opt">编</span><span class="shanchu2 file-list-opt">删</span></li>\
				<li hashid="dfa34edfgserasgdd45yuhgf" filetype="d/f/t"><span class="file-list-type">T</span><span class="file-list-name">文件名或文件夹名</span><span class="xiazai2 file-list-opt">下</span><span class="bianji2 file-list-opt">编</span><span class="shanchu2 file-list-opt">删</span></li>\
			</ul>\
		</div>\
	</div>'
		$("#resource-main-list").append(onebook)
		i++
	})
})

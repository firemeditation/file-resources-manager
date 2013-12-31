var iResourceList_from = 0;
var iResourceList_limit = 10;
var iResourceList_count = 0;

//begin 图书列表的点击动作
var closeOneBookAll = function(theone,type){
	if(type != "resource-all-info"){
		theone.children(".resource-all-info").hide(100);
		theone.children(".resource-all-info").attr('showit',"no");
	};
	if(type != "resource-all-file"){
		theone.children(".resource-all-file").hide(100);
		theone.children(".resource-all-file").attr('showit',"no");
	};
};
$("#resource-main-list .the-resource-name").click(function(){
	closeOneBookAll($(this).parent().parent(),"resource-all-info");
	var theone = $(this).parent().parent().children(".resource-all-info");
	if(theone.attr('showit') == "no"){
		theone.show(100);
		theone.attr('showit',"yes");
	}else{
		theone.hide(100);
		theone.attr('showit',"no");
	};
});
var resourceNameClick = function(self){
	closeOneBookAll($(self).parent().parent(),"resource-all-info");
	var theone = $(self).parent().parent().children(".resource-all-info");
	if(theone.attr('showit') == "no"){
		$("#resource-main-list .resource-all-info").each(function(){$(this).hide(100); $(this).attr('showit','no')});
		$("#resource-main-list .resource-all-file").each(function(){$(this).hide(100); $(this).attr('showit','no')});
		theone.show(100);
		theone.attr('showit',"yes");
	}else{
		theone.hide(100);
		theone.attr('showit',"no");
	};
};

$("#resource-main-list .liulan").click(function(){
	closeOneBookAll($(this).parent().parent().parent(),"resource-all-file");
	var theone = $(this).parents(".one-resource-main").children(".resource-all-file");
	if(theone.attr('showit') == "no"){
		theone.show(100);
		theone.attr('showit',"yes");
	}else{
		theone.hide(100);
		theone.attr('showit',"no");
	};
});

var resourceLiulanClick = function(self){
	closeOneBookAll($(self).parent().parent().parent(),"resource-all-file");
	var theone = $(self).parents(".one-resource-main").children(".resource-all-file");
	if(theone.attr('showit') == "no"){
		$("#resource-main-list .resource-all-info").each(function(){$(this).hide(100); $(this).attr('showit','no')});
		$("#resource-main-list .resource-all-file").each(function(){$(this).hide(100); $(this).attr('showit','no')});
		theone.show(100);
		theone.attr('showit',"yes");
	}else{
		theone.hide(100);
		theone.attr('showit',"no");
	};
};
//end 图书列表的点击动作

var lastOtime = function(utime){
	var theTime = utime * 1000;
	var timedate = new Date(theTime);
	var theTime = timedate.formatDate("yyyy年MM月dd日 hh:mm:ss");
	return theTime;
};
var getResourceListFromServer = function(){
	$("#nowloadbox").fadeIn(200);
	$("#resource-main-list").html("")
	$.get("webInterface?type=resource-list&from="+iResourceList_from+"&limit="+iResourceList_limit , function(data){
		var json = $.parseJSON(data);
		$("#resource-list .allListBookCount").text(json.Count);
		iResourceList_count = json.Count;
		if (iResourceList_from == 0){
			$("#next-and-prev .prev").hide();
		}else{
			$("#next-and-prev .prev").show();
		};
		if (iResourceList_from + iResourceList_limit >= iResourceList_count){
			$("#next-and-prev .next").hide();
		}else{
			$("#next-and-prev .next").show();
		};
		var i = 0;
		
		md_converter = new Markdown.Converter();
		
		$(json.List).each(function(){
			//var li = $.parseJSON(this.MetaData);
			var ptime = lastOtime(this.Table.Btime);
			
			var md_c = md_converter.makeHtml(this.Table.Info);
			
			var onebook = '<div class="one-resource-main" hashid="'+this.Table.HashId+'">\
			<div class="one-resource-total-info">\
				<div class="the-resource-name" onclick=resourceNameClick(this)>'+this.Table.Name+'</div>\
				<div class="the-resource-edits">\
					<div class="shanchu resource-edits-1">删</div>\
					<div class="shangchuan resource-edits-1">上</div>\
					<div class="bianji resource-edits-1">编</div>\
					<div class="liulan resource-edits-1" onclick=resourceLiulanClick(this)>浏</div>\
					<div class="xiazai resource-edits-1">下</div>\
				</div>\
			</div>\
			<div class="resource-all-info" showit="no"><p>类型：'+this.RSR.RtName+'&nbsp;&nbsp;最后操作人：'+this.RSR.UsersName+'&nbsp;&nbsp;最后操作时间：'+ptime+'</p>\
			<p>作者：'+this.MD.Author+'&nbsp;&nbsp;编辑：'+this.MD.Editor+'&nbsp;&nbsp;ISBN/ISSN：'+this.MD.ISBN+'&nbsp;&nbsp;</p>\
			<p>简介：</p>\
			<div class="markdown">'+md_c+'</div></div>\
			<div class="resource-all-file" showit="no">\
				<div class="now-dir">··/所在文件夹（如果上级没有则斜线后没内容）</div>\
				<ul class="file-list">\
					<li hashid="dfa34edfgserasgdd45yuhgf" filetype="d/f/t"><span class="file-list-type">D</span><span class="file-list-name">文件名或文件夹名</span><span class="xiazai2 file-list-opt">下</span><span class="bianji2 file-list-opt">编</span><span class="shanchu2 file-list-opt">删</span></li>\
					<li hashid="dfa34edfgserasgdd45yuhgf" filetype="d/f/t"><span class="file-list-type">F</span><span class="file-list-name">文件名或文件夹名</span><span class="xiazai2 file-list-opt">下</span><span class="bianji2 file-list-opt">编</span><span class="shanchu2 file-list-opt">删</span></li>\
					<li hashid="dfa34edfgserasgdd45yuhgf" filetype="d/f/t"><span class="file-list-type">T</span><span class="file-list-name">文件名或文件夹名</span><span class="xiazai2 file-list-opt">下</span><span class="bianji2 file-list-opt">编</span><span class="shanchu2 file-list-opt">删</span></li>\
				</ul>\
			</div>\
		</div>';
			$("#resource-main-list").append(onebook);
			i++;
		});
		$("#nowloadbox").fadeOut(200);
	});
};

getResourceListFromServer();

$("#next-and-prev .next").click(function(){
	iResourceList_from = iResourceList_from + iResourceList_limit;
	getResourceListFromServer();
});

$("#next-and-prev .prev").click(function(){
	iResourceList_from = iResourceList_from - iResourceList_limit;
	getResourceListFromServer();
});

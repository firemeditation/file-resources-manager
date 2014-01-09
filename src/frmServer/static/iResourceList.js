var iResourceList_from = 0;
var iResourceList_limit = 10;
var iResourceList_count = 0;

var theBigJSON;

//修改弹出的每本书的详细信息的框的高度
var changeFullResouceBoxHeight = function(){
	var windowHeight = $(window).height();
	var fullboxHeight = windowHeight - 120;
	var upponHeight = fullboxHeight - 60;
	$("#resource-one-full").height(fullboxHeight);
	$("#resource-one-full .uppon-info-show").height(upponHeight);
}

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
	var allinfo = $(self).parent().parent().children(".resource-all-info").html();
	var hashid = $(self).parent().parent().attr("hashid");
	var bookname = $(self).parent().parent().children(".one-resource-total-info").children(".the-resource-name").text();
	$('#resource-one-full').attr("hashid",hashid);
	$('#resource-one-full .uppon-info-show .resource-all-info').html(allinfo);
	$('#resource-one-full .the-resource-name').text(bookname);
	changeFullResouceBoxHeight();
	$("#allwhite").show();
	$('#resource-one-full').show();
	
};

$(window).resize(function(){ changeFullResouceBoxHeight(); });

var resourceCloseNow = function(self){
	$("#allwhite").hide();
	$('#resource-one-full').hide();
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
	};
});

var showBigJsonLevel = function(json){
	//$("#resource-one-full .resource-all-file .now-dir .true-now").text(path);
	//$("#resource-one-full .resource-all-file .file-list").html("");
	var onefile = ''
	$.each(json, function(name, value){
		if (value.IsDir == false){
			onefile += '<li hashid="'+value.HashId+'" filetype="f"><span class="file-list-type">F</span><span class="file-list-name">'+value.Name+'</span><span class="xiazai2 file-list-opt">下</span><span class="bianji2 file-list-opt">编</span><span class="shanchu2 file-list-opt">删</span></li>';
		}else{
			onefile += '<li hashid="'+value.HashId+'" filetype="d"><span class="file-list-type">D</span><span class="file-list-name" onclick=showChildList(this)>'+value.Name+'/</span><span class="xiazai2 file-list-opt">下</span><span class="bianji2 file-list-opt">编</span><span class="shanchu2 file-list-opt">删</span>';
			onefile += '<ul class="file-list-2" show="no">'
			onefile += showBigJsonLevel(json[value.Name].Files)
			onefile += '</ul></li>'
		}
		//$("#resource-one-full .resource-all-file .file-list").append(onefile);
	});
	return onefile;
};

var showChildList = function(self){
	theUl = $(self).parent().children(".file-list-2")
	if(theUl.attr("show") == "no"){
		$("#resource-one-full .resource-all-file .file-list-2").each(function(){
			$(this).hide().attr("show","no");
		});
		theUl.show();
		theUl.attr("show","yes")
	}else{
		theUl.hide();
		theUl.attr("show","no")
	}
}

var resourceLiulanClick = function(self){
	var allinfo = $('#resource-one-full .resource-all-info')
	var filelist = $('#resource-one-full .resource-all-file')
	if(allinfo.attr("showit") == "yes"){
		allinfo.hide().attr("showit","no");
		filelist.show().attr("showit","yes");
	}else{
		filelist.hide().attr("showit","no");
		allinfo.show().attr("showit","yes");
		return;
	};
	hashid = $('#resource-one-full').attr("hashid");
	$("#nowloadbox").fadeIn(200);
	$.get("webInterface?type=resource-file&hashid="+hashid , function(data){
		theBigJSON = $.parseJSON(data);
		thelist = showBigJsonLevel(theBigJSON);
		$("#resource-one-full .resource-all-file .file-list").html(thelist);
		$("#nowloadbox").fadeOut(200);
	});
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
	var server_word = "";
	if(search_word === ""){
		server_word = "webInterface?type=resource-list&from="+iResourceList_from+"&limit="+iResourceList_limit
		$("#resource-list .allListBookCountTishi").text("本社共有图书")
	}else{
		server_word = "webInterface?type=resource-list&key_word="+search_word+"&search_type="+search_type+"&from="+iResourceList_from+"&limit="+iResourceList_limit
		$("#resource-list .allListBookCountTishi").text("共找到图书")
	}
	$.get(server_word , function(data){
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
		if($("#resource-list .allListBookCount").text() == '0'){
			$("#next-and-prev .next").hide();
			$("#next-and-prev .prev").hide();
		}
		var i = 0;
		
		md_converter = new Markdown.Converter();
		
		$(json.List).each(function(){
			//var li = $.parseJSON(this.MetaData);
			var ptime = lastOtime(this.Table.Btime);
			
			var md_c = md_converter.makeHtml(this.Table.Info);
			
			var onebook = '<div class="one-resource-main" hashid="'+this.Table.HashId+'">\
			<div class="one-resource-total-info">\
				<div class="the-resource-name" onclick=resourceNameClick(this)>'+this.Table.Name+'</div>\
				<div class="the-little-info">类型：'+this.RSR.RtName+'&nbsp;&nbsp;&nbsp;&nbsp;作者：'+this.MD.Author+'<br>编辑：'+this.MD.Editor+'&nbsp;&nbsp;&nbsp;&nbsp;ISBN/ISSN：'+this.MD.ISBN+'</div>\
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

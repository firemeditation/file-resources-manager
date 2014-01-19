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
$(window).resize(function(){ changeFullResouceBoxHeight(); });


// 隐藏这本书的一切信息组
var irlHideAll = function(){
	$('#resource-one-full .resource-all-info').hide();
	$('#resource-one-full .resource-all-file').hide();
	$('#resource-one-full .resource-delete-all').hide();
}

// 资源图书列表下，点击书名，打开这本书的信息
var resourceNameClick = function(self){
	irlHideAll();
	var allinfo = $(self).parent().parent().children(".resource-all-info").html();
	var hashid = $(self).parent().parent().attr("hashid");
	var bookname = $(self).parent().parent().children(".one-resource-total-info").children(".the-resource-name").text();
	$('#resource-one-full').attr("hashid",hashid);
	$('#resource-one-full .uppon-info-show .resource-all-info').html(allinfo);
	$('#resource-one-full .the-resource-name').text(bookname);
	changeFullResouceBoxHeight();
	$("#allwhite").show();
	$('#resource-one-full').show();
	$('#resource-one-full .resource-all-info').show();	
};

// 关闭这本书的操作
var resourceCloseNow = function(self){
	$("#allwhite").hide();
	$('#resource-one-full').hide();
	var allinfo = $('#resource-one-full .resource-all-info')
	var filelist = $('#resource-one-full .resource-all-file')
	filelist.hide().attr("showit","no");
	allinfo.show().attr("showit","yes");
}

// 返回最后操作时间
var lastOtime = function(utime){
	var theTime = utime * 1000;
	var timedate = new Date(theTime);
	var theTime = timedate.formatDate("yyyy年MM月dd日 hh:mm:ss");
	return theTime;
};

// 从服务器上获取资源图书的列表
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
		if(json.err){alert(json.err); processServerError(json.err); return;}
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
			<div class="resource-all-info" showit="no"><p>类型：'+this.RSR.RtName+'&nbsp;&nbsp;最后操作人：'+this.RSR.UsersName+'&nbsp;&nbsp;创建时间：'+ptime+'</p>\
			<p>作者：'+this.MD.Author+'&nbsp;&nbsp;编辑：'+this.MD.Editor+'&nbsp;&nbsp;ISBN/ISSN：'+this.MD.ISBN+'&nbsp;&nbsp;</p>\
			<p>简介：</p>\
			<div class="markdown">'+md_c+'</div></div>\
		</div>';
			$("#resource-main-list").append(onebook);
			i++;
		});
		$("#nowloadbox").fadeOut(200);
	});
};

getResourceListFromServer();

//点击这本书的上传
$("#resource-one-full .one-resource-total-info .shangchuan").click(function(){
	if(login_user.UPower.resource.origin < 2){ return }
	var hashid = $('#resource-one-full').attr("hashid");
	var bookname = $('#resource-one-full .the-resource-name').text();
	$("#allwhite2").load("static/iAddResourceUpload.htm", function(){
		$("#allwhite2").attr("hashid", hashid);
		$("#allwhite2").attr("opentype", "irl");
		$("#allwhite2").attr("bookname", bookname);
		$("#allwhite2").attr("path", "");
		$.getScript("static/iAddResourceUpload.js").done(function(){
			$("#allwhite2").show();
		});
	});
});

// 将文件上传到指定路径下
var iRLUpFile = function(path){
	if(login_user.UPower.resource.origin < 2){ return }
	var hashid = $('#resource-one-full').attr("hashid");
	var bookname = $('#resource-one-full .the-resource-name').text();
	$("#allwhite2").load("static/iAddResourceUpload.htm", function(){
		$("#allwhite2").attr("hashid", hashid);
		$("#allwhite2").attr("opentype", "irl");
		$("#allwhite2").attr("bookname", bookname);
		$("#allwhite2").attr("path", path);
		$.getScript("static/iAddResourceUpload.js").done(function(){
			$("#allwhite2").show();
		});
	});
}

// 显示图书内所有文件的目录树
var showBigJsonLevel = function(json, path, jsoo){
	//$("#resource-one-full .resource-all-file .now-dir .true-now").text(path);
	//$("#resource-one-full .resource-all-file .file-list").html("");
	var onefile = ''
	$.each(json, function(name, value){
		if (value.IsDir == false){
			onefile += '<li hashid="'+value.HashId+'" filetype="f"><span>├</span><span class="file-list-type">F</span><span class="file-list-name">'+value.Name+'</span><span class="xiazai2 file-list-opt">下</span><span class="bianji2 file-list-opt">编</span><span class="shanchu2 file-list-opt" onclick=irlDeleteOneFile("'+value.HashId+'")>删</span></li>';
		}else{
			onefile += '<li filetype="d"><span>├</span><span class="file-list-type">D</span><span class="file-list-name" onclick=showChildList(this)>'+value.Name+'/</span><span class="xiazai2 file-list-opt">下</span><span class="shangchuan2 file-list-opt" onclick=iRLUpFile("'+path + value.Name + '/")>上</span><span class="bianji2 file-list-opt">编</span><span class="shanchu2 file-list-opt" onclick=irlDeletePartFile('+jsoo + '["' + value.Name + '"]' +'.Files)>删</span>';
			onefile += '<ul class="file-list-2" show="no">'
			onefile += showBigJsonLevel(json[value.Name].Files, path+value.Name+"/", jsoo + '["'+value.Name+'"].Files')
			onefile += '</ul></li>'
		}
		//$("#resource-one-full .resource-all-file .file-list").append(onefile);
	});
	return onefile;
};

// 展开目录树的下一级
var showChildList = function(self){
	theUl = $(self).parent().children(".file-list-2")
	if(theUl.attr("show") == "no"){
		//$("#resource-one-full .resource-all-file .file-list-2").each(function(){
		//	$(this).hide().attr("show","no");
		//});
		theUl.show();
		theUl.attr("show","yes")
	}else{
		theUl.hide();
		theUl.attr("show","no")
	}
}

// 显示这本书的详情页
var resourceXiangqingClick = function(){
	irlHideAll();
	$('#resource-one-full .resource-all-info').show();
}

// 显示这本书的文件树
var resourceLiulanClick = function(){
	irlHideAll();
	var filelist = $('#resource-one-full .resource-all-file');
	filelist.show();
	hashid = $('#resource-one-full').attr("hashid");
	$("#nowloadbox").fadeIn(200);
	$.get("webInterface?type=resource-file&hashid="+hashid , function(data){
		theBigJSON = $.parseJSON(data);
		if(theBigJSON.err){alert(theBigJSON.err); processServerError(theBigJSON.err); return;}
		thelist = showBigJsonLevel(theBigJSON,'',"theBigJSON");
		$("#resource-one-full .resource-all-file .file-list").html(thelist);
		$("#nowloadbox").fadeOut(200);
	});
};


// 获取此目录下所有文件的hashid
var irlGetAllHashid = function(json){
	var allhashid = '';
	$.each(json, function(name, value){
		if (value.IsDir == false){
			allhashid += "'"+value.HashId+"',";
		}else{
			allhashid += irlGetAllHashid(json[value.Name].Files);
		}
	});
	return allhashid;
}

/** begin 删除图书 **/

// 显示这本书的删除选项
var resourceDeleteAllClick = function(){
	if(login_user.UPower.resource.origin < 2){ return }
	irlHideAll();
	$('#resource-one-full .resource-delete-all').show();
}

//删除这本书的一切，包括里面的文件，以及这个条目本身
var irlDoDropAll = function(){
	hashid = $('#resource-one-full').attr("hashid");
	if(confirm("确定要删除这个条目的一切？一定要想清楚！")){
		$.get("webInterface?type=delete-resource-group&hashid="+hashid, function(data){
			theJSON = $.parseJSON(data);
			if(theJSON.err){alert(theJSON.err); processServerError(theJSON.err); return;}
			if(theJSON.ok){alert(theJSON.ok);}
			resourceCloseNow();
			getResourceListFromServer();
		});
	};
}

// 清空这本书的所有文件
var irlDoDelAllFile = function(){
	hashid = $('#resource-one-full').attr("hashid");
	if(confirm("确定要清空文件所有文件？一定要想清楚！")){
		$.get("webInterface?type=delete-resource-file&hashid="+hashid+"&dtype=all",function(data){
			theJSON = $.parseJSON(data);
			if(theJSON.err){alert(theJSON.err); processServerError(theJSON.err); return;}
			if(theJSON.ok){alert(theJSON.ok);}
			resourceXiangqingClick();
		});
	};
}

// 删除部分文件
var irlDeletePartFile = function(json){
	allfile = irlGetAllHashid(json);
	hashid = $('#resource-one-full').attr("hashid");
	if(confirm("确定要删除这些文件吗？这是不可逆操作，一定要想清楚！")){
		$.get("webInterface?type=delete-resource-file&hashid="+hashid+"&dtype=part&file="+allfile,function(data){
			theJSON = $.parseJSON(data);
			if(theJSON.err){alert(theJSON.err); processServerError(theJSON.err); return;}
			if(theJSON.ok){alert(theJSON.ok);}
			resourceLiulanClick();
		});
	}
};

// 删除一个文件
var irlDeleteOneFile = function(hashid){
	hashid = $('#resource-one-full').attr("hashid");
	if(confirm("确定要删除这个文件吗？这是不可逆操作，一定要想清楚！")){
		$.get("webInterface?type=delete-resource-file&hashid="+hashid+"&dtype=one&file="+hashid,function(data){
			theJSON = $.parseJSON(data);
			if(theJSON.err){alert(theJSON.err); processServerError(theJSON.err); return;}
			if(theJSON.ok){alert(theJSON.ok);}
			resourceLiulanClick();
		});
	}
};

/** end 删除图书 **/


//点击下一页
$("#next-and-prev .next").click(function(){
	iResourceList_from = iResourceList_from + iResourceList_limit;
	getResourceListFromServer();
});

//点击上一页
$("#next-and-prev .prev").click(function(){
	iResourceList_from = iResourceList_from - iResourceList_limit;
	getResourceListFromServer();
});

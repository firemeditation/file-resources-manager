var iaruPath = function(){
	path = $("#allwhite2").attr("path")
	$("#iAddEditResourceUpload .relativepath input").val(path);
}

iaruPath();


$("#iAddEditResourceUploadForm .submit input[name='close']").click(function(){
	$("#allwhite2").hide().html("");
	if( $("#allwhite2").attr("opentype") == "aer" ){
		hideAll();
		searchClean();
		showTheBasicResourceList();
	}else if($("#allwhite2").attr("opentype") == "irl"){
		$("#iAddEditResourceUploadForm").hide();
	}
});

$("#iAddEditResourceUploadForm .submit input[name='submit']").click(function(){
	var ckArray = [0,1];
	ckArray[0] = $.RequestProcess.Text('#iAddEditResourceUploadForm .localpath',0,1,1000);
    ckArray[1] = $.RequestProcess.Text('#iAddEditResourceUploadForm .relativepath',1,1,1000);
    if($.RequestProcess.ckAllOne(ckArray)==0){ return }
    var $localpath = inputSafe.CleanAll($("#iAddEditResourceUploadForm .localpath input").val());
    var $relativepath = inputSafe.CleanAll($("#iAddEditResourceUploadForm .relativepath input").val());
    var $hashid = inputSafe.CleanAll($("#allwhite2").attr("hashid"));
    var $user = login_user.HashId;
    var $bookname = inputSafe.CleanAll($("#allwhite2").attr("bookname"));
    $.getJSON("http://127.0.0.1:"+local_client_port+"/uploadFile?user="+$user+"&bookname="+$bookname+"&local="+$localpath+"&relative="+$relativepath+"&hashid="+$hashid+"&callback=?", function(data){
		if (data.err) {
			alert(data.err);
			return
		}else{
			alert("已经转向后台上传，具体请查看后台状态。")
		}
		$("#allwhite2").hide().html("");
		if( $("#allwhite2").attr("opentype") == "aer" ){
			hideAll();
			searchClean();
			showTheBasicResourceList();
		}else if($("#allwhite2").attr("opentype") == "irl"){
			$("#iAddEditResourceUploadForm").hide();
		}
	})
});

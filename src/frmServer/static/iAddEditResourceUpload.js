$("#iAddEditResourceUploadForm .submit input[name='close']").click(function(){
	$("#allwhite").hide().html("");
	if( $("#allwhite").attr("opentype") == "aer" ){
		hideAll();
		$("#main-box #resource-list").fadeIn();
	}
});

$("#iAddEditResourceUploadForm .submit input[name='submit']").click(function(){
	ckArray[0] = $.RequestProcess.Text('#iAddEditResourceUploadForm .localpath',0,1,1000);
    ckArray[1] = $.RequestProcess.Text('#iAddEditResourceUploadForm .relativepath',1,1,1000);
    if($.RequestProcess.ckAllOne(ckArray)==0){ return }
    var $localpath = inputSafe.CleanAll($("#iAddEditResourceUploadForm .localpath input").val());
    var $relativepath = inputSafe.CleanAll($("#iAddEditResourceUploadForm .relativepath input").val());
    var $hashid = inputSafe.CleanAll($("#allwhite").attr("hashid"));
    var $user = login_user.HashId;
    var $bookname = inputSafe.CleanAll($("#allwhite").attr("bookname"));
    $.getJSON("http://127.0.0.1:"+local_client_port+"/uploadFile?user="+$user+"&bookname="+$bookname+"&local="+$localpath+"&relative="+$relativepath+"&hashid="+$hashid+"&callback=?")
	$("#allwhite").hide().html("");
	if( $("#allwhite").attr("opentype") == "aer" ){
		hideAll();
		$("#main-box #resource-list").fadeIn();
	}
});

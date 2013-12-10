$.get("webInterface?type=get-resource-type",function(data){
	//alert(data);
	var json = $.parseJSON(data);
	$(json).each(function(){
		$('#iAddEditResource .resoucetype select').append("<option value='" + this.Id + "'>" + this.Name + "</option>")
	});
});

	$('#iAddEditResource .bookname input').keyup(function(){ $.RequestProcess.Text('#iAddEditResource .bookname',0,1,1000); });
    $('#iAddEditResource .author input').keyup(function(){ $.RequestProcess.Text('#iAddEditResource .author',1,1,1000); });
    $('#iAddEditResource .editor input').keyup(function(){ $.RequestProcess.Text('#iAddEditResource .editor',1,1,1000); });
    $('#iAddEditResource .isbn input').keyup(function(){ $.RequestProcess.Text('#iAddEditResource .isbn',1,1,1000); });
    $('#iAddEditResource .info input').keyup(function(){ $.RequestProcess.Textarea('#iAddEditResource .info',1,1,99999); });
    var ckArray = [0,1,1,1,1];
    iAddEditResource_add = function(){
        ckArray[0] = $.RequestProcess.Text('#iAddEditResource .bookname',0,1,1000);
        ckArray[1] = $.RequestProcess.Text('#iAddEditResource .author',1,1,1000);
        ckArray[2] = $.RequestProcess.Text('#iAddEditResource .editor',1,1,1000);
        ckArray[3] = $.RequestProcess.Text('#iAddEditResource .isbn',1,1,1000);
        ckArray[4] = $.RequestProcess.Textarea('#iAddEditResource .info',1,1,99999);
        if($.RequestProcess.ckAllOne(ckArray)==0){ alert("输入错误") ;}
    };

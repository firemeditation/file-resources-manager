/* cc_RequestProcess
 * CODEC的输入处理函数
 * 作为jQuery Plugin，需要配合jQuery使用
 * 为File Resource Manager做了部分修改
 */

jQuery.cc_RequestProcess = {
	Text: function($theID,$cannull,$min,$max){
		var $SID = $theID + ' .input_wrong_suggest';
		var $IID = $theID + ' input';
		var $ck = 0;
		var $theVal = cTrim($($IID).val(),0);
		var $theLength = mbStringLength($theVal);
		if($theLength == 0 && $cannull == 0){
			$($SID).text('不能为空').addClass('input_wrong_suggest_wrong');
			$ck = 0;
		}else if($theLength == 0 && $cannull == 1){
			$($SID).empty().removeClass('input_wrong_suggest_wrong');
			$ck = 1;
		}else{
			if($theLength > $max){
				$($SID).text('字数太多').addClass('input_wrong_suggest_wrong');
				$ck = 0;
			}else if($theLength < $min){
				$($SID).text('字数太少').addClass('input_wrong_suggest_wrong');
				$ck = 0;
			}else{
				$($SID).empty().removeClass('input_wrong_suggest_wrong');
				$ck = 1;
			}
		}
		return $ck;
	},
	Textarea: function($theID,$cannull,$min,$max){
		var $SID = $theID + ' .input_wrong_suggest';
		var $IID = $theID + ' textarea';
		var $ck = 0;
		var $theVal = cTrim($($IID).val(),0);
		var $theLength = mbStringLength($theVal);
		if($theLength == 0 && $cannull == 0){
			$($SID).text('不能为空').addClass('input_wrong_suggest_wrong');
			$ck = 0;
		}else if($theLength == 0 && $cannull == 1){
			$($SID).empty().removeClass('input_wrong_suggest_wrong');
			$ck = 1;
		}else{
			if($theLength > $max){
				$($SID).text('字数太多').addClass('input_wrong_suggest_wrong');
				$ck = 0;
			}else if($theLength < $min){
				$($SID).text('字数太少').addClass('input_wrong_suggest_wrong');
				$ck = 0;
			}else{
				$($SID).empty().removeClass('input_wrong_suggest_wrong');
				$ck = 1;
			}
		}
		return $ck;
	},
	Editor: function($theID,$content,$cannull,$min,$max){
		var $SID = $theID + ' .input_wrong_suggest';
		var $IID = $theID + ' textarea';
		var $ck = 0;
		var $theLength = mbStringLength($content);
		if($theLength == 0 && $cannull == 0){
			$($SID).text('不能为空').addClass('input_wrong_suggest_wrong');
			$ck = 0;
		}else if($theLength == 0 && $cannull == 1){
			$($SID).empty().removeClass('input_wrong_suggest_wrong');
			$ck = 1;
		}else{
			if($theLength > $max){
				$($SID).text('字数太多').addClass('input_wrong_suggest_wrong');
				$ck = 0;
			}else if($theLength < $min){
				$($SID).text('字数太少').addClass('input_wrong_suggest_wrong');
				$ck = 0;
			}else{
				$($SID).empty().removeClass('input_wrong_suggest_wrong');
				$ck = 1;
			}
		}
		return $ck;
	},
	Int: function($theID,$cannull,$min,$max){
		var $SID = $theID + ' .input_wrong_suggest';
		var $IID = $theID + ' input';
		var $ck = 0;
		var $theVal = cTrim($($IID).val(),0);
		var $theLength = mbStringLength($theVal);
		if($theLength == 0 && $cannull == 0){
			$($SID).text('不能为空').addClass('input_wrong_suggest_wrong');
			$ck = 0;
		}else if($theLength == 0 && $cannull == 1){
			$($SID).empty().removeClass('input_wrong_suggest_wrong');
			$ck = 1;
		}else{
			if(!/^([0-9-]+)$/.test($theVal)){
				$($SID).text('整数格式错误').addClass('input_wrong_suggest_wrong');
				$ck = 0;
			}else{
				if($theVal > $max){
					$($SID).text('数字过大').addClass('input_wrong_suggest_wrong');
					$ck = 0;
				}else if($theVal < $min){
					$($SID).text('数字过小').addClass('input_wrong_suggest_wrong');
					$ck = 0;
				}else{
					$($SID).empty().removeClass('input_wrong_suggest_wrong');
					$ck = 1;
				}
			}
		}
		return $ck;
	},
	Number: function($theID,$cannull,$min,$max){
		var $SID = $theID + ' .input_wrong_suggest';
		var $IID = $theID + ' input';
		var $ck = 0;
		var $theVal = cTrim($($IID).val(),0);
		var $theLength = mbStringLength($theVal);
		if($theLength == 0 && $cannull == 0){
			$($SID).text('不能为空').addClass('input_wrong_suggest_wrong');
			$ck = 0;
		}else if($theLength == 0 && $cannull == 1){
			$($SID).empty().removeClass('input_wrong_suggest_wrong');
			$ck = 1;
		}else{
			if(!/^-?([1-9]\d*\.\d*|0\.\d*[1-9]\d*|0?\.0+|[0-9]+)$/.test($theVal)){
				$($SID).text('数字格式错误').addClass('input_wrong_suggest_wrong');
				$ck = 0;
			}else{
				if($theVal > $max){
					$($SID).text('数字过大').addClass('input_wrong_suggest_wrong');
					$ck = 0;
				}else if($theVal < $min){
					$($SID).text('数字过小').addClass('input_wrong_suggest_wrong');
					$ck = 0;
				}else{
					$($SID).empty().removeClass('input_wrong_suggest_wrong');
					$ck = 1;
				}
			}
		}
		return $ck;
	},
	Money: function($theID,$cannull,$min,$max){
		var $SID = $theID + ' .input_wrong_suggest';
		var $IID = $theID + ' input';
		var $ck = 0;
		var $theVal = cTrim($($IID).val(),0);
		var $theLength = mbStringLength($theVal);
		if($theLength == 0 && $cannull == 0){
			$($SID).text('不能为空').addClass('input_wrong_suggest_wrong');
			$ck = 0;
		}else if($theLength == 0 && $cannull == 1){
			$($SID).empty().removeClass('input_wrong_suggest_wrong');
			$ck = 1;
		}else{
			if(!/^-?([1-9]\d*\.\d{1,2}|0\.\d*[1-9]\d{1,2}|0?\.0+|[0-9]+)$/.test($theVal)){
				$($SID).text('货币格式错误').addClass('input_wrong_suggest_wrong');
				$ck = 0;
			}else{
				if($theVal > $max){
					$($SID).text('数字过大').addClass('input_wrong_suggest_wrong');
					$ck = 0;
				}else if($theVal < $min){
					$($SID).text('数字过小').addClass('input_wrong_suggest_wrong');
					$ck = 0;
				}else{
					$($SID).empty().removeClass('input_wrong_suggest_wrong');
					$ck = 1;
				}
			}
		}
		return $ck;
	},
	Mark: function($theID,$cannull,$min,$max){
		var $SID = $theID + ' .input_wrong_suggest';
		var $IID = $theID + ' input';
		var $ck = 0;
		var $theVal = cTrim($($IID).val(),0);
		var $theLength = mbStringLength($theVal);
		if($theLength == 0 && $cannull == 0){
			$($SID).text('不能为空').addClass('input_wrong_suggest_wrong');
			$ck = 0;
		}else if($theLength == 0 && $cannull == 1){
			$($SID).empty().removeClass('input_wrong_suggest_wrong');
			$ck = 1;
		}else{
			if(!/^[A-Za-z0-9_-]+$/.test($theVal)){
				$($SID).text('标识格式错误').addClass('input_wrong_suggest_wrong');
				$ck = 0;
			}else if($theLength > $max){
				$($SID).text('字数太多').addClass('input_wrong_suggest_wrong');
				$ck = 0;
			}else if($theLength < $min){
				$($SID).text('字数太少').addClass('input_wrong_suggest_wrong');
				$ck = 0;
			}else{
				$($SID).empty().removeClass('input_wrong_suggest_wrong');
				$ck = 1;
			}
		}
		return $ck;
	},
	PathName: function($theID,$cannull,$min,$max){
		var $SID = $theID + ' .input_wrong_suggest';
		var $IID = $theID + ' input';
		var $ck = 0;
		var $theVal = cTrim($($IID).val(),0);
		var $theLength = mbStringLength($theVal);
		if($theLength == 0 && $cannull == 0){
			$($SID).text('不能为空').addClass('input_wrong_suggest_wrong');
			$ck = 0;
		}else if($theLength == 0 && $cannull == 1){
			$($SID).empty().removeClass('input_wrong_suggest_wrong');
			$ck = 1;
		}else{
			if(!/^[A-Za-z0-9]+[A-Za-z0-9_\/-]+[\/]$/.test($theVal)){
				$($SID).text('路径名格式错误').addClass('input_wrong_suggest_wrong');
				$ck = 0;
			}else if($theLength > $max){
				$($SID).text('字数太多').addClass('input_wrong_suggest_wrong');
				$ck = 0;
			}else if($theLength < $min){
				$($SID).text('字数太少').addClass('input_wrong_suggest_wrong');
				$ck = 0;
			}else{
				$($SID).empty().removeClass('input_wrong_suggest_wrong');
				$ck = 1;
			}
		}
		return $ck;
	},
	FileName: function($theID,$cannull,$min,$max){
		var $SID = $theID + ' .input_wrong_suggest';
		var $IID = $theID + ' input';
		var $ck = 0;
		var $theVal = cTrim($($IID).val(),0);
		var $theLength = mbStringLength($theVal);
		if($theLength == 0 && $cannull == 0){
			$($SID).text('不能为空').addClass('input_wrong_suggest_wrong');
			$ck = 0;
		}else if($theLength == 0 && $cannull == 1){
			$($SID).empty().removeClass('input_wrong_suggest_wrong');
			$ck = 1;
		}else{
			if(!/^[A-Za-z0-9#_]{1}[A-Za-z0-9_\.#-]*[A-Za-z0-9#]{1}$/.test($theVal)){
				$($SID).text('文件名格式错误').addClass('input_wrong_suggest_wrong');
				$ck = 0;
			}else if($theLength > $max){
				$($SID).text('字数太多').addClass('input_wrong_suggest_wrong');
				$ck = 0;
			}else if($theLength < $min){
				$($SID).text('字数太少').addClass('input_wrong_suggest_wrong');
				$ck = 0;
			}else{
				$($SID).empty().removeClass('input_wrong_suggest_wrong');
				$ck = 1;
			}
		}
		return $ck;
	},
	Password: function($theID,$cannull,$min){
		var $SID = $theID + ' .input_wrong_suggest';
		var $IID = $theID + ' input';
		var $ck = 0;
		var $theVal = cTrim($($IID).val(),0);
		var $theLength = mbStringLength($theVal);
		if($theLength == 0 && $cannull == 0){
			$($SID).text('不能为空').addClass('input_wrong_suggest_wrong');
			$ck = 0;
		}else if($theLength == 0 && $cannull == 1){
			$($SID).empty().removeClass('input_wrong_suggest_wrong');
			$ck = 1;
		}else{
			if($theLength < $min){
				$($SID).text('密码太短').addClass('input_wrong_suggest_wrong');
				$ck = 0;
			}else {
				$($SID).empty().removeClass('input_wrong_suggest_wrong');
				$ck = 1;
			}
		}
		return $ck;
	},
	PasswordOne: function($theID,$cannull,$min){
		var $SID = $theID + ' .input_wrong_suggest';
		var $IID = $theID + ' input';
		var $UID = $theID + ' .input_unique_suggest';
		var $ck = 0;
		var $theVal = cTrim($($IID).val(),0);
		var $theLength = mbStringLength($theVal);
		if($theLength == 0 && $cannull == 0){
			$($SID).text('不能为空').addClass('input_wrong_suggest_wrong');
			$ck = 0;
		}else if($theLength == 0 && $cannull == 1){
			$($SID).empty().removeClass('input_wrong_suggest_wrong');
			$ck = 1;
		}else{
			if($theLength < $min){
				$($UID).empty().removeClass('input_wrong_suggest_weak');
				$($SID).text('密码太短').addClass('input_wrong_suggest_wrong');
				$ck = 0;
			}else if(/^[0-9]+$/.test($theVal) || /^[A-Za-z-]+$/.test($theVal)){
				$($SID).empty().removeClass('input_wrong_suggest_wrong');
				$($UID).text('密码较弱').addClass('input_wrong_suggest_weak');
				$ck = 1;
			}else {
				$($SID).empty().removeClass('input_wrong_suggest_wrong');
				$($UID).empty().removeClass('input_wrong_suggest_weak');
				$ck = 1;
			}
		}
		return $ck;
	},
	PasswordTwo: function($theID1,$theID2,$cannull){
		var $IID1 = $theID1 + ' input';
		var $SID2 = $theID2 + ' .input_wrong_suggest';
		var $IID2 = $theID2 + ' input';
		var $ck = 0;
		var $theVal2 = cTrim($($IID2).val(),0);
		var $theVal1 = cTrim($($IID1).val(),0);
		var $theLength1 = mbStringLength($theVal1);
		var $theLength2 = mbStringLength($theVal2);
		if($theLength2 == 0 && $theLength1 != 0){
			$($SID2).text('不能为空').addClass('input_wrong_suggest_wrong');
			$ck = 0;
		}else if($theLength2 == 0 && $theLength1 == 0){
			$($SID2).empty().removeClass('input_wrong_suggest_wrong');
			$ck = 1;
		}else{
			if($theVal2 != $theVal1){
				$($SID2).text('两遍密码不一致').addClass('input_wrong_suggest_wrong');
				$ck = 0;
			}else{
				$($SID2).empty().removeClass('input_wrong_suggest_wrong');
				$ck = 1;
			}
		}
		return $ck;
	},
	Regular:function($theID,$cannull,$regular){
		var $SID = $theID + ' .input_wrong_suggest';
		var $IID = $theID + ' input';
		var $ck = 0;
		var $theVal = cTrim($($IID).val(),0);
		var $theLength = mbStringLength($theVal);
		if($theLength == 0 && $cannull == 0){
			$($SID).text('不能为空').addClass('input_wrong_suggest_wrong');
			$ck = 0;
		}else if($theLength == 0 && $cannull == 1){
			$($SID).empty().removeClass('input_wrong_suggest_wrong');
			$ck = 1;
		}else{
			var pattern = new RegExp($regular);
			if(pattern.test($theVal)){
				$($SID).empty().removeClass('input_wrong_suggest_wrong');
				$ck = 1;
			}else{
				$($SID).text('输入格式错误').addClass('input_wrong_suggest_wrong');
				$ck = 0;
			}
		}
		return $ck;
	},
	Date: function($theID,$cannull){
		var $SID = $theID + ' .input_wrong_suggest';
		var $IID = $theID + ' input';
		var $ck = 0;
		var $theVal = cTrim($($IID).val(),0);
		var $theLength = mbStringLength($theVal);
		if($theLength == 0 && $cannull == 0){
			$($SID).text('不能为空').addClass('input_wrong_suggest_wrong');
			$ck = 0;
		}else if($theLength == 0 && $cannull == 1){
			$($SID).empty().removeClass('input_wrong_suggest_wrong');
			$ck = 1;
		}else{
			if(!/((^((1[8-9]\d{2})|([2-9]\d{3}))-(10|12|0?[13578])-(3[01]|[12][0-9]|0?[1-9])$)|(^((1[8-9]\d{2})|([2-9]\d{3}))-(11|0?[469])-(30|[12][0-9]|0?[1-9])$)|(^((1[8-9]\d{2})|([2-9]\d{3}))-(0?2)-(2[0-8]|1[0-9]|0?[1-9])$)|(^([2468][048]00)-(0?2)-(29)$)|(^([3579][26]00)-(0?2)-(29)$)|(^([1][89][0][48])-(0?2)-(29)$)|(^([2-9][0-9][0][48])-(0?2)-(29)$)|(^([1][89][2468][048])-(0?2)-(29)$)|(^([2-9][0-9][2468][048])-(0?2)-(29)$)|(^([1][89][13579][26])-(0?2)-(29)$)|(^([2-9][0-9][13579][26])-(0?2)-(29)$))/.test($theVal)){
				$($SID).text('日期格式错误').addClass('input_wrong_suggest_wrong');
				$ck = 0;
			}else{
				$($SID).empty().removeClass('input_wrong_suggest_wrong');
				$ck = 1;
			}
		}
		return $ck;
	},
	Mail: function($theID,$cannull){
		var $SID = $theID + ' .input_wrong_suggest';
		var $IID = $theID + ' input';
		var $ck = 0;
		var $theVal = cTrim($($IID).val(),0);
		var $theLength = mbStringLength($theVal);
		if($theLength == 0 && $cannull == 0){
			$($SID).text('不能为空').addClass('input_wrong_suggest_wrong');
			$ck = 0;
		}else if($theLength == 0 && $cannull == 1){
			$($SID).empty().removeClass('input_wrong_suggest_wrong');
			$ck = 1;
		}else{
			if(!/^[A-Za-z0-9]+([_\.-][A-Za-z0-9]+)*@[A-Za-z0-9]+([_\.-][A-Za-z0-9]+)*\.([A-Za-z]){2,4}$/.test($theVal)){
				$($SID).text('邮箱格式错误').addClass('input_wrong_suggest_wrong');
				$ck = 0;
			}else{
				$($SID).empty().removeClass('input_wrong_suggest_wrong');
				$ck = 1;
			}
		}
		return $ck;
	},
	Url: function($theID,$cannull){
		var $SID = $theID + ' .input_wrong_suggest';
		var $IID = $theID + ' input';
		var $ck = 0;
		var $theVal = cTrim($($IID).val(),0);
		var $theLength = mbStringLength($theVal);
		if($theLength == 0 && $cannull == 0){
			$($SID).text('不能为空').addClass('input_wrong_suggest_wrong');
			$ck = 0;
		}else if($theLength == 0 && $cannull == 1){
			$($SID).empty().removeClass('input_wrong_suggest_wrong');
			$ck = 1;
		}else{
			if(!/^[a-zA-z]+:\/\/[^\s]*$/.test($theVal)){
				$($SID).text('网址格式错误').addClass('input_wrong_suggest_wrong');
				$ck = 0;
			}else{
				$($SID).empty().removeClass('input_wrong_suggest_wrong');
				$ck = 1;
			}
		}
		return $ck;
	},
	Unique: function($theUrl,$theField,$theType,$theID){
		var $UID = $theField + ' .input_unique_suggest';
		var $SID = $theField + ' .input_wrong_suggest';
		var $IID = $theField + ' input';
		var $content = $($IID).val();
		var timeStamp = new Date().getTime();
		var $ck = 0;
		$.ajax({
			type: "POST",
			url: $theUrl,
			cache: "false",
			dataType: 'html',
			data: {'field':$theField,'timeStamp':timeStamp,'type':$theType,'id':$theID,'content':$content},
			global: "false",
			success: function(html){
				//alert(html);
				if(html == 'OK'){
					$($UID).empty().removeClass('input_unique_suggest_wrong');
				}else{
					$($SID).empty().removeClass('input_wrong_suggest_wrong');
					$($UID).text('已经被使用，请更换').addClass('input_unique_suggest_wrong');
					//$($theSubmit).attr({"disabled": "disabled", "class": "input_submit_disabled"});
				};
			}
		});
	},
	ckAllOne: function($ckArray){
		var $theReturn = 1;
		for (var i in $ckArray){
			if($ckArray[i] == 0){
				$theReturn = 0;
			}
		}
		return $theReturn;
	},
	ckUnique: function($theID){
		$theID = $theID + ' .input_unique_suggest';
		$theTest = $($theID).text();
		$ck = 0;
		if(mbStringLength($theTest) == 0){
			$ck = 1;
		}else{
			$ck = 0;
		}
		return $ck;
	}
};

var inputSafe = {
	Clean : function($input){
		var $trans = {"'" : "&#039;" , "`" : "&acute;", '"' : '&quot;', "<script(.*)>(.*)<\/script>" : "&lt;script$1&gt;$2&lt;script&gt;"}
		for ($one in $trans){
			$re = new RegExp($one, "gi");
			$input = $input.replace($re,$trans[$one]);
		}
		return $input
	},
	CleanAll : function($input){
		var $trans = {"'" : "&#039;" , "`" : "&acute;", '"' : '&quot;', "<" : "&lt;", ">" : '&gt;'}
		for ($one in $trans){
			$re = new RegExp($one, "gi");
			$input = $input.replace($re,$trans[$one]);
		}
		return $input
	},
	CleanBack : function($input){
		var $trans = { "&#039;" : "'" , "&acute;" : "`" , '&quot;' : '"', "&lt;" : "<" , '&gt;' : ">"}
		for ($one in $trans){
			$re = new RegExp($one, "gi");
			$input = $input.replace($re,$trans[$one]);
		}
		return $input
	}
}

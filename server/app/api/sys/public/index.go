package public

import (
	"net/http"
)

const INDEX = `
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>edu/admin欢迎您</title>
<style>
body{
  margin:0; 
  padding:0; 
  overflow-y:hidden
}
</style>
<script src="http://libs.baidu.com/jquery/1.9.0/jquery.js"></script>
<script type="text/javascript"> 
window.onerror=function(){return true;} 
$(function(){ 
  headerH = 0;  
  var h=$(window).height();
  $("#iframe").height((h-headerH)+"px"); 
});
</script>
</head>
<body>
<iframe id="iframe" frameborder="0" src="http://doc.zhangwj.com/edu/admin-site/" style="width:100%;"></iframe>
</body>
</html>
`

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(INDEX))
}

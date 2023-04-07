/* 使用需要放在 jq后面 */

//公用Ajax Post 方法
function pubAjaxPOST(urlstr,senddatas,Func){
	var returnData;
	var returnInfo;
	urlstr = document.location.protocol+ "//" + window.location.host + urlstr;
	console.log(urlstr);
	$.ajax({
            url: urlstr,
            type:'post',
            dataType:'json',
            contentType:"application/x-www-form-urlencoded",
            async:true,//异步请求
            cache:false,
            data:JSON.stringify(senddatas),
            success:function(rdata) {
            	if(Func){
                    Func(rdata);
                }
                else{
                    console.log(rdata);
                }
            },
            //执行失败或错误的回调函数
            error:function(xhr) {
            	alert("后台请求出错！");
              	console.log(xhr);
              	location.reload();
            }
          });
	//console.log(returnData,returnInfo);
	//return returnData,returnInfo;
}

//公用Ajax Get 方法
function pubAjaxGET(urlstr,Func){
    urlstr = document.location.protocol+ "//" + window.location.host + urlstr;
    console.log(urlstr);
    $.ajax({
            url: urlstr,
            type:'get',
            dataType:'json',
            contentType:"application/x-www-form-urlencoded",
            async:true,//异步请求
            cache:false,
            success:function(rdata) {
                if(Func){
                    Func(rdata);
                }
                else{
                    console.log(rdata);
                }
            },
            //执行失败或错误的回调函数
            error:function(xhr) {
                alert("后台请求出错！");
                console.log(xhr);
                location.reload();
            }
          });
}


function Islogin(b){
  console.log("Islogin -- ");
  console.log(b);
  if(!b){
    top.location.reload();
  }
}
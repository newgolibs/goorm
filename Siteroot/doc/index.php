<?php

?>
<!DOCTYPE html>
<html>
<head>
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
    <meta name="viewport" content="width=device-width,initial-scale=1">
    <meta charset="UTF-8">
    <meta http-equiv="Expires" content="-1">
    <meta http-equiv="Pragma" content="no-cache">
    <meta http-equiv="Cache-control" content="no-cache">
    <meta http-equiv="Cache" content="no-cache">
    <link rel="stylesheet" href="https://static.xialintai.com/docsify/vue.css">
    <link rel="icon" type="image/png" href="favicon.png" />
    <title></title>
</head>
<body>
<div id="main">加载中</div>
<script>function adminlinkcheck(checkobj){var checked=checkobj.checked;Cookies.set('adminlinkchecked',checked);document.querySelectorAll('span.adminlink').forEach(function(item){item.style.display=(checked==true?'':'none')})}document.onkeypress=function(e){var ev=e||window.event;var obj=ev.relatedTarget||ev.srcElement||ev.target||ev.currentTarget;if(ev.keyCode==32){var c=document.querySelector('#showmanager');c.checked=!c.checked;adminlinkcheck(c);ev.preventDefault()}};window.$docsify={repo:'docsifyjs/docsify',name:'',themeColor:'#3F51B5',maxLevel:4,subMaxLevel:2,coverpage:true,auto2top:true,loadSidebar:true,mergeNavbar:true,loadNavbar:true,autoHeader:true,executeScript:true,externalLinkTarget:'_self',el:'#main',plantuml:{skin:'classic',renderSvgAsObject:true},}</script>

<script type="application/javascript">function x(){var el_object=event.currentTarget;var link=event.currentTarget.getAttribute('link');var edit_doc=document.querySelector("#edit_doc");if(edit_doc.style.display=='none'||edit_doc.src!=link){document.querySelectorAll('.ifmbtn').forEach(function(it){it.style.color="black"});el_object.style.color="red";if(edit_doc.src!=link){edit_doc.src=link}edit_doc.style.display='block'}else{el_object.style.color="black";edit_doc.style.display='none'}}</script>
</body>
</html>

<script type="application/javascript">var iframe={onclick:function($btn_object,$iframe_id){var $iframe_object=document.querySelector($iframe_id);var $url=$iframe_object.parentNode.querySelector('a').getAttribute('href');if($iframe_object.src==$url&&$iframe_object.status=='open'){$iframe_object.style.display='none';$btn_object.style.color='black';$iframe_object.status='close'}else{if($iframe_object.src!=$url)$iframe_object.src=$url;$iframe_object.style.display='block';$btn_object.style.color='red';$iframe_object.status='open'}},onmouseover:function($iframe_id){var $iframe_object=document.querySelector($iframe_id);var $url=$iframe_object.parentNode.querySelector('a').getAttribute('href');if($iframe_object.src==$url&&$iframe_object.status=='open'){window.t=window.setInterval(function(){$iframe_object.height=parseInt($iframe_object.height)+20;console.log(['i.height',$iframe_object.height])},100)}},onmouseover_less:function($iframe_id){var $iframe_object=document.querySelector($iframe_id);var $url=$iframe_object.parentNode.querySelector('a').getAttribute('href');if($iframe_object.src==$url&&$iframe_object.status=='open'){window.t=window.setInterval(function(){$iframe_object.height=parseInt($iframe_object.height)-20;console.log(['i.height',$iframe_object.height])},100)}},};</script>

<script type="application/javascript">
    //接收跨域iframe消息,用来调整本页面内嵌的iframe高度自适应
    window.addEventListener('message', function(e){
        document.querySelectorAll("iframe").forEach(function(item){
            if(item.src==e.data.url && e.data.height)
            {
                item.parentNode.style.height=e.data.height+10;
                item.height=e.data.height+10;
                console.log(["doc文档-最外层调整高度",e.data.message,e.data.height,e.data]);
            }
        });
    }, false);
</script>
<script src="https://static.xialintai.com/docsify/vue.min.js"></script>
<script src="https://static.xialintai.com/sso/OssVue/doc.js"></script>
<script src="https://static.xialintai.com/docsify/js.cookie.min.js"></script>
<script src="https://static.xialintai.com/docsify/docsify.min.js"></script>
<script src="https://static.xialintai.com/docsify/prism-bash.js"></script>
<script src="https://static.xialintai.com/docsify/prism-php.js"></script>
<script src="https://static.xialintai.com/docsify/prism-json.js"></script>
<script src="https://static.xialintai.com/docsify/zoom-image.js"></script>
<script src="https://static.xialintai.com/docsify/docsify-plantuml.min.js"></script>
<script src="https://static.xialintai.com/docsify/docsify-copy-code.min.js"></script>

<div id="main_iframe">
    <input id="showmanager" v-if="Cookies.get('username')=='夏琳泰'" v-bind:checked="Cookies.get('adminlinkchecked')=='true'" type="checkbox"  onchange="adminlinkcheck(this); ">
    <adminlink  name="编辑本页内容" link="/?c=Project_doc/Project_doc&prepage=30&webPageOrderBy=orderby&pageID=1&Project_doc_group=47&project_doc_navidoperation=%3D&status=%E9%80%9A%E8%BF%87&hidden_navigation=2" :show="Cookies.get('adminlinkchecked')"></adminlink>
    <adminlink  name="管理本大文档" link="/?c=Project_doc_group/Project_doc_group&project_nameoperation=%3D&prepage=30&id=47&project_name=goorm&hidden_navigation=2&status%5B%5D=%E9%80%9A%E8%BF%87" :show="Cookies.get('adminlinkchecked')"></adminlink>
    <div><iframe id="edit_doc" style="display: none;border: 2px solid darkred;min-height: 500px" src='' frameborder='0' scrolling='yes' width='100%'></iframe></div>
</div>


function openNewIframe(title, name, url) {
    var iframe = '<iframe frameborder="0" src="' + url + '" style="width:100%;height:100%;"></iframe>';
    var div = '<div id="' + name + '">' + iframe + '</div>';
}

//  模拟表单提交文件
function mockForm(action, file) {
    var form = $("<form></form>");
    form.attr("action", action);
    form.attr("method", "post");
    form.attr("enctype", "multipart/form-data");
    var input = $("<input type='file' name='file'>");
    form.append(input);
    input.files = [file];
    form.appendTo("body");
    form.css("display", "none");
    form.submit()
}
// 注：
// 侧边栏li的自定义属性lay-id需与url统一，方便后台操作


var $;
var layer;
var element;

//JavaScript代码区域
layui.use(['jquery', 'element', 'layer'], function() {
    $ = jQuery = layui.$;
    element = layui.element;
    layer = layui.layer;
    $("#tree li").on("click", openIframe)
});

function openIframe(e) {
    var text = $(this).text();
    var url = $(this).attr("lay-id");
    element.tabDelete('tab', url)
        //  iframe高度自适应问题待续...
    element.tabAdd('tab', {
        title: text,
        content: "<div><iframe frameborder='0' src='/admin/" + url + "' style='width:100%;min-height:800px;' id='iframe'></iframe></div>",
        id: url
    })
    element.tabChange('tab', url)
}
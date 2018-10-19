var layer;
layui.use(["layer"], function() {
    var layer = layui.layer;
    initData()
})

function initData() {
    layer.load(1, { shade: [0.2, '#000000'] });
    var id = getQeuryParam("id");
    asyncInvoke("/products.api", "GetItem", { id: id }, function(d) {
        if (d.code) {
            alert(d.data);
        } else {
            $("#article-content").html(d.data.htmlContent);
        }
        layer.closeAll('loading');
    })
}
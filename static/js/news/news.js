var pageIndex = parseInt(getQeuryParam("pageIndex")) || 1;
var rowsInPage = 10;
var total = 0;
var pageCount = 0;

var laypage = null;
var layer = null;
layui.use(["laypage", "layer"], function() {
    laypage = layui.laypage;
    initData();
})


function initData() {
    asyncInvoke("/news.api?pageIndex=" + pageIndex + "&rowsInPage=" + rowsInPage, "GetList", {}, function(d) {
        console.log(d);
        if (d.code) {
            alert(d.data);
            return;
        }
        total = d.data.total;
        pageCount = d.data.pageCount;
        renderProductList(d.data.rows);
        laypage.render({
            elem: 'page',
            count: d.data.pageCount,
            jump: pageChange
        })
    })
}

function pageChange(obj, bool) {
    if (!bool) {
        layer.load(1, { shade: [0.2, '#000000'] });
        asyncInvoke("/news.api?pageIndex=" + obj.curr + "&rowsInPage=10&searchKey=" + searchKey, "GetList", {}, function(d) {
            if (d.code) {
                errHandler(d.data);
            } else {
                renderList(d.data.rows);
            }
            layer.closeAll('loading');
        })
    }
}

function renderProductList(data) {
    var html = "";
    for (var i = 0; i < data.length; i++) {
        var title = data[i].title;
        var time = new Date(data[i].createTime).format("yyyy-MM-dd");
        var id = data[i].id;
        var icon = '<i class="icon"></i>'
        var url = data[i].original ? data[i].textContent + ' target="blank"' : '/newsDetail?id=' + id;
        html += '<li class="clearfix">' + icon + ' <a href=' + url + ' title=' + title + ' class="article-title">' + title + '</a> <span class="article-time">' + time + '</span> </li>'
    }
    $("#list").html(html);
}
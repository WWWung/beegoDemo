var pageIndex = Number.parseInt(getQeuryParam("pageIndex")) || 1;
var rowsInPage = 10;
var total = 0;
var pageCount = 0;


initData();

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
    })
}

function renderProductList(data) {
    var html = "";
    for (var i = 0; i < data.length; i++) {
        var title = data[i].title;
        var time = new Date(data[i].createTime).format("yyyy-MM-dd");
        var id = data[i].id;
        var icon = '<i class="icon"></i>'
        html += '<li class="clearfix">' + icon + ' <a href="/newsDetail?id=' + id + '" title=' + title + ' class="article-title">' + title + '</a> <span class="article-time">' + time + '</span> </li>'
    }
    $("#list").html(html);
}
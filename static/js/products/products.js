var pageIndex = Number.parseInt(getQeuryParam("pageIndex")) || 1;
var rowsInPage = 10;
var total = 0;
var pageCount = 0;

$("#turn-page").on("click", paginate);
$("#prev-page").on("click", toPrevPage);
$("#next-page").on("click", toNextPage);

initData();

function initData() {
    asyncInvoke("/products.api?pageIndex=" + pageIndex + "&rowsInPage=" + rowsInPage, "GetList", null, function(d) {
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
        html += '<li class="clearfix">' + icon + ' <a href="/productDetail?id=' + id + '" title=' + title + ' class="article-title">' + title + '</a> <span class="article-time">' + time + '</span> </li>'
    }
    $("#product-list").html(html);
}

function paginate() {
    var pageNum = Number.parseInt($("#page-num").val());
    if (pageNum !== pageNum) {
        $("#page-num").val("");
        return alert("请输入正确的页码数")
    }
    if (pageNum < 1 || pageNum > pageCount) {
        $("#page-num").val("");
        return alert("请输入正确的页码数")
    }
    window.location.href = "/products?pageIndex=" + pageNum;
}

function toPrevPage() {
    if (pageIndex - 1 < 1) {
        return alert("已经是第一页了")
    }
    window.location.href = "/products?pageIndex=" + (pageIndex - 1);
}

function toNextPage() {
    if (pageIndex + 1 > pageCount) {
        return alert("已经是最后一页了")
    }
    window.location.href = "/products?pageIndex=" + (pageIndex + 1);
}
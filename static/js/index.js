getProducts();
getNews();
getFriends();
getUtils();
getGuiFan();
getOption();

//  产品中心
function getProducts() {
    asyncInvoke("/products.api?pageIndex=1&rowsInPage=5&sort=sort", "GetList", null, function(d) {
        renderProductList(d.data.rows);
    })
}

//  新闻中心
function getNews() {
    asyncInvoke("/news.api?pageIndex=1&rowsInPage=5&sort=sort", "GetList", null, function(d) {
        renderNewsList(d.data.rows);
    })
}

//  新闻中心
function getFriends() {
    asyncInvoke("/friendUrl.api?pageIndex=1&rowsInPage=5&sort=sort", "GetList", null, function(d) {
        renderFriendList(d.data.rows);
    })
}

//  技术规范
function getGuiFan() {
    asyncInvoke("/filesManage.api?pageIndex=1&rowsInPage=10&sort=sort&searchKey=0", "GetList", null, function(d) {
        renderGuiFanList(d.data.rows);
    })
}

//  系统工具
function getUtils() {
    asyncInvoke("/filesManage.api?pageIndex=1&rowsInPage=10&sort=sort&searchKey=1", "GetList", null, function(d) {
        console.log(d)
        renderUtilsList(d.data.rows);
    })
}

//  操作说明
function getOption() {
    asyncInvoke("/filesManage.api?pageIndex=1&rowsInPage=10&sort=sort&searchKey=2", "GetList", null, function(d) {
        renderOptionList(d.data.rows);
    })
}


//  产品中心列表
function renderProductList(data) {
    var html = "";
    for (var i = 0; i < data.length; i++) {
        var title = data[i].title;
        var time = new Date(data[i].createTime).format("yyyy-MM-dd");
        var id = data[i].id;
        html += '<li> <a href="/productDetail?id=' + id + '" title=' + title + ' class="article-title">' + title + '</a> <span class="article-time">' + time + '</span> </li>'
    }
    $("#product-list").html(html);
}

//  新闻中心列表
function renderNewsList(data) {
    var html = "";
    for (var i = 0; i < data.length; i++) {
        var title = data[i].title;
        var time = new Date(data[i].createTime).format("yyyy-MM-dd");
        var id = data[i].id;
        html += '<li> <a href="/newsDetail?id=' + id + '" title=' + title + ' class="article-title">' + title + '</a> <span class="article-time">' + time + '</span> </li>'
    }
    $("#news-list").html(html);
}

//  友情链接列表
function renderFriendList(data) {
    var html = "";
    for (var i = 0; i < data.length; i++) {
        var name = data[i].name;
        var url = data[i].url;
        html += '<li> <a target="blank" href="' + url + '" title=' + name + ' class="article-title">' + name + '</a> </li>'
    }
    $("#friend-list").html(html);
}


//  技术文档
function renderGuiFanList(data) {
    var html = "";
    for (var i = 0; i < data.length; i++) {
        var name = data[i].name;
        var time = new Date(data[i].uploadTime).format("yyyy-MM-dd");
        var url = data[i].url;
        var suffix = data[i].extName;
        html += "<li><a href='javascript:;' onclick=downloadFile('" + url + "','" + name + "','" + suffix + "','" + data[i].rank + "')>" + name + "</a><span class='article-time'>" + time + "</span></li>"
    }
    $("#GuiFan-list").html(html);
}

//  技术文档
function renderUtilsList(data) {
    var html = "";
    for (var i = 0; i < data.length; i++) {
        var name = data[i].name;
        var time = new Date(data[i].uploadTime).format("yyyy-MM-dd");
        var url = data[i].url;
        var suffix = data[i].extName;
        html += "<li><a href='javascript:;' onclick=downloadFile('" + url + "','" + name + "','" + suffix + "','" + data[i].rank + "')>" + name + "</a><span class='article-time'>" + time + "</span></li>"
    }
    $("#utils-list").html(html);
}

//  技术文档
function renderOptionList(data) {
    var html = "";
    for (var i = 0; i < data.length; i++) {
        var name = data[i].name;
        var time = new Date(data[i].uploadTime).format("yyyy-MM-dd");
        var url = data[i].url;
        var suffix = data[i].extName;
        html += "<li><a href='javascript:;' onclick=downloadFile('" + url + "','" + name + "','" + suffix + "','" + data[i].rank + "')>" + name + "</a><span class='article-time'>" + time + "</span></li>"
    }
    $("#option-list").html(html);
}

function downloadFile(url, dName, suffix, rank) {
    asyncInvoke("/filesManage.api", "CheckPower", {}, function(d) {
        if (d.code) {
            alert(d.data);
            return;
        }
        if (d.data === 0) {
            if (confirm("未登录无法下载，是否跳转至登录界面？")) {
                window.location.href = "/login";
            }
        } else if (d.data < rank) {
            alert("该资源需要黄金会员权限才能下载,如需升级权限，请联系客服人员")
        } else {
            var arr = url.split("/");
            if (arr.length < 4) {
                alert("出现错误");
                return;
            }
            var u = "/filesManage.api?m=Download&type=" + arr[2] + "&name=" + arr[3] + "&dName=" + encodeURIComponent(dName) + "&suffix=" + suffix + "&url=" + url;
            var a = document.createElement("a");
            a.href = u;
            a.click();
        }
    })
}
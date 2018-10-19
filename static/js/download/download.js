$("#tree li").on("click", function() {
    window.location.hash = "#" + $(this).attr("id");
    initData();
})
var layer;
var laypage;
var searchKey = 0;

layui.use(['layer', 'laypage'], function() {
    layer = layui.layer;
    laypage = layui.laypage;
    initData();
})

function initData() {
    var hash = window.location.hash;
    if (hash === "#utils") {
        searchKey = 1;
        selectLi(1);
    } else if (hash === "#notice") {
        searchKey = 2;
        selectLi(2);
    } else if (hash === "#report") {
        searchKey = 3;
        selectLi(3);
    } else {
        searchKey = 0;
        selectLi(0);
    }
    layer.load(1, { shade: [0.2, '#000000'] });
    asyncInvoke("/filesManage.api?pageIndex=1&rowsInPage=10&searchKey=" + searchKey, "GetList", {}, function(d) {
        if (d.code) {
            errHandler(d.data);
        } else {
            renderList(d.data.rows);
            laypage.render({
                elem: 'page',
                count: d.data.pageCount,
                jump: pageChange
            })
        }
        layer.closeAll('loading');
    })
}

function pageChange(obj, bool) {
    if (!bool) {
        layer.load(1, { shade: [0.2, '#000000'] });
        asyncInvoke("/filesManage.api?pageIndex=" + obj.curr + "&rowsInPage=10&searchKey=" + searchKey, "GetList", {}, function(d) {
            if (d.code) {
                errHandler(d.data);
            } else {
                renderList(d.data.rows);
            }
            layer.closeAll('loading');
        })
    }
}

function selectLi(index) {
    var list = $("#tree").find("li");
    for (var i = 0; i < list.length; i++) {
        list.eq(i).css("background", "");
        list.eq(i).find("a").eq(0).css({
            color: "",
            "font-weight": ""
        })
    }
    list.eq(index).css("background", "linear-gradient(135deg, #95acf4, #0837d1)");
    var a = list.eq(index).find("a").eq(0);
    a.css({
        color: "#fff",
        "font-weight": "bold"
    })
    $("#item-title").text(a.text())
}

function renderList(data) {
    var html = "";
    if (!data.length) {
        $("#list").html("<li>暂无</li>");
        return
    }
    for (var i = 0; i < data.length; i++) {
        var name = data[i].name;
        var time = new Date(data[i].uploadTime).format("yyyy-MM-dd");
        var id = data[i].id;
        var url = data[i].url;
        var suffix = data[i].extName

        html += "<li><i class='icon-turnTo'></i>"
        html += "<span title=" + name + " class='article-title'>" + name + "</span>";
        html += "<time clss='article-time'>" + time + "</time>";
        html += "<a href='javascript:;' onclick=downloadFile('" + url + "','" + name + "','" + suffix + "','" + data[i].rank + "')>下载</a>"
        html += "<a href='javascript:;' onclick=preview('" + id + "','" + suffix + "')>预览</a>"
        html += "</li>";
    }
    $("#list").html(html);
}

function preview(id, suffix) {
    if (suffix !== "pdf") {
        alert("该文件不支持预览！")
        return;
    }
    window.open("http://localhost:8080/downloadCenter/preview?id=" + id);
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
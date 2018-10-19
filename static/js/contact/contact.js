$("#tree li").on("click", function() {
    window.location.hash = "#" + $(this).attr("id");
    initData();
})
$("#submit").on("click", submit)
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
    if (hash === "#phone") {
        selectLi(1);
        $("#list-wrap").css("display", "none");
        $("#online-wrap").css("display", "");

    } else if (hash === "#online" || hash === "") {
        selectLi(0);
        $("#list-wrap").css("display", "");
        $("#online-wrap").css("display", "none");
        layer.load(1, { shade: [0.2, '#000000'] });
        asyncInvoke("/contact.api?pageIndex=1&rowsInPage=10", "GetList", {}, function(d) {
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
}

function pageChange(obj, bool) {
    if (!bool) {
        layer.load(1, { shade: [0.2, '#000000'] });
        asyncInvoke("/contact.api?pageIndex=" + obj.curr + "&rowsInPage=10&searchKey=" + searchKey, "GetList", {}, function(d) {
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
        html = "<li>暂无</li>"
    }
    for (var i = 0; i < data.length; i++) {
        var name = data[i].userName;
        var content = data[i].content;
        var time = timeFormatter(data[i].time);
        html += "<li>";
        var p1 = "<p class=\"user-name\">" + name + "</p>";
        html += p1;
        var p2 = "<p class=\"content-wrap\">" + content + "</p>";
        html += p2;
        var p3 = "<p class=\"time\"><span>" + time + "</span></p>"
        html += p3;
        html += "</li>";
    }
    $("#list").html(html);
}

function submit() {
    var data = {};
    getStringFromInput("content", data);
    if (data.content.length > 300) {
        return alert("输入内容过长")
    }
    data.content = data.content.replace(/</ig, '&lt').replace(/>/ig, '&gt').replace(/\n/ig, '<br/>');
    itemJson = JSON.stringify(data)
    asyncInvoke("/contact.api", "Add", { data: itemJson }, function(d) {
        if (d.code) {
            alert(d.data);
        } else {
            initData();
            $("#content").val("");
        }
    })
}
$("#tree li").on("click", function() {
    window.location.hash = "#" + $(this).attr("id");
    // initData();
})
var layer;
var laypage;
var type = 0;

layui.use(['layer'], function() {
    layer = layui.layer;
    laypage = layui.laypage;
    initData();
})

function initData() {
    var hash = window.location.hash;
    if (hash === "#system") {
        type = 1;
        selectLi(1);
    } else {
        type = 0;
        selectLi(0);
    }
    layer.load(1, { shade: [0.2, '#000000'] });
    asyncInvoke("/summary.api?update=true&pageIndex=1&rowsInPage=10", "GetList", { type: type }, function(d) {
        if (d.code) {
            errHandler(d.data);
        } else {
            render(d.data.rows[0]);
        }
        layer.closeAll('loading');
    })
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

function render(data) {
    if (!data) {
        data = {
            title: "",
            clickNumber: "",
            time: "",
            htmlContent: ""
        }
    }
    $("#article-title").text(data.title);
    $("#clickNumber").text(data.clickNumber);
    $("#time").text(timeFormatter(data.time));
    $("#article-content").html(data.htmlContent);
}
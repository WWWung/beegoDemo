$("#logout").on("click", logout);
//  初始化悬浮窗
initAside()
$("#close-aside").on("click", function() {
    $("#contact-aside").hide();
})

layui.use('carousel', function() {
    var carousel = layui.carousel;
    var arr = ["../static/img/bg2.jpg", "../static/img/bg3.jpg"];
    var html = "";
    for (var i = 0; i < arr.length; i++) {
        html += "<div><img src='" + arr[i] + "' alt=''></div>"
    }
    $("#banner").html(html);
    carousel.render({
        elem: "#banner-wrap",
        width: "100%",
        height: "100%",
        anim: "fade"
    })
})

function logout() {
    asyncInvoke("user.api", "Logout", {}, function(d) {
        if (d.code) {
            errHandler(d.data);
        } else {
            window.location.reload();
        }
    })
}

//==
$(".hh-nav-sub").on("mouseenter", showNavList);
$(".hh-nav-sub").on("mouseleave", hideNavList);

function showNavList(e) {
    $(this).find(".hh-sub-list").stop(true, true);
    $(this).find(".hh-sub-list").eq(0).show();
    $(this).find(".hh-sub-list").eq(0).animate({
        "opacity": 1
    }, 200, "linear")
}

function hideNavList(e) {
    $(this).find(".hh-sub-list").stop(true, true);
    $(this).find(".hh-sub-list").eq(0).animate({
        "opacity": 0
    }, 200, "linear", function() {
        $(this).find(".hh-sub-list").eq(0).hide();
    })
}

function initAside() {
    var html = "<div id=\"contact-aside\">\n    <div id=\"contact-aside-title\">\n        \u5728\u7EBF\u5BA2\u670D\n        <i id=\"close-aside\"></i>\n    </div>\n    <ul id=\"qq-list\">\n        <li>\n            <div class=\"qq-list-row clearfix\">\n                <i class=\"service-phone-icon\"></i>\n                <span class=\"service-phone\">15972367296</span>\n            </div>\n            <a href=\"tencent://message/?uin=1281975873&Site=%E5%9C%A8%E7%BA%BF%E5%AE%A2%E6%9C%8D&Menu=yes\">\u5728\u7EBF\u4EA4\u8C08</a>\n        </li>\n        <li>\n            <div class=\"qq-list-row clearfix\">\n                <i class=\"service-phone-icon\"></i>\n                <span class=\"service-phone\">15623148932</span>\n            </div>\n            <a href=\"tencent://message/?uin=917816882&Site=%E5%9C%A8%E7%BA%BF%E5%AE%A2%E6%9C%8D&Menu=yes\">\u5728\u7EBF\u4EA4\u8C08</a>\n        </li>\n    </ul>\n</div>"
    $("body").append(html);
}

window.onhashchange = function() {
    window.location.reload()
}
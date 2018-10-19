$("#logout").on("click", logout);

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
    // var name = $("#login-wrap").find("strong").eq(0).val();
    // console.log(name);
    asyncInvoke("user.api", "Logout", {}, function(d) {
        console.log(d);
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
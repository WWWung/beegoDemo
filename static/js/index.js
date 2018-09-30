var currentBanner = 0;

var bannerTime = null;
getBanner();

$(".sub-item").on("mouseover", showNavList);
$(".sub-item").on("mouseout", hideNavList);
$("#banner-nav").on("click", function(e) {
    if (e.target.nodeName.toLowerCase() === "a") {
        clickBanner(e)
    }
})

function showNavList(e) {
    $(e.target).next().show();
    $(e.target).next().animate({
        "opacity": 1
    }, 200, "linear")
}

function hideNavList(e) {
    $(e.target).next().animate({
        "opacity": 0
    }, 200, "linear", function() {
        $(e.target).next().hide();
    })
}

//  生成轮播图
function getBanner(srcArr) {
    if (!srcArr || !srcArr.length) {
        srcArr = ["../static/img/bg2.jpg", "../static/img/bg3.jpg"];
    }
    var imgs = "";
    var btns = "";
    for (var i = 0; i < srcArr.length; i++) {
        var img = "<li><img src='" + srcArr[i] + "'></li>";
        imgs += img;
        var btn = "<a href='javascript:;'></a>";
        btns += btn;
    }
    $("#home-banner-list").html(imgs);
    $("#banner-nav").html(btns);
    bannerStart();
}

//  开始轮播
function bannerStart() {
    // clearBanner();
    $("#home-banner-list li").eq(currentBanner).show();
    $("#home-banner-list li").eq(currentBanner).css("opacity", 1);
    $("#banner-nav a").eq(currentBanner).addClass("active");
    bannerTime = setInterval(function() {
        bannerBroadcast();
    }, 3000);
}

//  轮播
function bannerBroadcast() {
    var els = $("#home-banner-list li");
    var currentEl = els.eq(currentBanner);
    var nextEl = null;
    if (currentBanner + 1 >= els.length) {
        nextEl = els.eq(0);
    } else {
        nextEl = els.eq(currentBanner + 1);
    }
    currentEl.animate({
        opacity: 0
    }, 500, "linear", function() {
        currentEl.hide();
    })
    nextEl.show();
    nextEl.animate({
        opacity: 1
    }, 500, "linear", function() {
        currentBanner++;
        if (currentBanner >= els.length) {
            currentBanner = 0;
        }
        $("#banner-nav a").removeClass("active");
        $("#banner-nav a").eq(currentBanner).addClass("active");
    });
}

function clickBanner(e) {
    clearInterval(bannerTime);
    var el = $(e.target);
    currentBanner = $("#banner-nav a").index(el) - 1;
    if (currentBanner < 0) {
        currentBanner = $("#banner-nav a").length - 1;
    }
    bannerBroadcast();
    bannerStart();
}

function clearBanner() {
    $("#banner-nav a").removeClass("active");
    $("#home-banner-list li").hide();
    $("#home-banner-list li").css("opacity", 0);
}
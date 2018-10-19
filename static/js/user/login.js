$("#login").on("click", login)
var data;

function login() {
    getDataFromDom();
    if (!data.name) {
        alert("请输入账号");
        return;
    }
    if (!data.password) {
        alert("请输入密码");
        return;
    }
    if (!data.captcha) {
        alert("请输入验证码");
        return;
    }
    data.captchaId = $("#captcha-id").val();
    asyncInvoke("/user.api", "Login", data, function(d) {
        if (d.code) {
            errHandler(d.data);
            getNewCaptcha();
        } else {
            if (document.referrer.indexOf("21zkb") > 0) {
                window.location.href = document.referrer;
            } else {
                window.location.href = "/";
            }
        }
    })
}

function captchaReload() {
    var src = $("#captcha-image").attr("src");
    var i = src.indexOf("?");
    if (i >= 0) {
        src = src.substr(0, i);
    }
    src = src + "?reload=" + new Date().getTime();
    $("#captcha-image").attr("src", src);
    return false;
}

function getNewCaptcha() {
    $.ajax({
        type: "post",
        url: "/captcha",
        success: function(d) {
            if (d.code) {
                alert(d.data);
            } else {
                $("#captcha-image").attr("src", "/captcha/" + d.captchaID + ".png");
                $("#captcha-id").val(d.captchaID);
            }
        }
    })
}

function getDataFromDom() {
    if (!data) {
        data = {};
    }
    getStringFromInput("name", data);
    getStringFromInput("password", data);
    getStringFromInput("captcha", data);
}
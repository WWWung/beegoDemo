// $("#captcha-image").on("click", captchaReload);
$("#regis").on("click", regis);

function regis() {
    var name = $("#name").val();
    var password = $("#password").val();
    if (!name) {
        alert("请输入用户名");
        return
    }
    if (!password) {
        alert("请输入密码");
        return
    }
    var captchaId = $("#captcha-id").val();
    var captcha = $("#captcha").val();
    if (!captcha) {
        alert("请输入验证码");
        return
    }
    if (window.location.href.indexOf("localhost") < 0) {
        alert("环境错误");
        return;
    }
    var data = {
        captchaId: captchaId,
        captcha: captcha,
        name: name,
        password: password,
        power: 0
    }
    $.ajax({
        type: "post",
        url: "/admin/regis",
        data: data,
        success: function(d) {
            if (d.code) {
                alert(d.data);
                getNewCaptcha();
            } else {
                alert("注册成功");
                window.location.href = "/admin";
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
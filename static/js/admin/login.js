$("body").on("keydown", function(e) {
    if (e.keyCode === 13) {
        loginin()
    }
})
$("#loginin").on("click", loginin);
$("#captcha-image").on("click", captchaReload)

function loginin() {
    var name = $("#name").val();
    var password = $("#password").val();
    var captcha = $("#captcha").val();
    if (!name) {
        alert("请输入账号名");
        return
    }
    if (!password) {
        alert("请输入密码");
        return
    }
    if (!captcha) {
        alert("请输入验证码");
        return
    }
    var captchaId = $("#captcha-id").val();
    var data = {
        name: name,
        password: password,
        captcha: captcha,
        captchaId: captchaId
    }
    $.ajax({
        type: "post",
        data: data,
        url: "/admin/loginin",
        success: function(d) {
            if (d.code) {
                alert(d.data)
                getNewCaptcha();
            } else {
                window.location.href = "/admin"
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
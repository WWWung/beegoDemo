$("#regis").on("click", regis);
var data;

function regis() {
    getDataFromDom();
    if (!checkData()) {
        return;
    };
    var itemJson = JSON.stringify(data);
    var d = {
        data: itemJson
    }
    getStringFromInput("captcha", d);
    getStringFromInput("password", d);
    d.captchaId = $("#captcha-id").val();
    asyncInvoke("/user.api", "Add", d, function(d) {
        if (d.code) {
            errHandler(d.data);
            getNewCaptcha();
        } else {
            alert("注册成功!");
            window.location.href = "/";
        }
    })
}

function getDataFromDom() {
    if (!data) {
        data = {};
    }
    getStringFromInput("name", data);
    getStringFromInput("password", data);
    getStringFromInput("phone", data);
    getStringFromInput("dwPhone", data);
    getStringFromInput("addr", data);
}

function checkData() {
    if (!data.name) {
        alert("请输入名字");
        return false;
    }
    if (!data.password) {
        alert("请输入密码");
        return false;
    }
    if (data.password !== $("#rePassword").val()) {
        alert("两次密码不一致");
        return false;
    }
    if (!checkPassWord(data.password)) {
        alert("密码由字母和数字组成且长度不小于6位");
        return false;
    }
    if (!data.phone) {
        alert("请输入手机号码");
        return false;
    }
    if (!data.dwPhone) {
        alert("请输入单位电话");
        return false;
    }
    if (!data.addr) {
        alert("请输入单位地址");
        return false;
    }
    if (!checkPhone(data.phone)) {
        alert("请输入有效手机号码");
        return false;
    }
    return true;
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
<!DOCTYPE html>
<html>

<head>
    <meta charset="utf-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>质控宝后台注册</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="../../static/css/clear.css">
    <link rel="stylesheet" type="text/css" media="screen" href="../../static/css/admin/register.css" />
    <script src="../../static/js/jq-min.js"></script>
</head>

<body>
    <div id="admin-regis-box">
        <div id="al-wrap" class="form">
            <div id="al-title">
                <h1>上海路岩科技有限公司</h1>
                <img src="../../static/img/logo.png" alt="" id="logo">
            </div>
            <p>后台管理账号注册</p>
            <div class="input-row">
                <input type="text" placeholder="用户名" id="name">
            </div>
            <div class="input-row">
                <input type="password" placeholder="密码" id="password">
            </div>
            <div class="input-row">
                <input type="text" placeholder="验证码" id="captcha">
            </div>
            <p><img id="captcha-image" src="/captcha/{{.CaptchaId}}.png" onclick="captchaReload()" alt="Captcha image"></p>
            <input type=hidden name=captchaId value="{{.CaptchaId}}" id="captcha-id">
            <div>
                <a href="javascript:;" id="regis">注册</a>
            </div>
        </div>
    </div>
    <script src="../../static/js/admin/register.js"></script>
</body>

</html>
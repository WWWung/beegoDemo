<!DOCTYPE html>
<html>

<head>
    <meta charset="utf-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>登录</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="../../static/layui/css/layui.css">
    <link rel="stylesheet" href="../../static/css/clear.css">
    <link rel="stylesheet" href="../../static/css/global.css">
    <link rel="stylesheet" href="../../static/css/header.css">
    <link rel="stylesheet" href="../../static/css/footer.css">
    <link rel="stylesheet" href="../../static/css/user/login.css" />
    <script src="../../static/js/jq-min.js"></script>
    <script src="../../static/layui/layui.all.js"></script>
    <script src="../../static/js/global.js"></script>
</head>

<body>
    <div id="root">
        <header id="home-head" class="clearfix">
            <div id="home-logo">
                <a href="/"></a>
            </div>
            <ul id="home-head-nav">
                <li class="hh-nav-sub"> <a href="/" class="sub-item">首页</a> </li>
                <li class="hh-nav-sub"> <a href="javascript:;" class="sub-item">关于我们</a>
                    <ul class="hh-sub-list">
                        <li> <a href="javascript:;">企业简介</a> </li>
                        <li> <a href="javascript:;">系统简介</a> </li>
                    </ul>
                </li>
                <li class="hh-nav-sub"> <a href="javascript:;" class="sub-item">产品中心</a>
                    <ul class="hh-sub-list">
                        <li> <a href="javascript:;">产品1</a> </li>
                        <li> <a href="javascript:;">产品2</a> </li>
                    </ul>
                </li>
                <li class="hh-nav-sub"> <a href="javascript:;" class="sub-item">新闻中心</a>
                    <ul class="hh-sub-list">
                        <li> <a href="javascript:;">新闻1</a> </li>
                        <li> <a href="javascript:;">新闻2</a> </li>
                    </ul>
                </li>
                <li class="hh-nav-sub"> <a href="javascript:;" class="sub-item">案例展示</a>
                    <ul class="hh-sub-list">
                        <li> <a href="javascript:;">客户1</a> </li>
                        <li> <a href="javascript:;">客户2</a> </li>
                    </ul>
                </li>
                <li class="hh-nav-sub"> <a href="/downloadCenter" class="sub-item">下载中心</a>
                    <ul class="hh-sub-list">
                        <li> <a href="/downloadCenter#doc">技术文档</a> </li>
                        <li> <a href="/downloadCenter#utils">系统工具</a> </li>
                        <li> <a href="/downloadCenter#notice">操作说明</a> </li>
                        <li> <a href="/downloadCenter#report">模拟报告</a> </li>
                    </ul>
                </li>
                <li class="hh-nav-sub"> <a href="javascript:;" class="sub-item">联系我们</a>
                    <ul class="hh-sub-list">
                        <li> <a href="javascript:;">联系方式</a> </li>
                        <li> <a href="javascript:;">在线留言</a> </li>
                    </ul>
                </li>
            </ul>
            <div id="login-wrap">
                {{str2html .tip}}
            </div>

        </header>
        <div id="main">
            <div class="form">
                <div class="row clearfix">
                    <div class="row-dt">
                        <span>账号:</span>
                    </div>
                    <div class="row-dd">
                        <input id="name" type="text" class="layui-input" placeholder="请输入账号">
                    </div>
                </div>
                <div class="row clearfix">
                    <div class="row-dt">
                        <span>密码:</span>
                    </div>
                    <div class="row-dd">
                        <input id="password" type="password" class="layui-input" placeholder="请输入密码">
                    </div>
                </div>
                <div class="row clearfix">
                    <div class="row-dt">
                        <span>验证码:</span>
                    </div>
                    <div class="row-dd">
                        <input id="captcha" type="text" class="layui-input" placeholder="请输入验证码">
                    </div>
                </div>
                <div class="row">
                    <img id="captcha-image" src="/captcha/{{.CaptchaId}}.png" onclick="captchaReload()" alt="Captcha image">
                    <input type=hidden name=captchaId value="{{.CaptchaId}}" id="captcha-id">
                </div>
                <div class="row">
                    <a href="/register" class="small-button edit" id="">注册</a>
                    <a href="javascript:;" class="small-button edit" id="login">登录</a>
                </div>
            </div>
        </div>
        <footer id="root-foot">
            <div class="foot-row">
                <span>电话:021-80370982</span>
                <span>邮编:200216</span>
            </div>
            <div class="foot-row">
                <span>单位：上海路岩信息科技有限公司</span>
            </div>
            <div class="foot-row">
                <span>地址：上海浦东新区历城路70号3071A室</span>
            </div>
        </footer>
    </div>
    <script src="../../static/js/header.js"></script>
    <script src="../../static/js/user/login.js"></script>
</body>

</html>
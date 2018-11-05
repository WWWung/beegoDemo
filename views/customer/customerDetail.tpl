<!DOCTYPE html>
<html>

<head>
    <meta charset="utf-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>Page Title</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="../../static/css/clear.css">
    <link rel="stylesheet" href="../../static/layui/css/layui.css">
    <link rel="stylesheet" href="../../static/css/global.css">
    <link rel="stylesheet" href="../../static/css/header.css">
    <link rel="stylesheet" href="../../static/css/footer.css">
    <link rel="stylesheet" href="../../static/css/customer/customerDetail.css" />
    <script src="../../static/js/jq-min.js"></script>
    <script src="../../static/layui/layui.all.js"></script>
    <script src="../../static/js/global.js"></script>
</head>

<body>
    <div id="root">
        <div>
            <header id="home-head" class="clearfix">
                <div id="home-logo">
                    <a href="/"></a>
                </div>
                <ul id="home-head-nav">
                    <li class="hh-nav-sub"> <a href="/" class="sub-item">首页</a> </li>
                    <li class="hh-nav-sub"> <a href="/summary" class="sub-item">关于我们</a>
                        <ul class="hh-sub-list">
                            <li> <a href="/summary#company">企业简介</a> </li>
                            <li> <a href="/summary#system">系统简介</a> </li>
                        </ul>
                    </li>
                    <li class="hh-nav-sub"> <a href="/products" class="sub-item">产品中心</a>
                    </li>
                    <li class="hh-nav-sub"> <a href="/news" class="sub-item">新闻中心</a>
                    </li>
                    <li class="hh-nav-sub"> <a href="/customer" class="sub-item">案例展示</a>
                    </li>
                    <li class="hh-nav-sub"> <a href="/downloadCenter" class="sub-item">下载中心</a>
                        <ul class="hh-sub-list">
                            <li> <a href="/downloadCenter#doc">技术文档</a> </li>
                            <li> <a href="/downloadCenter#utils">系统工具</a> </li>
                            <li> <a href="/downloadCenter#notice">操作说明</a> </li>
                            <li> <a href="/downloadCenter#report">模拟报告</a> </li>
                        </ul>
                    </li>
                    <li class="hh-nav-sub"> <a href="/contact" class="sub-item">联系我们</a>
                        <ul class="hh-sub-list">
                            <li> <a href="/contact#phone">联系方式</a> </li>
                            <li> <a href="/contact#online">在线留言</a> </li>
                        </ul>
                    </li>
                </ul>
                <div id="login-wrap">
                    {{str2html .tip}}
                </div>
            </header>
        </div>
        <div id="main">
            <div id="page-title" class="clearfix">
                <div id="product-logo">
                    <img src="../../static/img/product.png" alt="产品" id="pro-logo-img">
                </div>
                <div id="product-title">
                    <h2><a href="/customer">案例展示</a></h2>
                    <p>Customer</p>
                </div>
            </div>
            <div id="page-content">
                <div id="article-title-wrap">
                    <h3 id="article-title">
                        {{.title}}
                    </h3>
                    <div id="article-info">
                        <span>访问:</span>
                        <span id="clickNumber">{{.clickNumber}}</span>
                        <span>时间:</span>
                        <span id="updateTime">{{.createTime}}</span>
                    </div>
                </div>
                <div id="article-content">
                    {{str2html .htmlContent}}
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
    <script src="../../static/js/customer/customerDetail.js"></script>
</body>

</html>
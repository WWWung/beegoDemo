<!DOCTYPE html>
<html>

<head>
    <meta charset="utf-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>产品中心</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="../../static/layui/css/layui.css">
    <link rel="stylesheet" href="../../static/css/clear.css">
    <link rel="stylesheet" href="../../static/css/global.css">
    <link rel="stylesheet" href="../../static/css/header.css">
    <link rel="stylesheet" href="../../static/css/footer.css">
    <link rel="stylesheet" href="../../static/css/products/products.css" />
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
        <div id="home-banner">
            <div id="home-banner">
                <div class="layui-carousel" id="banner-wrap">
                    <div carousel-item id="banner">
                    </div>
                </div>
            </div>
        </div>
        <div id="list-title">
            <h2>
                <span>产品中心</span>
            </h2>
        </div>
        <div id="main">
            <ul id="product-list">

            </ul>
            <div id="pagination">
                <span>共{{.total}}条{{.pageCount}}页</span>
                <span>现在是第{{.pageIndex}}页</span>
                <a href="/products" id="prev-page">首页</a>
                <a href="/products?pageIndex={{.pageCount}}" id="prev-page">尾页</a>
                <a href="javascript:;" id="prev-page">上一页</a>
                <a href="javascript:;" id="next-page">下一页</a>
                <span>跳转到第</span>
                <input type="text" id="page-num">
                <span>页</span>
                <a href="javascript:;" id="turn-page">确定</a>
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
    <script src="../../static/js/products/products.js"></script>
</body>

</html>
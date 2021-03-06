<!DOCTYPE html>
<html>

<head>
    <meta charset="utf-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>联系我们</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="../../static/layui/css/layui.css">
    <link rel="stylesheet" href="../../static/css/clear.css">
    <link rel="stylesheet" href="../../static/css/global.css">
    <link rel="stylesheet" href="../../static/css/header.css">
    <link rel="stylesheet" href="../../static/css/footer.css">
    <link rel="stylesheet" href="../../static/css/contact/contact.css" />
    <script src="../../static/js/jq-min.js"></script>
    <script src="../../static/layui/layui.all.js"></script>
    <script src="../../static/js/global.js"></script>
</head>

<body>
    <!-- <div id="contact-aside">
        <div id="contact-aside-title">
            在线客服
            <i id="close-aside"></i>
        </div>
        <ul id="qq-list">
            <li>
                <div class="qq-list-row clearfix">
                    <i class="service-phone-icon"></i>
                    <span class="service-phone">15972367296</span>
                </div>
                <a href="tencent://message/?uin=1281975873&Site=%E5%9C%A8%E7%BA%BF%E5%AE%A2%E6%9C%8D&Menu=yes">在线交谈</a>
            </li>
            <li>
                <div class="qq-list-row clearfix">
                    <i class="service-phone-icon"></i>
                    <span class="service-phone">15623148932</span>
                </div>
                <a href="tencent://message/?uin=917816882&Site=%E5%9C%A8%E7%BA%BF%E5%AE%A2%E6%9C%8D&Menu=yes">在线交谈</a>
            </li>
        </ul>
    </div> -->
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
        <div id="home-banner">
            <div id="home-banner">
                <div class="layui-carousel" id="banner-wrap">
                    <div carousel-item id="banner">
                    </div>
                </div>
            </div>
        </div>
        <div id="page-title">
            <h2>
                <i></i>
                <span>联系我们</span>
            </h2>
            <div id="item-title">技术文档</div>
        </div>
        <div id="page-content" class="clearfix">
            <div id="aside">
                <div id="tree">
                    <ul>
                        <li id="online">
                            <a href="javascript:;">在线留言</a>
                        </li>
                        <li id="phone">
                            <a href="javascript:;">联系方式</a>
                        </li>
                    </ul>
                </div>
            </div>
            <div id="article">
                <div id="list-wrap">
                    <ul id="list">

                    </ul>
                    <div id="page"></div>
                    <div id="word-wrap">
                        <div id="textarea-wrap">
                            <textarea name="desc" placeholder="请输入内容(300字以内)" class="layui-textarea" id="content"></textarea>
                        </div>
                        <div>
                            <a href="javascript:;" class="small-button edit" id="submit">提交</a>
                        </div>
                    </div>
                </div>
                <div id="online-wrap">
                    <div class="row">
                        <p class="row-title">合作咨询:</p>
                        <p>
                            <span>021-80370982</span>
                            <span>15972367296</span>
                            <span>15623148932</span>
                        </p>
                    </div>
                    <div class="row">
                        <p class="row-title">技术服务:</p>
                        <p>
                            <span>18502748856</span>
                            <span>16607216366</span>
                            <span>13886672332</span>
                        </p>
                    </div>
                    <div class="row">
                        <p class="row-title">邮政编码:</p>
                        <p>
                            <span>021-80370982</span>
                        </p>
                    </div>
                    <div class="row">
                        <p class="row-title">单位地址:</p>
                        <p>上海浦东新区历城路70号3071A室</p>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <script src="../../static/js/header.js"></script>
    <script src="../../static/js/contact/contact.js"></script>
</body>

</html>
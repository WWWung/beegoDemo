<!DOCTYPE html>
<html>

<head>
    <meta charset="utf-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>质控宝-后台管理</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" type="text/css" media="screen" href="../../static/layui/css/layui.css" />
    <script src="../../static/layui/layui.js"></script>
</head>

<body class="layui-layout-body">
    <div class="layui-layout layui-layout-admin">
        <!-- 头部 -->
        <div class="layui-header">
            <div class="layui-logo">质控宝-后台管理</div>
            <ul class="layui-nav layui-layout-right">
                <li class="layui-nav-item">
                    {{.Name}}
                </li>
                <li class="layui-nav-item">
                    <a href="javascript:;">退了</a>
                </li>
            </ul>
        </div>
        <!-- 侧边栏 -->
        <div class="layui-side layui-bg-black">
            <div class="layui-side-scroll">
                <ul class="layui-nav layui-nav-tree" id="tree">
                    <li class="layui-nav-item" lay-id="products"><a href="javascript:;">产品中心</a></li>
                    <li class="layui-nav-item"><a href="javascript:;">产品中心</a></li>
                </ul>
            </div>
        </div>
        <!-- 内容 -->
        <div class="layui-body">
            <div class="layui-tab" lay-allowClose="true" lay-filter="tap">
                <ul class="layui-tab-title">
                </ul>
                <div class="layui-tab-content">
                </div>
            </div>
        </div>
    </div>
    <script src="../../static/js/global.js"></script>
    <script src="../../static/js/admin.js"></script>
</body>

</html>
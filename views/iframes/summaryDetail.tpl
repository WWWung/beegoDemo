<!DOCTYPE html>
<html>

<head>
    <meta charset="utf-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>Page Title</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" type="text/css" media="screen" href="../../static/layui/css/layui.css" />
    <link rel="stylesheet" type="text/css" media="screen" href="../../static/css/global.css" />
    <link rel="stylesheet" href="../../static/css/admin/summaryDetail.css">
    <script src="../../static/js/jq-min.js"></script>
    <script src="../../static/tinymce/tinymce.min.js"></script>
    <script src="../../static/layui/layui.js"></script>
    <script src="../../static/js/global.js"></script>
</head>

<body>
    <div class="wrap layui-form">
        <div class="title-wrap">
            <div class="item-title">
                标题信息
                <p></p>
            </div>
            <div class="layui-form-item">
                <label class="layui-form-label dt">标题</label>
                <div class="layui-input-block dd">
                    <input type="text" id="title" name="title" required lay-verify="required" placeholder="请输入标题" autocomplete="off" class="layui-input">
                </div>
            </div>
            <div class="layui-form-item">
                <label class="layui-form-label dt">更新时间</label>
                <input type="text" class="layui-input items-in-row layui-disabled" disabled id="time">
                <label class="layui-form-label dt">简介类型</label>
                <div class="items-in-row">
                    <select name="type" lay-filter="type" id="type" id="type" class="">
                        <option value="0">企业简介</option>
                        <option value="1">系统简介</option>
                    </select>
                </div>
            </div>
            <div class="layui-form-item">
                <label class="layui-form-label dt">点击数</label>
                <input type="text" name="title" required lay-verify="required" placeholder="点击数" autocomplete="off" class="layui-input items-in-row" id="clickNumber">
                <label class="layui-form-label dt">排序</label>
                <input type="text" name="title" required lay-verify="required" placeholder="1-1000升序排列" autocomplete="off" class="layui-input items-in-row" id="sort">
            </div>
        </div>
        <div class="content-wrap">
            <div class="item-title">
                内容
            </div>
            <div class="textarea-wrap">
                <div id="textarea"></div>
            </div>
        </div>
        <div class="btn-wrap">
            <a href="javascript:;" class="small-button edit" id="save">保存</a>
            <a href="javascript:;" class="small-button delete" id="cancel">返回</a>
        </div>
    </div>
    <script src="../../static/js/admin/summaryDetail.js "></script>
</body>

</html>
<!DOCTYPE html>
<html>

<head>
    <meta charset="utf-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>Page Title</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" type="text/css" media="screen" href="../../static/layui/css/layui.css" />
    <link rel="stylesheet" type="text/css" media="screen" href="../../static/css/global.css" />
    <link rel="stylesheet" href="../../static/css/admin/newsDetail.css">
    <script src="../../static/js/jq-min.js"></script>
    <script src="../../static/tinymce/tinymce.min.js"></script>
    <script src="../../static/layui/layui.js"></script>
    <script src="../../static/js/global.js"></script>
</head>

<body>
    <div class="wrap">
        <div class="title-wrap">
            <div class="item-title">
                标题信息
            </div>
            <div class="layui-form-item">
                <label class="layui-form-label dt">标题</label>
                <div class="layui-input-block dd">
                    <input type="text" id="title" name="title" required lay-verify="required" placeholder="请输入标题" autocomplete="off" class="layui-input">
                </div>
            </div>
            <div class="layui-form-item">
                <label class="layui-form-label dt">新闻来源</label>
                <div class="layui-input-block dd layui-form" lay-filter="form">
                    <select name="original" lay-filter="original" id="original">
                        <option value="0">原创</option>
                        <option value="1">引入</option>
                    </select>
                </div>
            </div>
            <div class="layui-form-item">
                <label class="layui-form-label dt">创建时间</label>
                <input type="text" class="layui-input layui-disabled items-in-row" disabled id="createTime">
                <label class="layui-form-label dt">更新时间</label>
                <input type="text" class="layui-input layui-disabled items-in-row " disabled id="updateTime">
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
    <script src="../../static/js/admin/newsDetail.js "></script>
</body>

</html>
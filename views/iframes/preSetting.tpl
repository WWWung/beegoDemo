<!DOCTYPE html>
<html>

<head>
    <meta charset="utf-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>Page Title</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" type="text/css" media="screen" href="../../static/layui/css/layui.css" />
    <link rel="stylesheet" type="text/css" media="screen" href="../../static/css/global.css" />
    <link rel="stylesheet" href="../../static/css/admin/preSetting.css">
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
                <label class="layui-form-label dt">排序</label>
                <input type="text" name="title" required lay-verify="required" placeholder="1-1000升序排列" autocomplete="off" class="layui-input items-in-row" id="sort">
                <a href="javascript:;" class="small-button func" id="upload">上传文件</a>
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
            <a href="/admin/filesManage" class="small-button edit" id="cancel">返回</a>
            <a href="javascript:;" class="small-button delete" id="del">删除</a>
        </div>
        <div class="imgWrap">
        </div>
    </div>
    <script src="../../static/js/admin/preSetting.js "></script>
</body>

</html>
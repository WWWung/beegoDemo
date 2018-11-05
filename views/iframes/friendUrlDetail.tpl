<!DOCTYPE html>
<html>

<head>
    <meta charset="utf-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>Page Title</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" type="text/css" media="screen" href="../../static/layui/css/layui.css" />
    <link rel="stylesheet" type="text/css" media="screen" href="../../static/css/global.css" />
    <link rel="stylesheet" href="../../static/css/admin/customerDetail.css">
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
                <label class="layui-form-label dt">名称</label>
                <div class="layui-input-block dd">
                    <input type="text" id="name" name="name" required lay-verify="required" placeholder="请输入名称" autocomplete="off" class="layui-input">
                </div>
            </div>
            <div class="layui-form-item">
                <label class="layui-form-label dt">链接</label>
                <div class="layui-input-block dd">
                    <input type="text" id="url" name="url" required lay-verify="required" placeholder="请输入链接地址" autocomplete="off" class="layui-input">
                </div>
            </div>
            <div class="layui-form-item">
                <label class="layui-form-label dt">排序</label>
                <div class="layui-input-block dd">
                    <input type="text" id="sort" name="sort" required lay-verify="required" placeholder="请输入排序" autocomplete="off" class="layui-input">
                </div>
            </div>
        </div>
        <div class="btn-wrap">
            <a href="javascript:;" class="small-button edit" id="save">保存</a>
            <a href="/admin/friendUrl" class="small-button delete" id="cancel">返回</a>
        </div>
    </div>
    <script src="../../static/js/admin/friendUrlDetail.js "></script>
</body>

</html>
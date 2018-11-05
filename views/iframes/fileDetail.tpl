<!DOCTYPE html>
<html>

<head>
    <meta charset="utf-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>文件明细</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" type="text/css" media="screen" href="../../static/layui/css/layui.css" />
    <link rel="stylesheet" type="text/css" media="screen" href="../../static/css/global.css" />
    <link rel="stylesheet" href="../../static/css/admin/fileDetail.css">
    <script src="../../static/js/jq-min.js"></script>
    <script src="../../static/layui/layui.js"></script>
    <script src="../../static/js/global.js"></script>
</head>

<body>
    <div id="wrap">
        <div class="title-wrap layui-form" lay-filter="form">
            <div class="item-title">
                文件信息
            </div>
            <div class="layui-form-item">
                <label class="layui-form-label dt">文件名字</label>
                <div class="layui-input-block dd">
                    <input type="text" id="name" name="name" required lay-verify="required" placeholder="请输入文件名字" autocomplete="off" class="layui-input">
                </div>
            </div>
            <div class="layui-form-item">
                <label class="layui-form-label dt">所需权限</label>
                <div id="rank-wrap">
                    <select name="rank" lay-filter="rank" id="rank">
                        <option value="0">无需权限</option>
                        <option value="1">普通会员</option>
                        <option value="2">黄金会员</option>
                    </select>
                </div>
                <label class="layui-form-label dt">可预览页数</label>
                <input type="text" name="preview" placeholder="可预览页数" autocomplete="off" class="layui-input items-in-row" id="preview">
            </div>
            <div class="layui-form-item">
                <label class="layui-form-label dt">上传时间</label>
                <input type="text" class="layui-input items-in-row" id="uploadTime" placeholder="上传时间">
                <label class="layui-form-label dt">文件大小</label>
                <input type="text" name="size" placeholder="文件大小" disabled autocomplete="off" class="layui-input layui-disabled items-in-row" id="size">
            </div>
            <div class="layui-form-item">
                <label class="layui-form-label dt">下载数量</label>
                <input type="text" name="downloadNumber" required lay-verify="required" placeholder="下载数量" autocomplete="off" class="layui-input items-in-row" id="downloadNumber">

                <label class="layui-form-label dt">排序</label>
                <input type="text" name="sort" required lay-verify="required" placeholder="1-1000升序排列" autocomplete="off" class="layui-input items-in-row" id="sort">
            </div>
            <div class="layui-form-item">
                <label class="layui-form-label dt">文件类型</label>
                <div class="layui-input-block dd">
                    <select name="type" lay-filter="type" id="type" id="type">
                        <option value="">文件所属板块</option>
                        <option value="0">技术文档</option>
                        <option value="1">系统工具</option>
                        <option value="2">操作说明</option>
                        <option value="3">模拟报告</option>
                    </select>
                </div>
            </div>
            <div class="layui-form-item">
                <label class="layui-form-label dt">文件描述</label>
                <div class="layui-input-block dd">
                    <input type="text" name="description" required lay-verify="required" placeholder="文件描述(50个字符以内)" autocomplete="off" class="layui-input" id="description">
                </div>
            </div>
            <div class="layui-form-item">
                <label class="layui-form-label dt">下载地址</label>
                <div class="layui-input-block dd">
                    <input type="text" name="url" disabled placeholder="下载地址(保存后自动生成)" autocomplete="off" class="layui-disabled layui-input" id="url">
                </div>
            </div>
        </div>
        <div id="content">

        </div>
        <div class="btn-wrap">
            <a href="javascript:;" class="small-button func" id="upload">上传文件</a>
            <a href="javascript:;" class="small-button edit" id="save">保存</a>
            <a href="/admin/filesManage" class="small-button delete" id="cancel">返回</a>
        </div>
        <div id="tip">
            <p>仅支持.rar/.doc/.pdf格式,且大小不能超过100M</p>
        </div>
    </div>
    <script src="../../static/js/admin/fileDetail.js"></script>
</body>

</html>
<!DOCTYPE html>
<html>

<head>
    <meta charset="utf-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>Page Title</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="../../static/css/clear.css">
    <link rel="stylesheet" href="../../static/css/download/preview.css">
    <script src="../../static/pdfjs/build/pdf.js"></script>
    <script src="../../static/js/jq-min.js"></script>
    <script src="../../static/js/global.js"></script>
</head>

<body>
    <div id="btn-wrap" class="clearfix">
        <div id="box">
            <div id="pages">
                <a href="javascript::" id="page-up" class="page-btn">上一页</a>
                <a href="javascript::" id="page-down" class="page-btn">下一页</a>
                <input type="text" id="page-to" value="1"> /
                <span id="page-count"></span>
            </div>
            <div id="zoom">
                <span id="zoom-in"></span>
                <span id="zoom-out"></span>
                <input type="text" id="scale" placeholder="40~200">
                <span id="percent">%</span>
            </div>
        </div>
    </div>
    <div id="canvas-wrap">
        <canvas id="canvas"></canvas>
    </div>
    <script src="../../static/js/download/preview.js"></script>
</body>

</html>
$("#zoom-in").on("click", zoomIn);
$("#zoom-out").on("click", zoomOut);

$("#page-up").on("click", toPrev);
$("#page-down").on("click", toNext);

$("#scale").on("blur", setScale);
$("#scale").on("keydown", function(e) {
    if (e.keyCode === 13) {
        setScale();
    }
});

$("#page-to").on("blur", pageChange);
$("#page-to").on("keydown", function(e) {
    if (e.keyCode === 13) {
        pageChange();
    }
});
var id;
var data;
var maxPage = 0;
var pageIndex = 1;
var scale = 1;
initData();

function zoomIn() {
    if (scale <= 2 && scale >= 0.4) {
        scale += 0.1;
        v = 100 * scale;
        $("#scale").val(Math.round(v));
        showPdf();
    }
}

function zoomOut() {
    if (scale <= 2 && scale >= 0.4) {
        scale -= 0.1;
        v = 100 * scale;
        $("#scale").val(Math.round(v));
        showPdf();
    }
}

function toPrev() {
    if (pageIndex > 1) {
        pageIndex--;
        showPdf();
        $("#page-to").val(pageIndex);
    } else {
        alert("已到达第一页");
    }
}

function toNext() {
    if (pageIndex < maxPage) {
        pageIndex++;
        showPdf();
        $("#page-to").val(pageIndex);
    } else {
        alert("已到达可预览最大页数，想了解更多内容可以登录后下载文档");
    }
}

function pageChange() {
    var p = getIntFromInput("page-to");
    if (p < 1) {
        alert("已到达第一页");
        p = 1;
    } else if (p > maxPage) {
        p = maxPage;
        alert("已到达可预览最大页数，想了解更多内容可以登录后下载文档");
    }
    pageIndex = p;
    $("#page-to").val(pageIndex);
    showPdf();
}

function showPdf() {
    // if (pageIndex > Math.max(data.))
    PDFJS.workerSrc = '../../static/pdfjs/build/pdf.worker.js'; //加载核心库
    PDFJS.getDocument(data.url).then(function getPdfHelloWorld(pdf) {
        //
        // 获取第一页数据
        //
        $("#page-count").text(pdf.numPages);
        maxPage = Math.min(data.preview, pdf.numPages);
        pdf.getPage(pageIndex).then(function getPageHelloWorld(page) {
            var viewport = page.getViewport(scale);

            //
            // Prepare canvas using PDF page dimensions
            //
            var canvas = document.getElementById('canvas');
            var context = canvas.getContext('2d');
            canvas.height = viewport.height;
            canvas.width = viewport.width;

            //
            // Render PDF page into canvas context
            //
            var renderContext = {
                canvasContext: context,
                viewport: viewport
            };
            page.render(renderContext);
        });
    });

}

function initData() {
    id = getQeuryParam("id");
    if (id) {
        asyncInvoke("/filesManage.api", "GetItem", { id: id }, function(d) {
            if (d.code) {
                alert(d.data);
            } else {
                data = d.data;
                showPdf();
            }
        })
    }
}

function setScale() {
    var v = getIntFromInput("scale");
    if (v !== v) {
        return;
    }
    if (v < 40) {
        v = 40;
    } else if (v > 200) {
        v = 200;
    } else {
        sub = v - 100;
        v = 100 + Math.round(sub / 10) * 10;
    }
    $("#scale").val(v);
    scale = v / 100;
    showPdf();
}
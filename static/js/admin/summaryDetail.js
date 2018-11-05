$("#save").on("click", save)

var laydate;
var data;
var form;
var id;

layui.use(['laydate', 'form', 'layer'], function() {
    laydate = layui.laydate;
    form = layui.form;
    laydate.render({
        elem: "#time",
        type: "date"
    })
    initData();
})

function initData() {
    id = getQeuryParam("id");
    tinymce.init({
        selector: "#textarea",
        language: "zh_CN",
        plugins: 'link lists image table colorpicker textcolor wordcount contextmenu codesample',
        fontsize_formats: '10px 11px 12px 14px 16px 18px 20px 24px',
        toolbar: 'bold italic underline strikethrough | fontsizeselect | forecolor backcolor | alignleft aligncenter alignright alignjustify | bullist numlist | outdent indent blockquote | undo redo | link unlink image | removeformat | fontselect ',
        font_formats: "宋体=宋体;微软雅黑=微软雅黑;新宋体=新宋体;",
        images_upload_handler: function(blobInfo, success, failure) {
            if (blobInfo.blob().size > 3148576) {
                failure('文件体积过大')
            }
            var formData = new FormData()
            formData.append('file', blobInfo.blob(), blobInfo.filename())
            $.ajax({
                url: "/upload?type=imgs",
                type: 'POST',
                cache: false,
                data: formData,
                processData: false,
                contentType: false,
                beforeSend: function() {
                    uploading = true;
                },
                success: function(d) {
                    success(d.data.src)
                }
            })
        }
    })
    if (id) {
        asyncInvoke("/summary.api", "GetItem", { id: id }, function(d) {
            if (d.code) {
                errHandler(d.data);
                return
            }
            data = d.data;
            setDataToDom();
        })
    }
}

function getDataFromDom() {
    if (!data) {
        data = {};
    }
    getStringFromInput("title", data);
    getStringFromInput("createTime", data);
    getIntFromInput("clickNumber", data);
    getIntFromInput("type", data);
    getIntFromInput("sort", data);
    if (data.createTime) {
        data.time = new Date(data.time);
    } else {
        data.time = new Date("1900-01-01");
    }

    //  获取富文本编辑器里的内容
    var activeEditor = tinymce.activeEditor;
    data.htmlContent = activeEditor.getContent();
    var editBody = activeEditor.getBody();
    activeEditor.selection.select(editBody);
    var text = activeEditor.selection.getContent({ "format": "text" });
    data.textContent = text;
}

function save() {
    getDataFromDom();
    var method = "Add";
    if (id) {
        method = "Update";
    }
    var jsonData = JSON.stringify(data);
    asyncInvoke("/summary.api", method, jsonData, function(d) {
        if (d.code === 0) {
            alert("保存成功")
            window.location.href = "/admin/summaryDetail?id=" + d.data
        } else {
            errHandler(d.data)
        }
    })
}

function setDataToDom() {
    $("#title").val(data.title);
    $("#time").val(new Date(data.time).format("yyyy-MM-dd"));
    $("#clickNumber").val(data.clickNumber);
    form.val("form", {
        "type": data.type
    })
    $("#sort").val(data.sort);

    var activeEditor = tinymce.activeEditor;
    activeEditor.setContent(data.htmlContent);
}
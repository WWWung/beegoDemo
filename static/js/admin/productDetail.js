$("#save").on("click", save)

var laydate;
var data;
var id;

layui.use('laydate', function() {
    laydate = layui.laydate;
    laydate.render({
        elem: "#createTime",
        type: "date"
    })
    laydate.render({
        elem: "#updateTime",
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
        toolbar: 'bold italic underline strikethrough | fontsizeselect | forecolor backcolor | alignleft aligncenter alignright alignjustify | bullist numlist | outdent indent blockquote | undo redo | link unlink image | removeformat | fontselect',
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
        asyncInvoke("/products.api", "GetItem", { id: id }, function(d) {
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
    getStringFromInput("updateTime", data);
    getIntFromInput("clickNumber", data);
    getStringFromInput("brand", data);
    getIntFromInput("sort", data);
    if (data.createTime) {
        data.createTime = new Date(data.createTime);
    } else {
        data.createTime = new Date("1960-01-01");
    }
    if (data.updateTime) {
        data.updateTime = new Date(data.updateTime);
    } else {
        data.updateTime = new Date("1960-01-01");
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
    console.log(data)
        // return
    asyncInvoke("/products.api", method, jsonData, function(d) {
        // console.log(d)
        // return
        if (d.code === 0) {
            alert("保存成功")
            window.location.href = "/admin/productDetail?id=" + d.data
        } else {
            errHandler(d.data)
        }
    })
}

function setDataToDom() {
    $("#title").val(data.title);
    $("#createTime").val(new Date(data.createTime).format("yyyy-MM-dd"));
    $("#updateTime").val(new Date(data.updateTime).format("yyyy-MM-dd"));
    $("#clickNumber").val(data.clickNumber);
    $("#brand").val(data.brand);
    $("#sort").val(data.sort);

    var activeEditor = tinymce.activeEditor;
    activeEditor.setContent(data.htmlContent);
}
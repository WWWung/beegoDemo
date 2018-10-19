$("#save").on("click", save);
$("#upload").on("click", upload);

var laydate;
var layer;
var form;
var data;
var id;
var formData = null;

layui.use(['laydate', 'form', 'layer'], function() {
    laydate = layui.laydate;
    layer = layui.layer;
    form = layui.form;
    laydate.render({
        elem: "#uploadTime",
        type: "date"
    })
    initData();
})

function initData() {
    id = getQeuryParam("id");
    if (id) {
        layer.load(1, { shade: [0.2, '#000000'] });
        asyncInvoke("/filesManage.api", "GetItem", { id: id }, function(d) {
            if (d.code) {
                errHandler(d.data);
            } else {
                data = d.data;
                setDataToDom();
            }
            layer.closeAll("loading")
        })
    }
}

function save() {
    if (!formData && !id) {
        alert("没有上传文件");
        return;
    }
    layer.load(1, { shade: [0.2, '#000000'] });
    var suffix = formData ? getExtName(formData.get('file').name) : data.suffix;
    if (!id) {
        $.ajax({
            url: "/upload?type=" + suffix + "s",
            type: 'POST',
            cache: false,
            data: formData,
            processData: false,
            contentType: false,
            beforeSend: function() {
                uploading = true;
            },
            success: function(d) {
                if (d.code) {
                    errHandler(d.data, layer);
                    layer.closeAll('loading');
                    return;
                }
                getDataFromDom();
                data.url = d.data.src;
                data.extName = suffix;
                var jsonData = JSON.stringify(data);
                asyncInvoke("/filesManage.api", "Add", jsonData, function(d) {
                    if (d.code === 0) {
                        alert("保存成功");
                        window.location.href = "/admin/fileDetail?id=" + d.data
                    } else {
                        errHandler(d.data)
                    }
                    layer.closeAll('loading');
                })
            }
        })
    } else {
        getDataFromDom();
        // data.extName = suffix;
        var jsonData = JSON.stringify(data);
        asyncInvoke("/filesManage.api", "Update", jsonData, function(d) {
            if (d.code === 0) {
                alert("保存成功");
                window.location.href += "?id=" + d.data
            } else {
                errHandler(d.data)
            }
            layer.closeAll('loading');
        })
    }
}

function upload() {
    var input = document.createElement("input");
    input.type = "file";
    $(input).on("change", fileChange)
    input.click();
}

function fileChange() {
    var f = this.files;
    if (!f || !f[0]) {
        return;
    }
    if (!checkExtName(f[0].name)) {
        alert("文件格式不对");
        return
    }
    var reader = new FileReader();
    reader.onload = function(evt) {
        var kbs = evt.total / 1024;
        if (kbs > 102400) {
            alert("文件体积过大");
            return;
        }
        formData = new FormData();
        formData.append('file', f[0], f[0].name);
        var size = calSize(kbs);
        $("#size").val(size);
    }
    reader.readAsDataURL(f[0]);
}

function checkExtName(file_name) {
    var extName = getExtName(file_name);
    return extName === "rar" || extName === "doc" || extName === "pdf";
}

function getExtName(file_name) {
    var index1 = file_name.lastIndexOf(".");
    var index2 = file_name.length;
    var suffix = file_name.substring(index1 + 1, index2);
    return suffix
}

function getDataFromDom() {
    if (!data) {
        data = {};
    }
    getStringFromInput("name", data);
    getIntFromInput("rank", data);
    getIntFromInput("preview", data);
    getIntFromInput("type", data);
    getDateFromInput("uploadTime", data);
    data.size = formData ? formData.get('file').size : data.size ? data.size : 0;
    getIntFromInput("downloadNumber", data);
    getIntFromInput("sort", data);
    getStringFromInput("description", data);
}

function calSize(kbs) {
    var size = 0;
    if (kbs > 10240) {
        size = Math.round(kbs / 1024).toFixed(1) + "MB";
    } else {
        size = kbs.toFixed(1) + "KB";
    }
    return size;
}

function setDataToDom() {
    $("#name").val(data.name);
    form.val("form", {
        "rank": data.rank
    })
    $("#preview").val(data.preview);
    $("#uploadTime").val(new Date(data.uploadTime).format("yyyy-MM-dd"));
    $("#downloadNumber").val(data.downloadNumber);
    $("#sort").val(data.sort);
    form.val("form", {
        "type": data.type
    })
    $("#description").val(data.description);
    $("#url").val(data.url);
    $("#size").val(calSize(data.size / 1024));
}

function checkData() {
    // if (!data.)
}
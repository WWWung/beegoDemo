var id;
var layer

$("#upload").on("click", updateImg)
$("#del").on("click", del)

layui.use(['laydate', 'form', 'layer'], function() {
    // laydate = layui.laydate;
    layer = layui.layer;
    // form = layui.form;
    initData();
})

function initData() {
    id = getQeuryParam("id");
    asyncInvoke("/previewImg.api?rowsInPage=999&pageIndex=1&fileId=" + id, "GetList", {}, function(d) {
        if (d.code === 0) {
            d.data.rows.forEach(function(data) {
                $(".imgWrap").append("<img src=" + data.src + " class='img'>")
            })
        } else {
            errHandler(d.data)
        }
    })
}

function del() {
    asyncInvoke("/previewImg.api", "Delete", { fileId: id }, function(d) {
        if (d.code) {
            alert(d.data)
        } else {
            alert("删除成功")
            window.location.href = "/admin/filesManage"
        }
    })
}

function updateImg() {
    if (!id) {
        return false;
    }
    var input = document.createElement("input");
    input.type = "file";
    console.log(input)
    $(input).on("change", fileChange)
    input.click();
}

function fileChange() {
    layer.load(1, { shade: [0.2, '#000000'] });
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
        if (kbs > 2048) {
            alert("文件大小不能超过2M");
            return;
        }
        var formData = new FormData();
        formData.append('file', f[0], f[0].name);
        var extName = getExtName(f[0].name)
        $.ajax({
            url: "/previewImg.api?m=Add&extName=" + extName + "&sort=" + getIntFromInput("sort") + "&fileId=" + id,
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
                $(".imgWrap").append("<img src=" + d.data + " class='img'>")
                alert("保存成功");
                layer.closeAll('loading');
            }
        })
    }
    reader.readAsDataURL(f[0]);
}

function checkExtName(file_name) {
    var extName = getExtName(file_name);
    return extName === "png" || extName === "jpg" || extName === "jpeg";
}

function getExtName(file_name) {
    var index1 = file_name.lastIndexOf(".");
    var index2 = file_name.length;
    var suffix = file_name.substring(index1 + 1, index2);
    return suffix
}
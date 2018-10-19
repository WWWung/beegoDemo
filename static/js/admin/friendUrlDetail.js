$("#save").on("click", save)

var laydate;
var data;
var id;

layui.use('laydate', function() {
    initData();
})

function initData() {
    id = getQeuryParam("id");
    if (id) {
        asyncInvoke("/friendUrl.api", "GetItem", { id: id }, function(d) {
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
    getStringFromInput("name", data);
    getStringFromInput("url", data);
    getIntFromInput("sort", data);
}

function save() {
    getDataFromDom();
    var method = "Add";
    if (id) {
        method = "Update";
    }
    if (!data.url || !data.sort || !data.name) {
        alert("数据不完整");
        return;
    }
    var jsonData = JSON.stringify(data);
    asyncInvoke("/friendUrl.api", method, jsonData, function(d) {
        if (d.code === 0) {
            alert("保存成功")
            window.location.href = "/admin/friendUrlDetail?id=" + d.data
        } else {
            errHandler(d.data)
        }
    })
}

function setDataToDom() {
    $("#name").val(data.name);
    $("#sort").val(data.sort);
    $("#url").val(data.url);
}
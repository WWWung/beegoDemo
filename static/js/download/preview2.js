var id;
// var layer

// $("#upload").on("click", updateImg)
// $("#del").on("click", del)

// layui.use(['laydate', 'form', 'layer'], function() {
//     // laydate = layui.laydate;
//     layer = layui.layer;
//     // form = layui.form;
//     initData();
// })
$(function() {
    initData();
})

function initData() {
    id = getQeuryParam("id");
    asyncInvoke("/previewImg.api?rowsInPage=999&pageIndex=1&fileId=" + id, "GetList", {}, function(d) {
        if (d.code === 0) {
            d.data.rows.forEach(function(data) {
                $("#imgs").append("<img src=" + data.src + " class='img'>")
            })
        } else {
            errHandler(d.data)
        }
    })
}
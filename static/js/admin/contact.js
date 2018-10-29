var table;
var layer;
var data;
var list;
var columns = [
    [
        { field: 'userName', title: '作者姓名', width: 200, align: "center" },
        { field: 'content', title: '内容', width: 600, align: "center" },
        { field: 'time', title: '时间', width: 120, align: "center" },
        { field: 'edit', title: '删除', width: 200, templet: getHandleBtn, align: "center" }
    ]
]

layui.use(['table', 'layer'], function() {
    table = layui.table;
    layer = layui.layer;
    initData()
})

function initData() {
    list = table.render({
        elem: "#list",
        cols: columns,
        page: true,
        limits: [5, 7, 10],
        limit: 5
    })
    asyncInvoke("/contact.api", "GetList", null, function(d) {
        data = d.data.rows;
        list.reload({
            data: d.data.rows
        })
    })
    return list;
}



function getHandleBtn(d) {
    var _delete = "<a href='javascript:;' class='small-button delete' style='margin:0 5px' onclick=del('" + d.id + "','" + d.LAY_TABLE_INDEX + "')>删除</button>";
    return _delete;
}


function del(id, index) {
    layer.confirm('确认删除？', {
        btn: ['是的', '取消'] //按钮
    }, function() {
        layer.load(1, { shade: [0.2, '#000000'] });
        asyncInvoke("/contact.api", "Delete", { id: id }, function(d) {
            if (d.code) {
                errHandler(d.data);
            } else {
                data.splice(index, 1);
                list.reload({
                    data: data
                });
                alert("删除成功");
            }
            layer.closeAll();
        })
    }, function() {
        // layer.msg('也可以这样', {
        //   time: 20000, //20s后自动关闭
        //   btn: ['明白了', '知道了']
        // });
    });
}
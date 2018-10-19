var table,
    layer,
    list,
    data;
var columns = [
    [
        { field: 'sort', title: '排序', width: 120, align: "center" },
        { field: 'title', title: '标题', width: 400, align: "center" },
        { field: 'type', title: '类型', width: 120, align: "center", templet: getType },
        { field: 'clickNumber', title: '点击量', width: 120, align: "center" },
        { field: 'time', title: '更新日期', width: 120, align: "center" },
        { field: 'edit', title: '编辑', width: 200, templet: getHandleBtn, align: "center" }
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
        // skin: "line",
        page: true,
        limits: [5, 7, 10],
        limit: 5
    })
    asyncInvoke("/summary.api?sort=sort", "GetList", null, function(d) {
        if (d.code) {
            alert(d.data)
        } else {
            data = d.data.rows;
            list.reload({
                data: d.data.rows
            })
        }
    })
    return list;
}

function getType(row) {
    var rsl = "";
    if (!row.type) {
        rsl = "企业简介";
    } else if (row.type === 1) {
        rsl = "系统简介";
    }
    return rsl
}

function getHandleBtn(d) {
    var edit = "<a href='javascript:;' class='small-button edit' style='margin:0 5px' onclick=edit('" + d.id + "')>编辑</a>";
    var _delete = "<a href='javascript:;' class='small-button delete' style='margin:0 5px' onclick=del('" + d.id + "','" + d.LAY_TABLE_INDEX + "')>删除</a>";
    return edit + _delete;
}

function edit(id) {
    window.location.href = "/admin/summaryDetail?id=" + id;
}

function del(id, index) {
    layer.confirm('确认删除？', {
        btn: ['是的', '取消'] //按钮
    }, function() {
        layer.load(1, { shade: [0.2, '#000000'] });
        asyncInvoke("/summary.api", "Delete", { id: id }, function(d) {
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
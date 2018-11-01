var table;
var layer;
var columns = [
    [
        { field: 'sort', title: '排序', width: 120, align: "center" },
        { field: 'name', title: '文件名', width: 400, align: "center" },
        { field: 'type', title: '类型', width: 120, align: "center", templet: getType },
        { field: 'rank', title: '下载所需权限', width: 120, align: "center", templet: getRank },
        { field: 'preview', title: '可预览页数', width: 120, align: "center" },
        { field: 'url', title: '下载地址', width: 120, align: "center" },
        { field: 'uploadTime', title: '创建日期', width: 120, align: "center" },
        { field: 'downloadNumber', title: '下载量', width: 100, align: "center" },
        { field: 'size', title: '文件大小', width: 100, align: "center" },
        { field: 'edit', title: '编辑', width: 200, templet: getHandleBtn, align: "center" }
    ]
]

layui.use(['table', 'layer'], function() {
    table = layui.table;
    layer = layui.layer;
    initData()
})

function initData() {
    var list = table.render({
        elem: "#filesManage-list",
        cols: columns,
        // skin: "line",
        page: true,
        limits: [5, 7, 10],
        limit: 5
    })
    asyncInvoke("/filesManage.api?sort=sort", "GetList", null, function(d) {
        list.reload({
            data: d.data.rows
        })
    })
    return list;
}

function getType(row) {
    var rsl = "";
    if (!row.type) {
        rsl = "技术文档";
    } else if (row.type === 1) {
        rsl = "系统工具";
    } else if (row.type === 2) {
        rsl = "操作说明";
    } else if (row.type === 3) {
        rsl = "模拟报告";
    }
    return rsl
}

function getRank(row) {
    var rsl = "";
    if (!row.rank) {
        rsl = "无需权限";
    } else if (row.rank === 1) {
        rsl = "普通会员";
    } else if (row.rank === 2) {
        rsl = "黄金会员";
    }
    return rsl
}

function getHandleBtn(d) {
    var edit = "<a href='javascript:;' class='small-button edit' style='margin:0 5px' onclick=edit('" + d.id + "')>编辑</button>";
    var _delete = "<a href='javascript:;' class='small-button delete' style='margin:0 5px' onclick=del('" + d.id + "','" + d.LAY_TABLE_INDEX + "')>删除</button>";
    return edit + _delete;
}

function edit(id) {
    window.location.href = "/admin/fileDetail?id=" + id;
}

function del(id, index) {
    layer.confirm('确认删除？', {
        btn: ['是的', '取消'] //按钮
    }, function() {
        layer.load(1, { shade: [0.2, '#000000'] });
        asyncInvoke("/filesManage.api", "Delete", { id: id }, function(d) {
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
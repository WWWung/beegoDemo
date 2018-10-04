var table;
var columns = [
    [
        { field: 'sort', title: '排序', width: 120, align: "center" },
        { field: 'title', title: '标题', width: 400, align: "center" },
        { field: 'brand', title: '品牌', width: 120, align: "center" },
        { field: 'clickNumber', title: '点击量', width: 120, align: "center" },
        { field: 'createTime', title: '创建日期', width: 120, align: "center" },
        { field: 'updateTime', title: '更新日期', width: 120, align: "center" },
        { field: 'edit', title: '编辑', width: 200, templet: getHandleBtn, align: "center" }
    ]
]

layui.use('table', function() {
    table = layui.table;
    initData()
})

function initData() {
    var list = table.render({
        elem: "#products-list",
        cols: columns,
        // skin: "line",
        page: true,
        limits: [5, 7, 10],
        limit: 5
    })
    $.ajax({
        type: 'post',
        url: '/admin/productsList?page=1&index=2',
        success: function(d) {
            console.log(d)
            list.reload({
                data: d.data.rows
            })
        }
    })
    return list;
}

function getHandleBtn(d) {
    var edit = "<a href='javascript:;' class='small-button edit' style='margin:0 5px' onclick=edit('" + d.id + "')>编辑</button>";
    var _delete = "<a href='javascript:;' class='small-button delete' style='margin:0 5px' onclick=del('" + d.id + "')>删除</button>";
    return edit + _delete;
}

function edit(id) {
    // console.log("edit", id);
    window.location.href = "/admin/productDetail?id=" + id;
}

function del(id) {
    console.log("delete", id);
}
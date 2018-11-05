var table;
var layer;
var data;
var columns = [
    [
        { field: 'name', title: '用户ID', width: 400, align: "center" },
        { field: 'phone', title: '联系电话', width: 120, align: "center" },
        { field: 'addr', title: '地址', width: 120, align: "center" },
        { field: 'lastLoginTime', title: '上次登录时间', width: 120, align: "center" },
        { field: 'power', title: '权限', width: 120, align: "center", templet: getRank, edit: "text" },
        { field: 'dwPhone', title: '单位电话', width: 120, align: "center" },
        { field: 'edit', title: '删除', width: 200, templet: getHandleBtn, align: "center" }
    ]
]

layui.use(['table', 'layer'], function() {
    table = layui.table;
    layer = layui.layer;
    initData()
})

function initData() {
    var list = table.render({
        elem: "#list",
        cols: columns,
        page: true,
        limits: [5, 7, 10],
        limit: 5
    })
    asyncInvoke("/user.api?sort=lastLoginTime", "GetList", null, function(d) {
        data = d.data.rows;
        list.reload({
            data: d.data.rows
        })
    })
    return list;
}


function getRank(row) {
    var rsl = "";
    if (row.power == 1) {
        rsl = "普通会员";
    } else if (row.power == 2) {
        rsl = "黄金会员";
    }
    return rsl
}

function getHandleBtn(d) {
    var save = "<a href='javascript:;' class='small-button delete' style='margin:0 5px' onclick=save('" + d.LAY_TABLE_INDEX + "')>保存</button>";
    var _delete = "<a href='javascript:;' class='small-button delete' style='margin:0 5px' onclick=del('" + d.id + "','" + d.LAY_TABLE_INDEX + "')>删除</button>";
    return save + _delete;
}

function save(index) {
    var item = data[index];
    item.power = item.power == "1" ? 1 : 2;
    var jsonData = JSON.stringify(item);
    asyncInvoke("/user.api", "Update", jsonData, function(d) {
        if (d.code === 0) {
            alert("保存成功");
        } else {
            errHandler(d.data);
        }
    })
}

function del(id, index) {
    layer.confirm('确认删除？', {
        btn: ['是的', '取消'] //按钮
    }, function() {
        layer.load(1, { shade: [0.2, '#000000'] });
        asyncInvoke("/user.api", "Delete", { id: id }, function(d) {
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
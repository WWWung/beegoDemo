// 对Date的扩展，将 Date 转化为指定格式的String
// 月(M)、日(d)、小时(h)、分(m)、秒(s)、季度(q) 可以用 1-2 个占位符， 
// 年(y)可以用 1-4 个占位符，毫秒(S)只能用 1 个占位符(是 1-3 位的数字) 
// 例子： 
// (new Date()).Format("yyyy-MM-dd HH:mm:ss.S") ==> 2006-07-02 08:09:04.423 
// (new Date()).Format("yyyy-M-d H:m:s.S")      ==> 2006-7-2 8:9:4.18 
Date.prototype.format = function(fmt) { //author: meizz 
    var o = {
        "M+": this.getMonth() + 1, //月份 
        "d+": this.getDate(), //日 
        "H+": this.getHours(), //小时 
        "m+": this.getMinutes(), //分 
        "s+": this.getSeconds(), //秒 
        "q+": Math.floor((this.getMonth() + 3) / 3), //季度 
        "S": this.getMilliseconds() //毫秒 
    };
    if (/(y+)/.test(fmt)) fmt = fmt.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));
    for (var k in o)
        if (new RegExp("(" + k + ")").test(fmt)) fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
    return fmt;
}

//  获取input值并且设置到对象data里(string)
function getStringFromInput(id, data) {
    var val = $("#" + id).val();
    if (data) {
        data[id] = val;
    }
    return val;
}

//  获取input值并且设置到对象data里（int）
function getIntFromInput(id, data) {
    var val = getStringFromInput(id);
    var v = parseInt(val);
    if (v !== v) {
        v = 0;
    }
    if (data) {
        data[id] = v;
    }
    return v;
}

function getDateFromInput(id, data) {
    var val = getStringFromInput(id);
    if (!val) {
        val = new Date("1960-01-01");
    } else {
        val = new Date(val);
    }
    if (data) {
        data[id] = val;
    }
    return val;
}

//  ajax(数据用json)
function asyncInvoke(url, method, data, successFunc, failureFunc) {
    var d = {};
    if (data instanceof Object) {
        d = Object.assign({ m: encodeURIComponent(method) }, data);
    } else {
        d = {
            m: encodeURIComponent(method),
            data: data
        };
    }
    if (!failureFunc) {
        failureFunc = function(XMLHttpRequest, textStatus, errorThrown) {
            invokeServiceError(XMLHttpRequest, textStatus, errorThrown);
        }
    }
    $.ajax({
        type: "post",
        url: url,
        success: successFunc,
        error: failureFunc,
        timeout: 20000,
        data: d
    })
}

//获取地址栏参数
function parseQueryString() {
    var url = window.document.location.href.toString();
    var u = url.split("?");
    if (typeof(u[1]) == "string") {
        u = u[1].split("&");
        var get = {};
        for (var i in u) {
            var j = u[i].split("=");
            get[j[0]] = j[1];
        }
        return get;
    } else {
        return {};
    }
}

function getQeuryParam(key) {
    var qs = parseQueryString();
    return qs[key] || "";
}

//  ajax失败
function invokeServiceError(XMLHttpRequest, textStatus, errorThrown) {
    console.log(XMLHttpRequest);
    console.log(textStatus);
    console.log(errorThrown);
    errHandler("error");
}

function timeFormatter(time) {
    if (new Date(time).getTime() < 0) {
        return "";
    }
    return new Date(time).format("yyyy-MM-dd")
}

//  错误处理，待补充...
function errHandler(errMsg, layer) {
    if (layer && layer.msg) {
        layer.msg(errMsg);
    } else {
        alert(errMsg);
    }
}
//  检查是否手机号码格式是否正确(开头为1，第二位数字为3、4、5、8，长度为11)
function checkPhone(phone) {
    return (/^1[3|4|5|8][0-9]\d{4,8}$/.test(phone)) || phone.length === 11
}
//  检查字符串是否全是汉字
function checkChinese(temp) {
    var re = /^[\u4e00-\u9fa5]+$/;
    return re.test(temp);
}
//必须为字母加数字且长度不小于6位
function checkPassWord(password) {
    var str = password;
    if (str == null || str.length < 6) {
        return false;
    }
    var reg1 = new RegExp(/^[0-9A-Za-z]+$/);
    if (!reg1.test(str)) {
        return false;
    }
    var reg = new RegExp(/[A-Za-z].*[0-9]|[0-9].*[A-Za-z]/);
    if (reg.test(str)) {
        return true;
    } else {
        return false;
    }
}
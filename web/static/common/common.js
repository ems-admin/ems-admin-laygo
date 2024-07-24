
//  成功消息
function successMsg(layer, msg) {
    layer.msg(msg, {icon: 1})
}

//  失败消息
function errorMsg(msg){
    layer.msg(msg, {icon: 2})
}

//  警告消息
function warningMsg(layer, msg){
    layer.msg(msg, {icon: 3})
}

//  信息消息
function infoMsg(layer, msg) {
    layer.msg(msg, {icon: 0})
}


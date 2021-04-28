import request from '@/utils/request'

// 查询OmsOrderOperateHistory列表
export function listOmsOrderOperateHistory(query) {
return request({
url: '/admin/v1/omsorderoperatehistoryList',
method: 'get',
params: query
})
}

// 查询OmsOrderOperateHistory详细
export function getOmsOrderOperateHistory (id) {
return request({
url: '/admin/v1/omsorderoperatehistory/' + id,
method: 'get'
})
}


// 新增OmsOrderOperateHistory
export function addOmsOrderOperateHistory(data) {
return request({
url: '/admin/v1/omsorderoperatehistory',
method: 'post',
data: data
})
}

// 修改OmsOrderOperateHistory
export function updateOmsOrderOperateHistory(data) {
return request({
url: '/admin/v1/omsorderoperatehistory',
method: 'put',
data: data
})
}

// 删除OmsOrderOperateHistory
export function delOmsOrderOperateHistory(id) {
return request({
url: '/admin/v1/omsorderoperatehistory/' + id,
method: 'delete'
})
}


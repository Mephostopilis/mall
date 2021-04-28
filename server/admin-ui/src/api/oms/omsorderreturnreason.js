import request from '@/utils/request'

// 查询OmsOrderReturnReason列表
export function listOmsOrderReturnReason(query) {
return request({
url: '/admin/v1/omsorderreturnreasonList',
method: 'get',
params: query
})
}

// 查询OmsOrderReturnReason详细
export function getOmsOrderReturnReason (id) {
return request({
url: '/admin/v1/omsorderreturnreason/' + id,
method: 'get'
})
}


// 新增OmsOrderReturnReason
export function addOmsOrderReturnReason(data) {
return request({
url: '/admin/v1/omsorderreturnreason',
method: 'post',
data: data
})
}

// 修改OmsOrderReturnReason
export function updateOmsOrderReturnReason(data) {
return request({
url: '/admin/v1/omsorderreturnreason',
method: 'put',
data: data
})
}

// 删除OmsOrderReturnReason
export function delOmsOrderReturnReason(id) {
return request({
url: '/admin/v1/omsorderreturnreason/' + id,
method: 'delete'
})
}


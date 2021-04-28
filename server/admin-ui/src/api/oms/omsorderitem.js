import request from '@/utils/request'

// 查询OmsOrderItem列表
export function listOmsOrderItem(query) {
return request({
url: '/admin/v1/omsorderitemList',
method: 'get',
params: query
})
}

// 查询OmsOrderItem详细
export function getOmsOrderItem (id) {
return request({
url: '/admin/v1/omsorderitem/' + id,
method: 'get'
})
}


// 新增OmsOrderItem
export function addOmsOrderItem(data) {
return request({
url: '/admin/v1/omsorderitem',
method: 'post',
data: data
})
}

// 修改OmsOrderItem
export function updateOmsOrderItem(data) {
return request({
url: '/admin/v1/omsorderitem',
method: 'put',
data: data
})
}

// 删除OmsOrderItem
export function delOmsOrderItem(id) {
return request({
url: '/admin/v1/omsorderitem/' + id,
method: 'delete'
})
}


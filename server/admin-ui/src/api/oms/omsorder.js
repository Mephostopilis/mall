import request from '@/utils/request'

// 查询OmsOrder列表
export function listOmsOrder(query) {
return request({
url: '/admin/v1/omsorderList',
method: 'get',
params: query
})
}

// 查询OmsOrder详细
export function getOmsOrder (id) {
return request({
url: '/admin/v1/omsorder/' + id,
method: 'get'
})
}


// 新增OmsOrder
export function addOmsOrder(data) {
return request({
url: '/admin/v1/omsorder',
method: 'post',
data: data
})
}

// 修改OmsOrder
export function updateOmsOrder(data) {
return request({
url: '/admin/v1/omsorder',
method: 'put',
data: data
})
}

// 删除OmsOrder
export function delOmsOrder(id) {
return request({
url: '/admin/v1/omsorder/' + id,
method: 'delete'
})
}


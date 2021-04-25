import request from '@/utils/request'

// 查询SmsCoupon列表
export function listSmsCoupon(query) {
return request({
url: '/admin/v1/smscouponList',
method: 'get',
params: query
})
}

// 查询SmsCoupon详细
export function getSmsCoupon (id) {
return request({
url: '/admin/v1/smscoupon/' + id,
method: 'get'
})
}


// 新增SmsCoupon
export function addSmsCoupon(data) {
return request({
url: '/admin/v1/smscoupon',
method: 'post',
data: data
})
}

// 修改SmsCoupon
export function updateSmsCoupon(data) {
return request({
url: '/admin/v1/smscoupon',
method: 'put',
data: data
})
}

// 删除SmsCoupon
export function delSmsCoupon(id) {
return request({
url: '/admin/v1/smscoupon/' + id,
method: 'delete'
})
}


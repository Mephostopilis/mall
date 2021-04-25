import request from '@/utils/request'

// 查询SmsCouponProductRelation列表
export function listSmsCouponProductRelation(query) {
return request({
url: '/admin/v1/smscouponproductrelationList',
method: 'get',
params: query
})
}

// 查询SmsCouponProductRelation详细
export function getSmsCouponProductRelation (id) {
return request({
url: '/admin/v1/smscouponproductrelation/' + id,
method: 'get'
})
}


// 新增SmsCouponProductRelation
export function addSmsCouponProductRelation(data) {
return request({
url: '/admin/v1/smscouponproductrelation',
method: 'post',
data: data
})
}

// 修改SmsCouponProductRelation
export function updateSmsCouponProductRelation(data) {
return request({
url: '/admin/v1/smscouponproductrelation',
method: 'put',
data: data
})
}

// 删除SmsCouponProductRelation
export function delSmsCouponProductRelation(id) {
return request({
url: '/admin/v1/smscouponproductrelation/' + id,
method: 'delete'
})
}


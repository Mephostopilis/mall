import request from '@/utils/request'

// 查询SmsCouponProductCategoryRelation列表
export function listSmsCouponProductCategoryRelation(query) {
return request({
url: '/admin/v1/smscouponproductcategoryrelationList',
method: 'get',
params: query
})
}

// 查询SmsCouponProductCategoryRelation详细
export function getSmsCouponProductCategoryRelation (id) {
return request({
url: '/admin/v1/smscouponproductcategoryrelation/' + id,
method: 'get'
})
}


// 新增SmsCouponProductCategoryRelation
export function addSmsCouponProductCategoryRelation(data) {
return request({
url: '/admin/v1/smscouponproductcategoryrelation',
method: 'post',
data: data
})
}

// 修改SmsCouponProductCategoryRelation
export function updateSmsCouponProductCategoryRelation(data) {
return request({
url: '/admin/v1/smscouponproductcategoryrelation',
method: 'put',
data: data
})
}

// 删除SmsCouponProductCategoryRelation
export function delSmsCouponProductCategoryRelation(id) {
return request({
url: '/admin/v1/smscouponproductcategoryrelation/' + id,
method: 'delete'
})
}


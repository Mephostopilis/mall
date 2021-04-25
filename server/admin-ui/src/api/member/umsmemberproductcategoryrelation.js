import request from '@/utils/request'

// 查询UmsMemberProductCategoryRelation列表
export function listUmsMemberProductCategoryRelation(query) {
return request({
url: '/admin/v1/umsmemberproductcategoryrelationList',
method: 'get',
params: query
})
}

// 查询UmsMemberProductCategoryRelation详细
export function getUmsMemberProductCategoryRelation (id) {
return request({
url: '/admin/v1/umsmemberproductcategoryrelation/' + id,
method: 'get'
})
}


// 新增UmsMemberProductCategoryRelation
export function addUmsMemberProductCategoryRelation(data) {
return request({
url: '/admin/v1/umsmemberproductcategoryrelation',
method: 'post',
data: data
})
}

// 修改UmsMemberProductCategoryRelation
export function updateUmsMemberProductCategoryRelation(data) {
return request({
url: '/admin/v1/umsmemberproductcategoryrelation',
method: 'put',
data: data
})
}

// 删除UmsMemberProductCategoryRelation
export function delUmsMemberProductCategoryRelation(id) {
return request({
url: '/admin/v1/umsmemberproductcategoryrelation/' + id,
method: 'delete'
})
}


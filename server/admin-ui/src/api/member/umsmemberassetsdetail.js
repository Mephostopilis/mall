import request from '@/utils/request'

// 查询UmsMemberAssetsDetail列表
export function listUmsMemberAssetsDetail(query) {
return request({
url: '/admin/v1/umsmemberassetsdetailList',
method: 'get',
params: query
})
}

// 查询UmsMemberAssetsDetail详细
export function getUmsMemberAssetsDetail (id) {
return request({
url: '/admin/v1/umsmemberassetsdetail/' + id,
method: 'get'
})
}


// 新增UmsMemberAssetsDetail
export function addUmsMemberAssetsDetail(data) {
return request({
url: '/admin/v1/umsmemberassetsdetail',
method: 'post',
data: data
})
}

// 修改UmsMemberAssetsDetail
export function updateUmsMemberAssetsDetail(data) {
return request({
url: '/admin/v1/umsmemberassetsdetail',
method: 'put',
data: data
})
}

// 删除UmsMemberAssetsDetail
export function delUmsMemberAssetsDetail(id) {
return request({
url: '/admin/v1/umsmemberassetsdetail/' + id,
method: 'delete'
})
}


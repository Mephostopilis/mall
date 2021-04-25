import request from '@/utils/request'

// 查询UmsMemberTag列表
export function listUmsMemberTag(query) {
return request({
url: '/admin/v1/umsmembertagList',
method: 'get',
params: query
})
}

// 查询UmsMemberTag详细
export function getUmsMemberTag (id) {
return request({
url: '/admin/v1/umsmembertag/' + id,
method: 'get'
})
}


// 新增UmsMemberTag
export function addUmsMemberTag(data) {
return request({
url: '/admin/v1/umsmembertag',
method: 'post',
data: data
})
}

// 修改UmsMemberTag
export function updateUmsMemberTag(data) {
return request({
url: '/admin/v1/umsmembertag',
method: 'put',
data: data
})
}

// 删除UmsMemberTag
export function delUmsMemberTag(id) {
return request({
url: '/admin/v1/umsmembertag/' + id,
method: 'delete'
})
}


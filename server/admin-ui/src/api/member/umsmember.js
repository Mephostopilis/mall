import request from '@/utils/request'

// 查询UmsMember列表
export function listUmsMember(query) {
return request({
url: '/admin/v1/umsmemberList',
method: 'get',
params: query
})
}

// 查询UmsMember详细
export function getUmsMember (id) {
return request({
url: '/admin/v1/umsmember/' + id,
method: 'get'
})
}


// 新增UmsMember
export function addUmsMember(data) {
return request({
url: '/admin/v1/umsmember',
method: 'post',
data: data
})
}

// 修改UmsMember
export function updateUmsMember(data) {
return request({
url: '/admin/v1/umsmember',
method: 'put',
data: data
})
}

// 删除UmsMember
export function delUmsMember(id) {
return request({
url: '/admin/v1/umsmember/' + id,
method: 'delete'
})
}


import request from '@/utils/request'

// 查询UmsMemberLevel列表
export function listUmsMemberLevel(query) {
return request({
url: '/admin/v1/umsmemberlevelList',
method: 'get',
params: query
})
}

// 查询UmsMemberLevel详细
export function getUmsMemberLevel (id) {
return request({
url: '/admin/v1/umsmemberlevel/' + id,
method: 'get'
})
}


// 新增UmsMemberLevel
export function addUmsMemberLevel(data) {
return request({
url: '/admin/v1/umsmemberlevel',
method: 'post',
data: data
})
}

// 修改UmsMemberLevel
export function updateUmsMemberLevel(data) {
return request({
url: '/admin/v1/umsmemberlevel',
method: 'put',
data: data
})
}

// 删除UmsMemberLevel
export function delUmsMemberLevel(id) {
return request({
url: '/admin/v1/umsmemberlevel/' + id,
method: 'delete'
})
}


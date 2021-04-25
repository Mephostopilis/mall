import request from '@/utils/request'

// 查询UmsMemberTask列表
export function listUmsMemberTask(query) {
return request({
url: '/admin/v1/umsmembertaskList',
method: 'get',
params: query
})
}

// 查询UmsMemberTask详细
export function getUmsMemberTask (id) {
return request({
url: '/admin/v1/umsmembertask/' + id,
method: 'get'
})
}


// 新增UmsMemberTask
export function addUmsMemberTask(data) {
return request({
url: '/admin/v1/umsmembertask',
method: 'post',
data: data
})
}

// 修改UmsMemberTask
export function updateUmsMemberTask(data) {
return request({
url: '/admin/v1/umsmembertask',
method: 'put',
data: data
})
}

// 删除UmsMemberTask
export function delUmsMemberTask(id) {
return request({
url: '/admin/v1/umsmembertask/' + id,
method: 'delete'
})
}


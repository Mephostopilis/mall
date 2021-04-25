import request from '@/utils/request'

// 查询UmsMemberMemberTagRelation列表
export function listUmsMemberMemberTagRelation(query) {
return request({
url: '/admin/v1/umsmembermembertagrelationList',
method: 'get',
params: query
})
}

// 查询UmsMemberMemberTagRelation详细
export function getUmsMemberMemberTagRelation (id) {
return request({
url: '/admin/v1/umsmembermembertagrelation/' + id,
method: 'get'
})
}


// 新增UmsMemberMemberTagRelation
export function addUmsMemberMemberTagRelation(data) {
return request({
url: '/admin/v1/umsmembermembertagrelation',
method: 'post',
data: data
})
}

// 修改UmsMemberMemberTagRelation
export function updateUmsMemberMemberTagRelation(data) {
return request({
url: '/admin/v1/umsmembermembertagrelation',
method: 'put',
data: data
})
}

// 删除UmsMemberMemberTagRelation
export function delUmsMemberMemberTagRelation(id) {
return request({
url: '/admin/v1/umsmembermembertagrelation/' + id,
method: 'delete'
})
}


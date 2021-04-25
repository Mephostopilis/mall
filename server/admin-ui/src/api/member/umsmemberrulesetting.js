import request from '@/utils/request'

// 查询UmsMemberRuleSetting列表
export function listUmsMemberRuleSetting(query) {
return request({
url: '/admin/v1/umsmemberrulesettingList',
method: 'get',
params: query
})
}

// 查询UmsMemberRuleSetting详细
export function getUmsMemberRuleSetting (id) {
return request({
url: '/admin/v1/umsmemberrulesetting/' + id,
method: 'get'
})
}


// 新增UmsMemberRuleSetting
export function addUmsMemberRuleSetting(data) {
return request({
url: '/admin/v1/umsmemberrulesetting',
method: 'post',
data: data
})
}

// 修改UmsMemberRuleSetting
export function updateUmsMemberRuleSetting(data) {
return request({
url: '/admin/v1/umsmemberrulesetting',
method: 'put',
data: data
})
}

// 删除UmsMemberRuleSetting
export function delUmsMemberRuleSetting(id) {
return request({
url: '/admin/v1/umsmemberrulesetting/' + id,
method: 'delete'
})
}


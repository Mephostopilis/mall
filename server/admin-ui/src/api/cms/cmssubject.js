import request from '@/utils/request'

// 查询CmsSubject列表
export function listCmsSubject(query) {
return request({
url: '/admin/v1/cmssubjectList',
method: 'get',
params: query
})
}

// 查询CmsSubject详细
export function getCmsSubject (id) {
return request({
url: '/admin/v1/cmssubject/' + id,
method: 'get'
})
}


// 新增CmsSubject
export function addCmsSubject(data) {
return request({
url: '/admin/v1/cmssubject',
method: 'post',
data: data
})
}

// 修改CmsSubject
export function updateCmsSubject(data) {
return request({
url: '/admin/v1/cmssubject',
method: 'put',
data: data
})
}

// 删除CmsSubject
export function delCmsSubject(id) {
return request({
url: '/admin/v1/cmssubject/' + id,
method: 'delete'
})
}


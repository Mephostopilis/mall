import request from '@/utils/request'

// 查询CmsSubjectComment列表
export function listCmsSubjectComment(query) {
return request({
url: '/admin/v1/cmssubjectcommentList',
method: 'get',
params: query
})
}

// 查询CmsSubjectComment详细
export function getCmsSubjectComment (id) {
return request({
url: '/admin/v1/cmssubjectcomment/' + id,
method: 'get'
})
}


// 新增CmsSubjectComment
export function addCmsSubjectComment(data) {
return request({
url: '/admin/v1/cmssubjectcomment',
method: 'post',
data: data
})
}

// 修改CmsSubjectComment
export function updateCmsSubjectComment(data) {
return request({
url: '/admin/v1/cmssubjectcomment',
method: 'put',
data: data
})
}

// 删除CmsSubjectComment
export function delCmsSubjectComment(id) {
return request({
url: '/admin/v1/cmssubjectcomment/' + id,
method: 'delete'
})
}


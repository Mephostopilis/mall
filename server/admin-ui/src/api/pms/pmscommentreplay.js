import request from '@/utils/request'

// 查询PmsCommentReplay列表
export function listPmsCommentReplay(query) {
return request({
url: '/admin/v1/pmscommentreplayList',
method: 'get',
params: query
})
}

// 查询PmsCommentReplay详细
export function getPmsCommentReplay (id) {
return request({
url: '/admin/v1/pmscommentreplay/' + id,
method: 'get'
})
}


// 新增PmsCommentReplay
export function addPmsCommentReplay(data) {
return request({
url: '/admin/v1/pmscommentreplay',
method: 'post',
data: data
})
}

// 修改PmsCommentReplay
export function updatePmsCommentReplay(data) {
return request({
url: '/admin/v1/pmscommentreplay',
method: 'put',
data: data
})
}

// 删除PmsCommentReplay
export function delPmsCommentReplay(id) {
return request({
url: '/admin/v1/pmscommentreplay/' + id,
method: 'delete'
})
}


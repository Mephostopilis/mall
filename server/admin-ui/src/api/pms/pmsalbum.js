import request from '@/utils/request'

// 查询PmsAlbum列表
export function listPmsAlbum(query) {
return request({
url: '/admin/v1/pmsalbumList',
method: 'get',
params: query
})
}

// 查询PmsAlbum详细
export function getPmsAlbum (id) {
return request({
url: '/admin/v1/pmsalbum/' + id,
method: 'get'
})
}


// 新增PmsAlbum
export function addPmsAlbum(data) {
return request({
url: '/admin/v1/pmsalbum',
method: 'post',
data: data
})
}

// 修改PmsAlbum
export function updatePmsAlbum(data) {
return request({
url: '/admin/v1/pmsalbum',
method: 'put',
data: data
})
}

// 删除PmsAlbum
export function delPmsAlbum(id) {
return request({
url: '/admin/v1/pmsalbum/' + id,
method: 'delete'
})
}


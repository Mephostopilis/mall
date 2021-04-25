import request from '@/utils/request'

// 查询PmsProductLadder列表
export function listPmsProductLadder(query) {
return request({
url: '/admin/v1/pmsproductladderList',
method: 'get',
params: query
})
}

// 查询PmsProductLadder详细
export function getPmsProductLadder (id) {
return request({
url: '/admin/v1/pmsproductladder/' + id,
method: 'get'
})
}


// 新增PmsProductLadder
export function addPmsProductLadder(data) {
return request({
url: '/admin/v1/pmsproductladder',
method: 'post',
data: data
})
}

// 修改PmsProductLadder
export function updatePmsProductLadder(data) {
return request({
url: '/admin/v1/pmsproductladder',
method: 'put',
data: data
})
}

// 删除PmsProductLadder
export function delPmsProductLadder(id) {
return request({
url: '/admin/v1/pmsproductladder/' + id,
method: 'delete'
})
}


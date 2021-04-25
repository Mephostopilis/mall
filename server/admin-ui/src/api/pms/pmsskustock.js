import request from '@/utils/request'

// 查询PmsSkuStock列表
export function listPmsSkuStock(query) {
return request({
url: '/admin/v1/pmsskustockList',
method: 'get',
params: query
})
}

// 查询PmsSkuStock详细
export function getPmsSkuStock (id) {
return request({
url: '/admin/v1/pmsskustock/' + id,
method: 'get'
})
}


// 新增PmsSkuStock
export function addPmsSkuStock(data) {
return request({
url: '/admin/v1/pmsskustock',
method: 'post',
data: data
})
}

// 修改PmsSkuStock
export function updatePmsSkuStock(data) {
return request({
url: '/admin/v1/pmsskustock',
method: 'put',
data: data
})
}

// 删除PmsSkuStock
export function delPmsSkuStock(id) {
return request({
url: '/admin/v1/pmsskustock/' + id,
method: 'delete'
})
}


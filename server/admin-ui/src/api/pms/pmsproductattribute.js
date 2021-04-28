import request from '@/utils/request'

// 查询PmsProductAttribute列表
export function listPmsProductAttribute(query) {
return request({
url: '/admin/v1/pmsproductattributeList',
method: 'get',
params: query
})
}

// 查询PmsProductAttribute详细
export function getPmsProductAttribute (id) {
return request({
url: '/admin/v1/pmsproductattribute/' + id,
method: 'get'
})
}


// 新增PmsProductAttribute
export function addPmsProductAttribute(data) {
return request({
headers: {
    'Content-Type': 'application/x-www-form-urlencoded'
},
url: '/admin/v1/pmsproductattribute',
method: 'post',
data: data
})
}

// 修改PmsProductAttribute
export function updatePmsProductAttribute(data) {
return request({
headers: {
    'Content-Type': 'application/x-www-form-urlencoded'
},
url: '/admin/v1/pmsproductattribute',
method: 'put',
data: data
})
}

// 删除PmsProductAttribute
export function delPmsProductAttribute(id) {
return request({
url: '/admin/v1/pmsproductattribute/' + id,
method: 'delete'
})
}


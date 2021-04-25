import request from '@/utils/request'

// 查询OmsCompanyAddress列表
export function listOmsCompanyAddress(query) {
return request({
url: '/admin/v1/omscompanyaddressList',
method: 'get',
params: query
})
}

// 查询OmsCompanyAddress详细
export function getOmsCompanyAddress (id) {
return request({
url: '/admin/v1/omscompanyaddress/' + id,
method: 'get'
})
}


// 新增OmsCompanyAddress
export function addOmsCompanyAddress(data) {
return request({
url: '/admin/v1/omscompanyaddress',
method: 'post',
data: data
})
}

// 修改OmsCompanyAddress
export function updateOmsCompanyAddress(data) {
return request({
url: '/admin/v1/omscompanyaddress',
method: 'put',
data: data
})
}

// 删除OmsCompanyAddress
export function delOmsCompanyAddress(id) {
return request({
url: '/admin/v1/omscompanyaddress/' + id,
method: 'delete'
})
}


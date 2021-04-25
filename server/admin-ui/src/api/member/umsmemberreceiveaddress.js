import request from '@/utils/request'

// 查询UmsMemberReceiveAddress列表
export function listUmsMemberReceiveAddress(query) {
return request({
url: '/admin/v1/umsmemberreceiveaddressList',
method: 'get',
params: query
})
}

// 查询UmsMemberReceiveAddress详细
export function getUmsMemberReceiveAddress (id) {
return request({
url: '/admin/v1/umsmemberreceiveaddress/' + id,
method: 'get'
})
}


// 新增UmsMemberReceiveAddress
export function addUmsMemberReceiveAddress(data) {
return request({
url: '/admin/v1/umsmemberreceiveaddress',
method: 'post',
data: data
})
}

// 修改UmsMemberReceiveAddress
export function updateUmsMemberReceiveAddress(data) {
return request({
url: '/admin/v1/umsmemberreceiveaddress',
method: 'put',
data: data
})
}

// 删除UmsMemberReceiveAddress
export function delUmsMemberReceiveAddress(id) {
return request({
url: '/admin/v1/umsmemberreceiveaddress/' + id,
method: 'delete'
})
}


import request from '@/utils/request'

// 查询UmsMemberAssets列表
export function listUmsMemberAssets(query) {
return request({
url: '/admin/v1/umsmemberassetsList',
method: 'get',
params: query
})
}

// 查询UmsMemberAssets详细
export function getUmsMemberAssets (username) {
return request({
url: '/admin/v1/umsmemberassets/' + username,
method: 'get'
})
}


// 新增UmsMemberAssets
export function addUmsMemberAssets(data) {
return request({
url: '/admin/v1/umsmemberassets',
method: 'post',
data: data
})
}

// 修改UmsMemberAssets
export function updateUmsMemberAssets(data) {
return request({
url: '/admin/v1/umsmemberassets',
method: 'put',
data: data
})
}

// 删除UmsMemberAssets
export function delUmsMemberAssets(username) {
return request({
url: '/admin/v1/umsmemberassets/' + username,
method: 'delete'
})
}


import request from '@/utils/request'

// 查询SmsCouponHistory列表
export function listSmsCouponHistory(query) {
return request({
url: '/admin/v1/smscouponhistoryList',
method: 'get',
params: query
})
}

// 查询SmsCouponHistory详细
export function getSmsCouponHistory (id) {
return request({
url: '/admin/v1/smscouponhistory/' + id,
method: 'get'
})
}


// 新增SmsCouponHistory
export function addSmsCouponHistory(data) {
return request({
url: '/admin/v1/smscouponhistory',
method: 'post',
data: data
})
}

// 修改SmsCouponHistory
export function updateSmsCouponHistory(data) {
return request({
url: '/admin/v1/smscouponhistory',
method: 'put',
data: data
})
}

// 删除SmsCouponHistory
export function delSmsCouponHistory(id) {
return request({
url: '/admin/v1/smscouponhistory/' + id,
method: 'delete'
})
}


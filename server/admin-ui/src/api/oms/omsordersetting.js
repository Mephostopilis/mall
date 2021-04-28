import request from '@/utils/request'

// 查询OmsOrderSetting列表
export function listOmsOrderSetting(query) {
return request({
url: '/admin/v1/omsordersettingList',
method: 'get',
params: query
})
}

// 查询OmsOrderSetting详细
export function getOmsOrderSetting (id) {
return request({
url: '/admin/v1/omsordersetting/' + id,
method: 'get'
})
}


// 新增OmsOrderSetting
export function addOmsOrderSetting(data) {
return request({
url: '/admin/v1/omsordersetting',
method: 'post',
data: data
})
}

// 修改OmsOrderSetting
export function updateOmsOrderSetting(data) {
return request({
url: '/admin/v1/omsordersetting',
method: 'put',
data: data
})
}

// 删除OmsOrderSetting
export function delOmsOrderSetting(id) {
return request({
url: '/admin/v1/omsordersetting/' + id,
method: 'delete'
})
}


import request from '@/utils/request'

// 查询UmsMemberStatisticsInfo列表
export function listUmsMemberStatisticsInfo(query) {
return request({
url: '/admin/v1/umsmemberstatisticsinfoList',
method: 'get',
params: query
})
}

// 查询UmsMemberStatisticsInfo详细
export function getUmsMemberStatisticsInfo (id) {
return request({
url: '/admin/v1/umsmemberstatisticsinfo/' + id,
method: 'get'
})
}


// 新增UmsMemberStatisticsInfo
export function addUmsMemberStatisticsInfo(data) {
return request({
url: '/admin/v1/umsmemberstatisticsinfo',
method: 'post',
data: data
})
}

// 修改UmsMemberStatisticsInfo
export function updateUmsMemberStatisticsInfo(data) {
return request({
url: '/admin/v1/umsmemberstatisticsinfo',
method: 'put',
data: data
})
}

// 删除UmsMemberStatisticsInfo
export function delUmsMemberStatisticsInfo(id) {
return request({
url: '/admin/v1/umsmemberstatisticsinfo/' + id,
method: 'delete'
})
}


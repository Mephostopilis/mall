import request from '@/utils/request'

// 查询CmsHelp列表
export function listCmsHelp(query) {
return request({
url: '/admin/v1/cmshelpList',
method: 'get',
params: query
})
}

// 查询CmsHelp详细
export function getCmsHelp (id) {
return request({
url: '/admin/v1/cmshelp/' + id,
method: 'get'
})
}


// 新增CmsHelp
export function addCmsHelp(data) {
return request({
url: '/admin/v1/cmshelp',
method: 'post',
data: data
})
}

// 修改CmsHelp
export function updateCmsHelp(data) {
return request({
url: '/admin/v1/cmshelp',
method: 'put',
data: data
})
}

// 删除CmsHelp
export function delCmsHelp(id) {
return request({
url: '/admin/v1/cmshelp/' + id,
method: 'delete'
})
}


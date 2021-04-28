import request from '@/utils/request'

// 查询CmsHelpCategory列表
export function listCmsHelpCategory(query) {
  return request({
    url: '/admin/v1/cmshelpcategoryList',
    method: 'get',
    params: query
  })
}

// 查询CmsHelpCategory详细
export function getCmsHelpCategory(id) {
  return request({
    url: '/admin/v1/cmshelpcategory/' + id,
    method: 'get'
  })
}


// 新增CmsHelpCategory
export function addCmsHelpCategory(data) {
  return request({
    url: '/admin/v1/cmshelpcategory',
    method: 'post',
    data: data
  })
}

// 修改CmsHelpCategory
export function updateCmsHelpCategory(data) {
  return request({
    url: '/admin/v1/cmshelpcategory',
    method: 'put',
    data: data
  })
}

// 删除CmsHelpCategory
export function delCmsHelpCategory(id) {
  return request({
    url: '/admin/v1/cmshelpcategory/' + id,
    method: 'delete'
  })
}

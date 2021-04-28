import request from '@/utils/request'

// 查询CmsSubjectCategory列表
export function listCmsSubjectCategory(query) {
  return request({
    url: '/admin/v1/cmssubjectcategoryList',
    method: 'get',
    params: query
  })
}

// 查询CmsSubjectCategory详细
export function getCmsSubjectCategory(id) {
  return request({
    url: '/admin/v1/cmssubjectcategory/' + id,
    method: 'get'
  })
}


// 新增CmsSubjectCategory
export function addCmsSubjectCategory(data) {
  return request({
    url: '/admin/v1/cmssubjectcategory',
    method: 'post',
    data: data
  })
}

// 修改CmsSubjectCategory
export function updateCmsSubjectCategory(data) {
  return request({
    url: '/admin/v1/cmssubjectcategory',
    method: 'put',
    data: data
  })
}

// 删除CmsSubjectCategory
export function delCmsSubjectCategory(id) {
  return request({
    url: '/admin/v1/cmssubjectcategory/' + id,
    method: 'delete'
  })
}


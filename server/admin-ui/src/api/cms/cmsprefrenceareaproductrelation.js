import request from '@/utils/request'

// 查询CmsPrefrenceAreaProductRelation列表
export function listCmsPrefrenceAreaProductRelation(query) {
  return request({
    url: '/api/v1/cmsprefrenceareaproductrelationList',
    method: 'get',
    params: query
  })
}

// 查询CmsPrefrenceAreaProductRelation详细
export function getCmsPrefrenceAreaProductRelation(id) {
  return request({
    url: '/api/v1/cmsprefrenceareaproductrelation/' + id,
    method: 'get'
  })
}


// 新增CmsPrefrenceAreaProductRelation
export function addCmsPrefrenceAreaProductRelation(data) {
  return request({
    url: '/api/v1/cmsprefrenceareaproductrelation',
    method: 'post',
    data: data
  })
}

// 修改CmsPrefrenceAreaProductRelation
export function updateCmsPrefrenceAreaProductRelation(data) {
  return request({
    url: '/api/v1/cmsprefrenceareaproductrelation',
    method: 'put',
    data: data
  })
}

// 删除CmsPrefrenceAreaProductRelation
export function delCmsPrefrenceAreaProductRelation(id) {
  return request({
    url: '/api/v1/cmsprefrenceareaproductrelation/' + id,
    method: 'delete'
  })
}


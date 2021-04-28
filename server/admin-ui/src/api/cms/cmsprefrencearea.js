import request from '@/utils/request'

// 查询CmsPrefrenceArea列表
export function listCmsPrefrenceArea(query) {
  return request({
    url: '/api/v1/cmsprefrenceareaList',
    method: 'get',
    params: query
  })
}

// 查询CmsPrefrenceArea详细
export function getCmsPrefrenceArea(id) {
  return request({
    url: '/api/v1/cmsprefrencearea/' + id,
    method: 'get'
  })
}


// 新增CmsPrefrenceArea
export function addCmsPrefrenceArea(data) {
  return request({
    url: '/api/v1/cmsprefrencearea',
    method: 'post',
    data: data
  })
}

// 修改CmsPrefrenceArea
export function updateCmsPrefrenceArea(data) {
  return request({
    url: '/api/v1/cmsprefrencearea',
    method: 'put',
    data: data
  })
}

// 删除CmsPrefrenceArea
export function delCmsPrefrenceArea(id) {
  return request({
    url: '/api/v1/cmsprefrencearea/' + id,
    method: 'delete'
  })
}


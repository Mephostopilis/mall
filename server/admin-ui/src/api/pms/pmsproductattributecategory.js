import request from '@/utils/request'
import qs from 'qs'

// 查询PmsProductAttributeCategory列表
export function listPmsProductAttributeCategory(query) {
  return request({
    url: '/admin/v1/pmsproductattributecategoryList',
    method: 'get',
    params: query
  })
}

// 查询PmsProductAttributeCategory详细
export function getPmsProductAttributeCategory(id) {
  return request({
    url: '/admin/v1/pmsproductattributecategory/' + id,
    method: 'get'
  })
}


// 新增PmsProductAttributeCategory
export function addPmsProductAttributeCategory(data) {
  return request({
    url: '/admin/v1/pmsproductattributecategory',
    method: 'post',
    data: data
  })
}

// 修改PmsProductAttributeCategory
export function updatePmsProductAttributeCategory(data) {
  return request({
    url: '/admin/v1/pmsproductattributecategory',
    method: 'put',
    data: data
  })
}

// 删除PmsProductAttributeCategory
export function delPmsProductAttributeCategory(id) {
  return request({
    url: '/admin/v1/pmsproductattributecategory/' + id,
    method: 'delete'
  })
}


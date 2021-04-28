import request from '@/utils/request'
import qs from 'qs'

// 查询PmsProductCategoryAttributeRelation列表
export function listPmsProductCategoryAttributeRelation(query) {
  return request({
    url: '/admin/v1/pmsproductcategoryattributerelationList',
    method: 'get',
    params: query
  })
}

// 查询PmsProductCategoryAttributeRelation详细
export function getPmsProductCategoryAttributeRelation(id) {
  return request({
    url: '/admin/v1/pmsproductcategoryattributerelation/' + id,
    method: 'get'
  })
}

// 新增PmsProductCategoryAttributeRelation
export function addPmsProductCategoryAttributeRelation(data) {
  return request({
    url: '/admin/v1/pmsproductcategoryattributerelation',
    method: 'post',
    data: data
  })
}

// 修改PmsProductCategoryAttributeRelation
export function updatePmsProductCategoryAttributeRelation(data) {
  return request({
    url: '/admin/v1/pmsproductcategoryattributerelation',
    method: 'put',
    data: data
  })
}

// 删除PmsProductCategoryAttributeRelation
export function delPmsProductCategoryAttributeRelation(id) {
  return request({
    url: '/admin/v1/pmsproductcategoryattributerelation/' + id,
    method: 'delete'
  })
}

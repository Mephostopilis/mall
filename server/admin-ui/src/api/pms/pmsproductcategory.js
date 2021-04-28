import request from '@/utils/request'
import qs from 'qs'

// 查询PmsProductCategory列表
export function listPmsProductCategory(query) {
  return request({
    url: '/admin/v1/pmsproductcategoryList',
    method: 'get',
    params: query
  })
}

// 查询PmsProductCategory列表
export function listPmsProductCategoryTree(query) {
  return request({
    url: '/admin/v1/pmsproductcategorytreeList',
    method: 'get',
    params: query
  })
}

// 查询PmsProductCategory详细
export function getPmsProductCategory(id) {
  return request({
    url: '/admin/v1/pmsproductcategory/' + id,
    method: 'get'
  })
}

// 新增PmsProductCategory
export function addPmsProductCategory(data) {
  return request({
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded'
    },
    url: '/admin/v1/pmsproductcategory',
    method: 'post',
    data: data
  })
}

// 修改PmsProductCategory
export function updatePmsProductCategory(data) {
  return request({
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded'
    },
    url: '/admin/v1/pmsproductcategory',
    method: 'put',
    data: data
  })
}

// 删除PmsProductCategory
export function delPmsProductCategory(id) {
  return request({
    url: '/admin/v1/pmsproductcategory/' + id,
    method: 'delete'
  })
}


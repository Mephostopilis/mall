import request from '@/utils/request'
import qs from 'qs'

// 查询PmsProduct列表
export function listPmsProduct(query) {
  return request({
    url: '/admin/v1/pmsproductList',
    method: 'get',
    params: query
  })
}

// 查询PmsProduct详细
export function getPmsProduct(id) {
  return request({
    url: '/admin/v1/pmsproduct/' + id,
    method: 'get'
  })
}


// 新增PmsProduct
export function addPmsProduct(data) {
  return request({
    url: '/admin/v1/pmsproduct',
    method: 'post',
    data: data
  })
}

// 修改PmsProduct
export function updatePmsProduct(data) {
  return request({
    url: '/admin/v1/pmsproduct',
    method: 'put',
    data: data
  })
}

// 删除PmsProduct
export function delPmsProduct(id) {
  return request({
    url: '/admin/v1/pmsproduct/' + id,
    method: 'delete'
  })
}


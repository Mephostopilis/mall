import request from '@/utils/request'
import qs from 'qs'

// 查询PmsProductFullReduction列表
export function listPmsProductFullReduction(query) {
  return request({
    url: '/admin/v1/pmsproductfullreductionList',
    method: 'get',
    params: query
  })
}

// 查询PmsProductFullReduction详细
export function getPmsProductFullReduction(id) {
  return request({
    url: '/admin/v1/pmsproductfullreduction/' + id,
    method: 'get'
  })
}

// 新增PmsProductFullReduction
export function addPmsProductFullReduction(data) {
  return request({
    url: '/admin/v1/pmsproductfullreduction',
    method: 'post',
    data: data
  })
}

// 修改PmsProductFullReduction
export function updatePmsProductFullReduction(data) {
  return request({
    url: '/admin/v1/pmsproductfullreduction',
    method: 'put',
    data: data
  })
}

// 删除PmsProductFullReduction
export function delPmsProductFullReduction(id) {
  return request({
    url: '/admin/v1/pmsproductfullreduction/' + id,
    method: 'delete'
  })
}


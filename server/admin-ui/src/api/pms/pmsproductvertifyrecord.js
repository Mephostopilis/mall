import request from '@/utils/request'
import qs from 'qs'

// 查询PmsProductVertifyRecord列表
export function listPmsProductVertifyRecord(query) {
  return request({
    url: '/admin/v1/pmsproductvertifyrecordList',
    method: 'get',
    params: query
  })
}

// 查询PmsProductVertifyRecord详细
export function getPmsProductVertifyRecord(id) {
  return request({
    url: '/admin/v1/pmsproductvertifyrecord/' + id,
    method: 'get'
  })
}


// 新增PmsProductVertifyRecord
export function addPmsProductVertifyRecord(data) {
  return request({
    url: '/admin/v1/pmsproductvertifyrecord',
    method: 'post',
    data: data
  })
}

// 修改PmsProductVertifyRecord
export function updatePmsProductVertifyRecord(data) {
  return request({
    url: '/admin/v1/pmsproductvertifyrecord',
    method: 'put',
    data: data
  })
}

// 删除PmsProductVertifyRecord
export function delPmsProductVertifyRecord(id) {
  return request({
    url: '/admin/v1/pmsproductvertifyrecord/' + id,
    method: 'delete'
  })
}


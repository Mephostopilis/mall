import request from '@/utils/request'
import qs from 'qs'

// 查询PmsBrand列表
export function listPmsBrand(query) {
  return request({
    url: '/admin/v1/pmsbrandList',
    method: 'get',
    params: query
  })
}

// 查询PmsBrand详细
export function getPmsBrand(id) {
  return request({
    url: '/admin/v1/pmsbrand/' + id,
    method: 'get'
  })
}


// 新增PmsBrand
export function addPmsBrand(data) {
  return request({
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded'
    },
    url: '/admin/v1/pmsbrand',
    method: 'post',
    data: data
  })
}

// 修改PmsBrand
export function updatePmsBrand(data) {
  return request({
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded'
    },
    url: '/admin/v1/pmsbrand',
    method: 'put',
    data: data
  })
}

// 删除PmsBrand
export function delPmsBrand(id) {
  return request({
    url: '/admin/v1/pmsbrand/' + id,
    method: 'delete'
  })
}


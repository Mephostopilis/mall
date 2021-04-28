import request from '@/utils/request'
import qs from 'qs'

// 查询PmsComment列表
export function listPmsComment(query) {
  return request({
    url: '/admin/v1/pmscommentList',
    method: 'get',
    params: query
  })
}

// 查询PmsComment详细
export function getPmsComment(id) {
  return request({
    url: '/admin/v1/pmscomment/' + id,
    method: 'get'
  })
}

// 新增PmsComment
export function addPmsComment(data) {
  return request({
    url: '/admin/v1/pmscomment',
    method: 'post',
    data: data
  })
}

// 修改PmsComment
export function updatePmsComment(data) {
  return request({
    url: '/admin/v1/pmscomment',
    method: 'put',
    data: data
  })
}

// 删除PmsComment
export function delPmsComment(id) {
  return request({
    url: '/admin/v1/pmscomment/' + id,
    method: 'delete'
  })
}


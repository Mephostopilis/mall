import request from '@/utils/request'
import qs from 'qs'

// 查询PmsAlbum列表
export function listPmsAlbum(query) {
  return request({
    url: '/admin/v1/pmsalbumList',
    method: 'get',
    params: query
  })
}

// 查询PmsAlbum详细
export function getPmsAlbum(id) {
  return request({
    url: '/admin/v1/pmsalbum/' + id,
    method: 'get'
  })
}

// 新增PmsAlbum
export function addPmsAlbum(data) {
  return request({
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded'
    },
    url: '/admin/v1/pmsalbum',
    method: 'post',
    data: qs.stringify(data)
  })
}

// 修改PmsAlbum
export function updatePmsAlbum(data) {
  return request({
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded'
    },
    url: '/admin/v1/pmsalbum',
    method: 'put',
    data: data
  })
}

// 删除PmsAlbum
export function delPmsAlbum(id) {
  return request({
    url: '/admin/v1/pmsalbum/' + id,
    method: 'delete'
  })
}


import request from '@/utils/request'
import qs from 'qs'

// 查询PmsAlbumPic列表
export function listPmsAlbumPic(query) {
  return request({
    url: '/admin/v1/pmsalbumpicList',
    method: 'get',
    params: query
  })
}

// 查询PmsAlbumPic详细
export function getPmsAlbumPic(id) {
  return request({
    url: '/admin/v1/pmsalbumpic/' + id,
    method: 'get'
  })
}

// 新增PmsAlbumPic
export function addPmsAlbumPic(data) {
  return request({
    url: '/admin/v1/pmsalbumpic',
    method: 'post',
    data: data
  })
}

// 修改PmsAlbumPic
export function updatePmsAlbumPic(data) {
  return request({
    url: '/admin/v1/pmsalbumpic',
    method: 'put',
    data: data
  })
}

// 删除PmsAlbumPic
export function delPmsAlbumPic(id) {
  return request({
    url: '/admin/v1/pmsalbumpic/' + id,
    method: 'delete'
  })
}


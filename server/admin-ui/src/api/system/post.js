import request from '@/utils/request'
import qs from 'qs'

// 查询岗位列表
export function listPost(query) {
  return request({
    url: '/admin/v1/postlist',
    method: 'get',
    params: query
  })
}

// 查询岗位详细
export function getPost(postId) {
  return request({
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded'
    },
    url: '/admin/v1/post/' + postId,
    method: 'get'
  })
}

// 新增岗位
export function addPost(data) {
  return request({
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded'
    },
    url: '/admin/v1/post',
    method: 'post',
    data: qs.stringify(data)
  })
}

// 修改岗位
export function updatePost(data) {
  return request({
    url: '/admin/v1/post',
    method: 'put',
    data: data
  })
}

// 删除岗位
export function delPost(postId) {
  return request({
    url: '/admin/v1/post/' + postId,
    method: 'delete'
  })
}


import request from '@/utils/request'
import qs from 'qs'

export function login(data) {
  return request({
    url: '/admin/login',
    method: 'post',
    data: data
  })
}

export function refreshtoken(data) {
  return request({
    url: '/admin/refreshtoken',
    method: 'post',
    data: data
  })
}

export function getInfo() {
  return request({
    url: '/admin/v1/getinfo',
    method: 'get'
  })
}

export function logout() {
  return request({
    url: '/admin/v1/logout',
    method: 'post'
  })
}

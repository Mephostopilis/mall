import request from '@/utils/request'
import qs from 'qs'

export function login(data) {
  return request({
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded'
    },
    url: '/admin/login',
    method: 'post',
    data: qs.stringify(data)
  })
}

export function refreshtoken(data) {
  return request({
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded'
    },
    url: '/admin/refreshtoken',
    method: 'post',
    data: qs.stringify(data)
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
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded'
    },
    url: '/admin/v1/logout',
    method: 'post'
  })
}

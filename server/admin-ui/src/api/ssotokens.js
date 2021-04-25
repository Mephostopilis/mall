import request from '@/utils/request'

// 查询SsoToken列表
export function listSsoToken(query) {
  return request({
    url: '/admin/v1/ssotokensList',
    method: 'get',
    params: query
  })
}

// 查询SsoToken详细
export function getSsoToken(id) {
  return request({
    url: '/admin/v1/ssotokens/' + id,
    method: 'get'
  })
}

// 新增SsoToken
export function addSsoToken(data) {
  return request({
    url: '/admin/v1/ssotokens',
    method: 'post',
    data: data
  })
}

// 修改SsoToken
export function updateSsoToken(data) {
  return request({
    url: '/admin/v1/ssotokens',
    method: 'put',
    data: data
  })
}

// 删除SsoToken
export function delSsoToken(id) {
  return request({
    url: '/admin/v1/ssotokens/' + id,
    method: 'delete'
  })
}


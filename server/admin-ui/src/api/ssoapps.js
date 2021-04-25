import request from '@/utils/request'

// 查询SsoApp列表
export function listSsoApp(query) {
  return request({
    url: '/admin/v1/ssoappsList',
    method: 'get',
    params: query
  })
}

// 查询SsoApp详细
export function getSsoApp(id) {
  return request({
    url: '/admin/v1/ssoapps/' + id,
    method: 'get'
  })
}

// 新增SsoApp
export function addSsoApp(data) {
  return request({
    url: '/admin/v1/ssoapps',
    method: 'post',
    data: data
  })
}

// 修改SsoApp
export function updateSsoApp(data) {
  return request({
    url: '/admin/v1/ssoapps',
    method: 'put',
    data: data
  })
}

// 删除SsoApp
export function delSsoApp(id) {
  return request({
    url: '/admin/v1/ssoapps/' + id,
    method: 'delete'
  })
}


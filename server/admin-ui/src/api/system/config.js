import request from '@/utils/request'
import qs from 'qs'

// 查询参数列表
export function listConfig(query) {
  return request({
    url: '/admin/v1/configList',
    method: 'get',
    params: query
  })
}

// 查询参数详细
export function getConfig(configId) {
  return request({
    url: '/admin/v1/config/' + configId,
    method: 'get'
  })
}

// 根据参数键名查询参数值
export function getConfigKey(configKey) {
  return request({
    url: '/admin/v1/configKey/' + configKey,
    method: 'get'
  })
}

// 新增参数配置
export function addConfig(data) {
  return request({
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded'
    },
    url: '/admin/v1/config',
    method: 'post',
    data: qs.stringify(data)
  })
}

// 修改参数配置
export function updateConfig(data) {
  return request({
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded'
    },
    url: '/admin/v1/config',
    method: 'put',
    data: qs.stringify(data)
  })
}

// 删除参数配置
export function delConfig(configId) {
  return request({
    url: '/admin/v1/config/' + configId,
    method: 'delete'
  })
}

// 导出参数
export function exportConfig(query) {
  return request({
    url: '/admin/v1/config/export',
    method: 'get',
    params: query
  })
}

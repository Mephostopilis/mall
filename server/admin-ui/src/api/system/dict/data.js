import request from '@/utils/request'

// 查询字典数据列表
export function listData(query) {
  return request({
    url: '/admin/v1/dict/datalist',
    method: 'get',
    params: query
  })
}

// 根据字典类型查询字典数据信息
export function getDicts(dictType) {
  return request({
    url: '/admin/v1/dict/databytype/' + dictType,
    method: 'get'
  })
}

// 查询字典数据详细
export function getData(dictCode) {
  return request({
    url: '/admin/v1/dict/data/' + dictCode,
    method: 'get'
  })
}

// 新增字典数据
export function addData(data) {
  return request({
    url: '/admin/v1/dict/data',
    method: 'post',
    data: data
  })
}

// 修改字典数据
export function updateData(data) {
  return request({
    url: '/admin/v1/dict/data/',
    method: 'put',
    data: data
  })
}

// 删除字典数据
export function delData(dictCode) {
  return request({
    url: '/admin/v1/dict/data/' + dictCode,
    method: 'delete'
  })
}

// 导出字典数据
export function exportData(query) {
  return request({
    url: '/admin/v1/dict/data/export',
    method: 'get',
    params: query
  })
}

import request from '@/utils/request'

// 查询SysResource列表
export function listSysResource(query) {
  return request({
    url: '/api/v1/sysresourceList',
    method: 'get',
    params: query
  })
}

// 查询SysResource详细
export function getSysResource(id) {
  return request({
    url: '/api/v1/sysresource/' + id,
    method: 'get'
  })
}


// 新增SysResource
export function addSysResource(data) {
  return request({
    url: '/api/v1/sysresource',
    method: 'post',
    data: data
  })
}

// 修改SysResource
export function updateSysResource(data) {
  return request({
    url: '/api/v1/sysresource',
    method: 'put',
    data: data
  })
}

// 删除SysResource
export function delSysResource(id) {
  return request({
    url: '/api/v1/sysresource/' + id,
    method: 'delete'
  })
}


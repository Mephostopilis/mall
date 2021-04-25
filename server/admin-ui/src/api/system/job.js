import request from '@/utils/request'
import qs from 'qs'

// 查询SysJob列表
export function listSysJob(query) {
  return request({
    url: '/admin/v1/sysjob',
    method: 'get',
    params: query
  })
}

// 查询SysJob详细
export function getSysJob(jobId) {
  return request({
    url: '/admin/v1/sysjob/' + jobId,
    method: 'get'
  })
}

// 新增SysJob
export function addSysJob(data) {
  return request({
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded'
    },
    url: '/admin/v1/sysjob',
    method: 'post',
    data: data
  })
}

// 修改SysJob
export function updateSysJob(data) {
  return request({
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded'
    },
    url: '/admin/v1/sysjob',
    method: 'put',
    data: data
  })
}

// 删除SysJob
export function delSysJob(jobId) {
  return request({
    url: '/admin/v1/sysjob/' + jobId,
    method: 'delete'
  })
}

// 移除SysJob
export function removeJob(jobId) {
  return request({
    url: '/admin/v1/job/remove/' + jobId,
    method: 'get'
  })
}

// 启动SysJob
export function startJob(jobId) {
  return request({
    url: '/admin/v1/job/start/' + jobId,
    method: 'get'
  })
}


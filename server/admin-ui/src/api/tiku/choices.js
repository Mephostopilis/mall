import request from '@/utils/request'

// 查询Choices列表
export function listChoices(query) {
  return request({
    url: '/admin/v1/tiku/choicesList',
    method: 'get',
    params: query
  })
}

// 查询Choices详细
export function getChoices(id) {
  return request({
    url: '/admin/v1/tiku/choices/' + id,
    method: 'get'
  })
}

// 新增Choices
export function addChoices(data) {
  return request({
    url: '/admin/v1/tiku/choices',
    method: 'post',
    data: data
  })
}

// 修改Choices
export function updateChoices(data) {
  return request({
    url: '/admin/v1/tiku/choices',
    method: 'put',
    data: data
  })
}

// 删除Choices
export function delChoices(id) {
  return request({
    url: '/admin/v1/tiku/choices/' + id,
    method: 'delete'
  })
}


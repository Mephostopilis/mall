import request from '@/utils/request'

// 查询Exercise列表
export function listExercise(query) {
  return request({
    url: '/admin/v1/tiku/exerciseList',
    method: 'get',
    params: query
  })
}

// 查询Exercise详细
export function getExercise(id) {
  return request({
    url: '/admin/v1/tiku/exercise/' + id,
    method: 'get'
  })
}


// 新增Exercise
export function addExercise(data) {
  return request({
    url: '/admin/v1/tiku/exercise',
    method: 'post',
    data: data
  })
}

// 修改Exercise
export function updateExercise(data) {
  return request({
    url: '/admin/v1/tiku/exercise',
    method: 'put',
    data: data
  })
}

// 删除Exercise
export function delExercise(id) {
  return request({
    url: '/admin/v1/tiku/exercise/' + id,
    method: 'delete'
  })
}


import request from '@/utils/request'

// 查询MemberAssetsDetail列表
export function listMemberAssetsDetail(query) {
  return request({
    url: '/admin/v1/memberassetsdetailList',
    method: 'get',
    params: query
  })
}

// 查询MemberAssetsDetail详细
export function getMemberAssetsDetail(id) {
  return request({
    url: '/admin/v1/memberassetsdetail/' + id,
    method: 'get'
  })
}


// 新增MemberAssetsDetail
export function addMemberAssetsDetail(data) {
  return request({
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded'
    },
    url: '/admin/v1/memberassetsdetail',
    method: 'post',
    data: data
  })
}

// 修改MemberAssetsDetail
export function updateMemberAssetsDetail(data) {
  return request({
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded'
    },
    url: '/admin/v1/memberassetsdetail',
    method: 'put',
    data: data
  })
}

// 删除MemberAssetsDetail
export function delMemberAssetsDetail(id) {
  return request({
    url: '/admin/v1/memberassetsdetail/' + id,
    method: 'delete'
  })
}


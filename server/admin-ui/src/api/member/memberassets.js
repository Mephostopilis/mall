import request from '@/utils/request'

// 查询MemberAssets列表
export function listMemberAssets(query) {
  return request({
    url: '/admin/v1/memberassetsList',
    method: 'get',
    params: query
  })
}

// 查询MemberAssets详细
export function getMemberAssets(memberId) {
  return request({
    url: '/admin/v1/memberassets/' + memberId,
    method: 'get'
  })
}


// 新增MemberAssets
export function addMemberAssets(data) {
  return request({
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded'
    },
    url: '/admin/v1/memberassets',
    method: 'post',
    data: data
  })
}

// 修改MemberAssets
export function updateMemberAssets(data) {
  return request({
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded'
    },
    url: '/admin/v1/memberassets',
    method: 'put',
    data: data
  })
}

// 删除MemberAssets
export function delMemberAssets(memberId) {
  return request({
    url: '/admin/v1/memberassets/' + memberId,
    method: 'delete'
  })
}


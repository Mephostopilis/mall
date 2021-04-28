import request from '@/utils/request'
import qs from 'qs'

// 查询菜单列表
export function listMenu(query) {
  return request({
    url: '/admin/v1/menulist',
    method: 'get',
    params: query
  })
}

// 查询菜单详细
export function getMenu(menuId) {
  return request({
    url: '/admin/v1/menu/' + menuId,
    method: 'get'
  })
}

// 查询菜单下拉树结构
export function treeselect() {
  return request({
    url: '/admin/v1/menuTreeselect',
    method: 'get'
  })
}

// 根据角色ID查询菜单下拉树结构
export function roleMenuTreeselect(roleId) {
  return request({
    url: '/admin/v1/roleMenuTreeselect/' + roleId,
    method: 'get'
  })
}

// 新增菜单
export function addMenu(data) {
  return request({
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded'
    },
    url: '/admin/v1/menu',
    method: 'post',
    data: qs.stringify(data)
  })
}

// 修改菜单
export function updateMenu(data) {
  return request({
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded'
    },
    url: '/admin/v1/menu',
    method: 'put',
    data: qs.stringify(data)
  })
}

// 删除菜单
export function delMenu(menuId) {
  return request({
    url: '/admin/v1/menu/' + menuId,
    method: 'delete'
  })
}

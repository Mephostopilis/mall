import http from '../utils/http.js'
const appVersion = 'api/v1';


/**
 * 用户登录
 */
export const login = (params) => {
	return http.post(appVersion + '/auth/login/token',params)
}

/**
 * 用户注册
 */
export const register = (params) => {
	return http.post(appVersion + '/user/register',params)
}

/**
 * 忘记密码
 */
export const forgetPass = (params) => {
	return http.post(appVersion + '/user/retrieve/password',params)
}

/**
 * 获取用户信息
 */
export const personalInfo = (params) => {
	return http.get(appVersion + '/mall/user/info',params)
}

/**
 * 获取用户设置
 */
export const personalSet = (params) => {
	return http.get(appVersion + '/mall/user/set',params)
}

/**
 * 获取用户资产
 */
export const personalAsset = (params) => {
	return http.get(appVersion + '/mall/user/asset/'+params)
}

/**
 * 获取收款账户列表
 */
export const receiptAccount = (params) => {
	return http.get(appVersion + '/user/receipt/account',params)
}

/**
 * 获取我的邀请
 */
export const getReferee = (params) => {
	return http.get(appVersion + '/user/referee',params)
}

/**
 * 获取公告列表
 */
export const getNoticeList = (params) => {
	return http.get(appVersion + '/content/notice/list',{params})
}

/**
 * 获取收货地址列表
 */
export const getAddressList = (params) => {
	return http.get(appVersion + '/user/receipt/address',{params})
}

/**
 * 修改用户昵称
 */
export const putNickname = (params) => {
	return http.put(appVersion + '/user/nickname',params)
}

/**
 * 设置资金密码
 */
export const setPayPassword = (params) => {
	return http.post(appVersion + '/user/payPassword',params)
}

/**
 * 修改资金密码
 */
export const alterPayPassword = (params) => {
	return http.put(appVersion + '/user/payPassword',params)
}


/**
 * 保存收款账户
 */
export const userReceiptAccount = (params) => {
	return http.put(appVersion + '/user/receipt/account',params)
}
/**
 * 删除收款账户
 */
export const delettReceiptAccount = (params) => {
	return http.delete(appVersion + '/user/receipt/account/'+params)
}

/**
 * 获取资产明细
 */
export const assetsdetailList = (params) => {
	return http.get(appVersion + '/mall/user/assets_detail/list',{params})
}

/**
 * zhuanzeng
 */
export const assetsTransfer = (params) => {
	return http.post(appVersion + '/mall/user/asset/transfer',params)
}

/**
 * 获取提现额度配置
 */
export const withdrawConfig = (params) => {
	return http.get(appVersion + '/mall/withdraw/config',{params})
}

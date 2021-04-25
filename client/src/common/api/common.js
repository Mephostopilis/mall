import http from '../utils/http.js'
const appVersion = 'api/v1';

/**
 * 获取图片验证码
 */
export const getVerifyCode = (params) => {
	return http.get(appVersion + '/common/image-code', {
		params: params
	})
}

/**
 *获取加密公钥
 */
export const getPublicKey = (params) => {
	return http.get(appVersion + '/common/publicKey', {
		params: params
	})
}

/**
 *发送短信验证
 */
export const sendMobileCode = (params) => {
	return http.post(appVersion + '/common/sendMobileCode',params)
}

/**
 *找回密码发送短信验证
 */
export const sendForgetMobileCode = (params) => {
	return http.post(appVersion + '/user/retrieve/mobile',params)
}




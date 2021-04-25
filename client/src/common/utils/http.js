import Request from 'luch-request'

const http = new Request();

http.setConfig(config => { //设置全局配置
	config.baseURL = 'http://192.168.21.71:9991/'
	config.custom.isNeedToken = true
	return config;
})

http.interceptors.request.use(config => {
	/* 请求之前拦截器。可以使用async await 做异步操作 */
	const Authorization = uni.getStorageSync("Authorization");
	config.header = {
		'Content-Type': 'application/x-www-form-urlencoded',
		"Authorization": Authorization,
	}
	return config;
}, error => {
	uni.showToast({
		title: '网络请求加载超时',
		icon: 'none'
	})
	return Promise.reject(error)
})

http.interceptors.response.use((response) => {
	/* 对响应成功做点什么 可使用async await 做异步操作*/
	if (response.statusCode === 200) {
		if (!response.data) {
			return Promise.resolve(response.data)
		} else if (response.data.code === '0') {
			return Promise.resolve(response.data.data)
		} else {
			switch (response.data.code) {
				case 401:
					uni.showToast({
						title: '未授权，请登录',
						icon: 'none'
					})
					//跳转登录页面
					uni.navigateTo({
						url: '../my/login',
						fail: (err) => {
							console.log(err)
						}
					})
					break;
				case 403:
					uni.showToast({
						title: '登录过期，请重新登录',
						icon: 'none'
					})
					//跳转登录页面
					uni.navigateTo({
						url: '../my/login',
						fail: (err) => {
							console.log(err)
						}
					})
					break;
				case '500':
					uni.showToast({
						title: response.data.message,
						icon: 'none'
					})
					break;
				default:
					uni.showToast({
						title: response.data.msg,
						icon: 'none'
					})
			}
			return Promise.reject(response.data);
		}
	} else {
		return Promise.reject(response.data);
	}
}, err)

const err = function (error) {
	if (error && error.response) {
		switch (error.response.status) {
			case 400:
				uni.showToast({
					title: '错误请求',
					icon: 'none'
				})
				break;
			case 401:
				uni.showToast({
					title: '未授权，请登录',
					icon: 'none'
				})
				//跳转登录
				break;
			case 403:
				uni.showToast({
					title: '登录过期，请重新登录',
					icon: 'none'
				})
				//跳转登录
				break;
			case 404:
				uni.showToast({
					title: '请求错误,未找到该资源',
					icon: 'none'
				})
				break;
			case 405:
				uni.showToast({
					title: '请求方法未允许',
					icon: 'none'
				})
				break;
			case 408:
				uni.showToast({
					title: '请求超时',
					icon: 'none'
				})
				break;
			case 500:
				uni.showToast({
					title: '服务器端出错',
					icon: 'none'
				})
				break;
			case 501:
				uni.showToast({
					title: '网络未实现',
					icon: 'none'
				})
				break;
			case 502:
				uni.showToast({
					title: '网络错误',
					icon: 'none'
				})
				break;
			case 503:
				uni.showToast({
					title: '服务不可用',
					icon: 'none'
				})
				break;
			case 504:
				uni.showToast({
					title: '网络超时',
					icon: 'none'
				})
				break;
			case 505:
				uni.showToast({
					title: 'http版本不支持该请求',
					icon: 'none'
				})
				break;
			default:
				uni.showToast({
					title: `连接错误${error.response.status}`,
					icon: 'none'
				})
		}
	} else {
		// 超时处理
		if (JSON.stringify(error).includes('timeout')) {
			error.message = '服务器响应超时，请刷新当前页'
		} else {
			error.message = '连接服务器失败'
		}
	}
	uni.showToast({
		title: error.message,
		icon: 'none'
	})
	return Promise.reject(error)
}

export default http

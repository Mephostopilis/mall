<script>
	export default {
		onLaunch: function() {
			console.log('App Launch')
		},
		onShow: function() {
			console.log('App Show')
		},
		onHide: function() {
			console.log('App Hide')
		}
	}
</script>

<style lang="scss">
	@import "./static/css/iconfont.css";
	@import "uview-ui/index.scss";
  @import "uview-ui/theme.scss";

	page {
		background: #ffffff;
		height: 100%;
		font-family: PingFang SC;
	}

	.divider {
		position: relative;
		display: inline-block;
		width: 1rpx;
		height: 36rpx;
		margin: 0 16rpx;
		background: #e8e8e8;
		vertical-align: middle;
	}
</style>

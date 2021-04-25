import Vue from 'vue'
import App from './App'
import uView from "uview-ui"
import MescrollBody from "mescroll-uni/mescroll-body.vue";
import MescrollUni from "mescroll-uni/mescroll-uni.vue";
Vue.component('mescroll-body', MescrollBody)
Vue.component('mescroll-uni', MescrollUni)
Vue.use(uView)

Vue.config.productionTip = false

App.mpType = 'app'

const app = new Vue({
    ...App
})
app.$mount()

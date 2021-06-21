import Vue from 'vue'
import App from './App.vue'
import axios from 'axios'
import './assets/common.less'

import dataV from '@jiaminghi/data-view'
Vue.prototype.$axios = axios
Vue.config.productionTip = false

Vue.use(dataV)

new Vue({
  render: h => h(App)
}).$mount('#app')

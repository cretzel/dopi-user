import Vue from 'vue'
import App from './App.vue'
import router from './router'
import moment from 'moment';

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')

Vue.filter('formatDate', function(value) {
    if (value) {
        return moment(value).format('L') + " " + moment(value).format('LTS')
    }
});
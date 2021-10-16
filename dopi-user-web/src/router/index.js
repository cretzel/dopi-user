import Vue from 'vue'
import VueRouter from 'vue-router'
import UserList from '../views/UserList.vue'
import EditUser from '../views/EditUser.vue'
import NewUser from "@/views/NewUser";

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'UserList',
    component: UserList
  },
  {
    path: '/new',
    name: 'NewUser',
    component: NewUser
  },
  {
    path: '/:username',
    name: 'EditUser',
    component: EditUser
  }

]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router

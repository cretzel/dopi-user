import Vue from 'vue'
import VueRouter from 'vue-router'
import UserList from '../views/UserList.vue'
import EditUser from '../views/EditUser.vue'
import NewUser from "@/views/NewUser";
import Login from "@/views/Login";
import store from "../store/Store";

Vue.use(VueRouter)

const routes = [
    {
        path: '/login',
        name: 'Login',
        component: Login
    },
    {
        path: '/',
        name: 'UserList',
        component: UserList,
        meta: {
            requiresAuth: true,
            role: 'admin'
        }
    },
    {
        path: '/new',
        name: 'NewUser',
        component: NewUser,
        meta: {
            requiresAuth: true,
            role: 'admin'
        }
    },
    {
        path: '/:username',
        name: 'EditUser',
        component: EditUser,
        meta: {
            requiresAuth: true,
            role: 'admin'
        }
    },

]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
})

router.beforeEach((to, from, next) => {
    if (to.matched.some(record => record.meta.requiresAuth)) {
        const userInfo = store.userInfo
        if (userInfo == null) {
            next({path: '/login', params: {nextUrl: to.fullPath}})
            return
        }
        if (to.meta.role && !userInfo.roles.includes(to.meta.role)) {
          next({path: '/login', params: {nextUrl: to.fullPath}})
          return
        }
        next()
        return
    } 
    next()
    
})
export default router

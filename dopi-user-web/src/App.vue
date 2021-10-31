<template>
  <div id="app">

    <nav class="navbar primary" role="navigation" aria-label="main navigation">
      <div class="navbar-brand">
        <a class="navbar-item" href="/">
          <h1>Dopi</h1>
        </a>

        <a role="button" class="navbar-burger" aria-label="menu" aria-expanded="false"
           onclick="document.querySelector('.navbar-menu').classList.toggle('is-active');">
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
        </a>
      </div>

      <div class="navbar-menu">
        <div class="navbar-start">
          <router-link class="navbar-item" :to="{ name: 'UserList'}">User Management</router-link>
        </div>

        <div class="navbar-end">
          <a class="navbar-item logged-in-user" v-if="store.loggedIn">{{ store.userInfo.username }}</a>
        </div>
      </div>
    </nav>


    <section class="section">
      <div class="container">
        <div class="col-sm-12">
          <message/>
          <router-view/>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
import store from './store/Store.js'
import loginService from './service/LoginService.js'
import Message from "@/components/Message";

export default {
  name: 'App',
  components: {Message},
  data: function () {
    return {
      store: store
    }
  },
  created: function () {
    loginService.startTokenRefresher();
  }
}
</script>



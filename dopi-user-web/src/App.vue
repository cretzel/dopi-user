<template>
  <div id="app">

    <nav class="navbar primary" role="navigation" aria-label="main navigation">
          <div class="navbar-brand">
            <a class="navbar-item" href="/">
              <h1>Dopi</h1>
            </a>
        
            <a role="button" class="navbar-burger" aria-label="menu" aria-expanded="false" onclick="document.querySelector('.navbar-menu').classList.toggle('is-active');">
                <span aria-hidden="true"></span>
                <span aria-hidden="true"></span>
                <span aria-hidden="true"></span>
              </a>
          </div>
        
          <div class="navbar-menu">
            <div class="navbar-start">
              <a class="navbar-item">User Management</a>
            </div>
        
            <div class="navbar-end">
              <a class="navbar-item" v-if="store.loggedIn">{{ store.userInfo.username }}</a>              
            </div>
          </div>
    </nav>


    <section class="section">
        <div class="container">
          <div v-if="store.loggedIn" class="row">
            <div class="col-sm-12">
              <router-view/>
            </div>
          </div>
          <div v-else class="row">
            <div class="col-sm-12">
            You are not logged in. Click <a href="/">here</a> to login.
            </div>
          </div>

        </div>
    </section>
  </div>
</template>

<script>
import store from './store/Store.js'
import loginService from './service/LoginService.js'

export default {
  name: 'App',
  components: {
  },
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



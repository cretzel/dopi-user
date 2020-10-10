<template>
  <div id="app">
    <header class="row">
      <a href="/" class="logo col-sm col-md">Dopi</a>
      <label for="doc-drawer-checkbox" class="button drawer-toggle col-sm-last"></label>
      <a class="doc username col-sm-1 col-md-last" v-if="store.loggedIn">{{ store.userInfo.username }}</a>
    </header>


    <div class="row" id="doc-wrapper">
      <input id="doc-drawer-checkbox" class="drawer" value="on" type="checkbox">
      <nav class="col-md-4 col-lg-3" id="nav-drawer">
        <h4>Dopi User Management</h4>
        <label for="doc-drawer-checkbox" class="button drawer-close"></label>
        <router-link to="/">Users</router-link>
      </nav>

      <main class="col-sm-12 col-md-8 col-lg-9" id="doc-content">
        <div class="fluid">

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

        </div>
      </main>

    </div>
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



<template>
  <div id="app">
    <header class="row">
      <span class="logo col-sm-1">Dopi</span>
      <span class="col-sm-8"></span>
      <span class="col-sm-2 username" v-if="store.loggedIn">{{ store.userInfo.username }}</span>
    </header>

    <div class="row" id="doc-wrapper">

      <main class="col-sm-12 col-md-12 col-lg-12" id="doc-content">
        <div class="fluid">

          <div class="container">
          <div v-if="store.loggedIn" class="row">
            <div class="col-sm-12">
              <UserPage></UserPage>
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
import UserPage from './components/UserPage.vue'
import store from './store/Store.js'
import loginService from './service/LoginService.js'

export default {
  name: 'App',
  components: {
    UserPage
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



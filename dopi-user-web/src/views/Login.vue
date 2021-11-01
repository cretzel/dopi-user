<template>

  <div class="login">
    <h1 class="title">Login</h1>
    <div>

      <div class="columns">
        <div class="column is-4">

          <div class="field">
            <label class="label">Username</label>
            <div class="control">
              <input class="input" type="text" id="username" v-model="login.username" v-on:keyup.enter="submit"/>
            </div>
          </div>

          <div class="field">
            <label class="label">Password</label>
            <div class="control">
              <input class="input" type="password" id="password" v-model="login.password" v-on:keyup.enter="submit"/>
            </div>
          </div>

        </div>
      </div>

    </div>

    <div class="columns">
      <div class="column is-1">
        <button v-on:click="submit" id="login-button" class="button is-primary mt-3">Login</button>
      </div>
    </div>
  </div>

</template>

<script>
import loginService from '../service/LoginService.js'
import store from '../store/Store.js'

export default {
  name: 'login',
  data: function () {
    return {
      login: {}
    }
  },
  methods: {
    submit() {
      loginService.login(this.login)
          .then(() => {
            store.setMessage({text: 'Login successful', type: 'success'});
          })
          .catch(() => {
            store.setMessage({text: 'Cannot login', type: 'danger'});
          });
    }
  }
}
</script>

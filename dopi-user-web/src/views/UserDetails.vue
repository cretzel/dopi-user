<template>

  <div class="user-details">
    <h1 class="title">User Details</h1>
    <h2 class="subtitle">
      Edit the details of this user like e-mail, roles, etc.
    </h2>

    <div v-if="user">

      <div class="columns">
        <div class="column is-half">

          <div class="field">
            <label class="label">Username</label>
            <div class="control">
              <input class="input" type="text" id="username" disabled placeholder="Roles" v-bind:value="user.username"/>
            </div>
          </div>

          <div class="field">
            <label class="label">Roles</label>
            <div class="control">
              <input class="input" type="text" id="roles" placeholder="Roles" v-model="user.roles"/>
            </div>
            <p class="help">The roles or permissions a user is assigned (admin, user, ...)</p>
          </div>

        </div>
      </div>

    </div>

    <button v-on:click="saveUser" id="save" class="button is-primary mt-3">Save</button>

  </div>
</template>

<script>
import userService from '../service/UserService.js'
import store from '../store/Store.js'

export default {
  name: 'User',
  data: function () {
    return {
      user: {}
    }
  },
  methods: {
    getUser() {
      userService.getUser(this.$route.params.username)
          .then(userDto => {
            this.user = {
              username: userDto.username,
              roles: userDto.roles.join(', ')
            }
          })
          .catch(() => {
            store.setMessage({text: 'Cannot get user', type: 'danger'});
          });
    },
    saveUser() {
      let userDto = {
        username: this.user.username,
        roles: this.user.roles.split(/[,\s]+/)
      }
      userService.putUser(userDto)
          .then(() => {
            store.setMessage({text: 'User saved', type: 'success'});
          })
          .catch(() => {
            store.setMessage({text: 'Cannot save user', type: 'danger'});
          });
    }
  },
  mounted: function () {
    this.getUser();
  }
}
</script>

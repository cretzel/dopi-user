<template>
  <div class="container">
    <div class="row">

      <div class="col-sm-12">
        <h2>Users</h2>
        <table class="cardized">
          <thead>
            <tr>
              <th>Username</th>
              <th>Roles</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="item in store.users" :key="item.username">
              <td data-label="Username">
                  <router-link :to="{ name: 'UserDetails', params: { username: item.username }}">{{ item.username }}</router-link>                  
              </td>
              <td data-label="Roles">
                test
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script>
import userService from '../service/UserService.js'
import store from '../store/Store.js'

export default {
  name: 'UserList',
  props: {
  },
  data: function () {
    return {
      store: store
    }
  },
  created: function () {
    userService.getUsers();
  }
}
</script>

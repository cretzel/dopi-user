<template>
  <div class="container">
    <div class="row">

      <div class="col-sm-12">
        <h2>Users</h2>
        <table style="width:100%;">
          <thead>
            <th>Username</th>
            <th>Roles</th>
          </thead>
          <tbody>
            <tr v-for="item in store.users" :key="item.username">
              <td>
                  {{ item.username }}
              </td>
              <td>
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
  name: 'LoginForm',
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

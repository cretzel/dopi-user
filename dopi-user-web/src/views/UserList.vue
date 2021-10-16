<template>

  <div class="user-list">
    <h1 class="title">Users</h1>
    <h2 class="subtitle">
      View all your dopi users, create new users or edit them.
    </h2>

    <table class="table is-hoverable is-fullwidth">
      <thead>
        <tr>
          <th>Username</th>
          <th>Roles</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in users" :key="item.username" v-bind:data-user="item.username">
          <td>
              <router-link :to="{ name: 'UserDetails', params: { username: item.username }}">{{ item.username }}</router-link>                  
          </td>
          <td>
            {{item.roles}}
          </td>
        </tr>
      </tbody>
    </table>

  </div>
</template>

<script>
import userService from '../service/UserService.js'
import store from '../store/Store.js'

export default {
  name: 'UserList',
  data: function () {
    return {
      users: []
    }
  },
  methods: {
    getUsers() {
      userService.getUsers()
          .then(userDtos => {
            console.log(userDtos);
            this.users = userDtos;
          })
          .catch(() => {
            store.setMessage({text: 'Cannot get users', type: 'danger'});
          });
    }
  },
  created: function () {
    this.getUsers();
  }
}
</script>

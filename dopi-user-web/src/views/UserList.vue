<template>
  <div class="user-list">
    <h1 class="title">Users</h1>
    <h2 class="subtitle">
      View all your dopi users, create new users or edit them.
    </h2>

    <div class="columns is-pulled-right">
      <div class="column">
        <router-link
          id="create-user-button"
          :to="{ name: 'NewUser' }"
          tag="button"
          class="button is-primary mt-3"
          >New User</router-link
        >
      </div>
    </div>

    <table class="table is-hoverable is-fullwidth">
      <thead>
        <tr>
          <th>Username</th>
          <th>Roles</th>
          <th>Created</th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="item in users"
          :key="item.username"
          v-bind:data-user="item.username"
        >
          <td>
            <router-link
              :to="{ name: 'EditUser', params: { username: item.username } }"
              >{{ item.username }}</router-link
            >
          </td>
          <td>
            <span v-for="role in item.roles" :key="role">
              {{ role }}
            </span>
          </td>
          <td class="has-text-grey">
            {{ item.createdAt | formatDate }}
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import userService from "../service/UserService.js";
import store from "../store/Store.js";

export default {
  name: "UserList",
  data: function () {
    return {
      users: [],
    };
  },
  methods: {
    getUsers() {
      userService
        .getUsers()
        .then((userDtos) => {
          this.users = userDtos;
        })
        .catch(() => {
          store.setMessage({ text: "Cannot get users", type: "danger" });
        });
    },
  },
  created: function () {
    this.getUsers();
  },
};
</script>

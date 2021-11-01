<template>

  <div class="new-user">
    <h1 class="title">Create User</h1>
    <h2 class="subtitle">
      Enter the details of this user like e-mail, roles, etc.
    </h2>


    <div v-if="user">

      <div class="columns">
        <div class="column is-half">

          <div class="field">
            <label class="label">Username</label>
            <div class="control">
              <input class="input" type="text" id="username" placeholder="Username" v-model="user.username" v-on:keyup.enter="saveUser"/>
            </div>
          </div>

          <div class="field">
            <label class="label">Roles</label>
            <div class="control">
              <input class="input" type="text" id="roles" placeholder="admin, user" v-model="user.roles" v-on:keyup.enter="saveUser"/>
            </div>
            <p class="help">The roles or permissions a user is assigned (admin, user, ...)</p>
          </div>

          <div class="field">
            <label class="label">Password</label>
            <div class="control">
              <input class="input" type="password" id="password" v-model="user.password" v-on:keyup.enter="saveUser"/>
            </div>
          </div>

        </div>
      </div>

    </div>

    <button v-on:click="saveUser" id="save" class="button is-primary mt-3">Save</button>

  </div>
</template>

<script>
import userService from "../service/UserService.js";
import store from "../store/Store.js";
import router from "@/router";

export default {
  name: "NewUser",
  data: function () {
    return {
      user: {
        username: "",
        roles: "",
      },
    };
  },
  methods: {
    validate(userDto) {
      const errors = [];
      if (!userDto.username) {
        errors.push("Username is required")
      }
      if (userDto.roles.length == 0) {
        errors.push("Roles is required")
      }
      if (!userDto.password) {
        errors.push("Password is required")
      }
      return errors;
    },

    saveUser() {
      let userDto = {
        username: this.user.username,
        roles: this.user.roles ? this.user.roles.split(/[,\s]+/) : [],
        password: this.user.password,
      };
           
      const errors = this.validate(userDto);
      if (errors.length > 0) {
        store.setMessage({ texts: errors, type: "danger" });
        return
      }
      userService
        .postUser(userDto)
        .then(() => {
          router.push({ name: "UserList" });
          store.setMessage({ text: "User created", type: "success" });
        })
        .catch((e) => {
          console.log(e);
          store.setMessage({ text: e, type: "danger" });
        });
    },
  },
};
</script>

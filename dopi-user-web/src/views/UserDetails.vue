<template>

<div>
    <h1 class="title">User Details</h1>
    <h2 class="subtitle">
      Edit the details of this user like e-mail, roles, etc.
    </h2>

    <div v-if="store.user">

      <div class="columns">
        <div class="column is-half">

          <div class="field">
            <label class="label">Username</label>
            <div class="control">
              <input class="input" type="text" id="Roles" disabled placeholder="Roles" v-bind:value="store.user.username" />
            </div>
          </div>    

          <div class="field">
            <label class="label">Roles</label>
            <div class="control">
              <input class="input" type="text" id="Roles" placeholder="Roles" v-bind:value="store.user.roles" />
            </div>
            <p class="help">The roles or permissions a user is assigned</p>
          </div>    

        </div>
      </div>

    </div>


    <button v-on:click="toggleEditMode" class="button is-primary mt-3">Save</button>

</div>
</template>

<script>
import userService from '../service/UserService.js'
import store from '../store/Store.js'

export default {
  name: 'User',
  props: {
  },
  data: function () {
    return {
      editMode: false,
      store: store
    }
  },
  created: function () {
    store.setUser(null);
    userService.getUser(this.$route.params.username);
  },
  methods: {
    toggleEditMode: function() {
      this.editMode = !this.editMode;
    }
  }

 
}
</script>

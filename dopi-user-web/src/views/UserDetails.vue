<template>
  <div class="container">
    <div class="row">

      <div class="col-sm-12">
        <div v-if="store.user">
          <h2>{{store.user.username}}</h2>

          <form>
              <div class="row responsive-label">
                <div class="col-sm-12 col-md-3">
                  <label for="username">Username</label>
                </div>
                <div class="col-sm-12 col-md">
                  <span v-bind:class="{hidden: editMode}">{{ store.user.username }}</span>
                  <input type="text" id="Username" placeholder="Username" v-bind:value="store.user.username" v-bind:class="{invisible: !editMode}"/>
                </div>

              </div>

              <div class="row responsive-label">
                <div class="col-sm-12 col-md-3">
                  <label for="roles">Roles</label>
                </div>
                <div class="col-sm-12 col-md">
                  <span v-bind:class="{hidden: editMode}">{{ store.user.roles }}</span>
                  <input type="text" id="Roles" placeholder="Roles" v-bind:value="store.user.roles" v-bind:class="{invisible: !editMode}"/>
                </div>
              </div>
          </form>

          <button v-on:click="toggleEditMode" v-bind:class="{hidden: editMode}">Edit</button>
          <button v-on:click="toggleEditMode" v-bind:class="{hidden: !editMode}">Save</button>

        </div>
      </div>
    </div>
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

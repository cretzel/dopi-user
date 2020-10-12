<template>
  <div class="container">
    <div class="row">

      <div class="col-sm-12">
        <div v-if="store.user">
          <h2>{{store.user.username}}</h2>

          <table class=" horizontal">
            <thead>
                <tr>
                  <th>Username</th>
                  <th>Roles</th>
                </tr>
            </thead>
            <tbody>
              <tr>
                <td data-label="Username">
                    {{ store.user.username }}
                </td>
                <td data-label="Roles">
                  {{ store.user.roles }}
                </td>
              </tr>
            </tbody>
          </table>
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
      store: store
    }
  },
  created: function () {
    store.setUser(null);
    userService.getUser(this.$route.params.username);
  }

 
}
</script>

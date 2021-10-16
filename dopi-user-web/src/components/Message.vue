<template>
  <article class="message" :class="stateType" v-if="text != null && stateVisisble">
    <div class="message-body">{{ text }}</div>
  </article>
</template>

<script>
import store from '../store/Store.js'

export default {
  name: 'message',
  data: function () {
    return {
      stateVisisble: true
    }
  },
  computed: {
    stateType() {
      return 'is-' + (store.message != null ? store.message.type : 'none');
    },
    text() {
      return store.message != null ? store.message.text : null;
    }
  },
  watch: {
    text: function () {
      window.setTimeout(() => store.setMessage(null), 5000);
    }
  }
}
</script>

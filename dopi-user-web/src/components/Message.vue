<template>
  <article class="message" :class="stateType" v-if="texts.length > 0 && stateVisisble">
    <div class="message-body">
      <div v-for="text in texts" :key="text">
      {{ text }}
      </div>
    </div>
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
    texts() {
      const texts = [];
      if (store.message) {
        if (store.message.text) {
          texts.push(store.message.text);
        }
        if (store.message.texts) {
          store.message.texts.forEach(element => {
            texts.push(element);
          });
        }
      }
      return texts;
    }
  },
  watch: {
    texts: function () {
      window.setTimeout(() => store.setMessage(null), 5000);
    }
  }
}
</script>

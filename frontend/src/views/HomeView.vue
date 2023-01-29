<template>
  <default-layout>
    <template v-slot:sidebar-content>
      <vhost-list>

      </vhost-list>
    </template>
    <template v-slot:main-content>Select a vhost or create a new one from the list</template>
  </default-layout>
</template>

<script>
import DefaultLayout from '../components/layouts/DefaultLayout.vue'
import VhostList from '../components/list/VhostList.vue'

export default {
  name: "HomeView",
  components: {
    'default-layout': DefaultLayout,
    'vhost-list': VhostList
  }
}
</script>

<style scoped>

</style>
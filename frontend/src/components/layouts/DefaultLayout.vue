
<template>
  <div class="flex h-screen">
    <div class="w-1/4 bg-gray-300">
      <div class="h-full overflow-y-auto relative">
        <div class="p-4">
          <slot name="sidebar-content"></slot>
        </div>
        <div class="bg-gray-400 p-4 flex justify-between absolute inset-x-0 bottom-0">
          <!-- Add a plus icon link that take to vue route /vhost/single -->
          <a href="#" class="text-lg text-gray-700" @click="addVhost()">
            <i class="fa fa-plus"></i>
          </a>
          <!-- Add a refresh icon -->
          <a href="#" class="text-lg text-gray-700" @click="restartApache()">
            <i class="fa fa-refresh"></i>
          </a>
        </div>
      </div>
    </div>
    <div class="w-3/4 p-4">
      <slot name="main-content"></slot>
    </div>
  </div>
</template>

<script>

import {RestartApache} from "../../../wailsjs/go/controller/VhostController";

export default {
  name: "DefaultLayout",
  methods: {
    addVhost: function (){
      this.$router.push({ name: 'vhost' })
    },
    restartApache: function (){
      RestartApache()
        .then(function (response) {
          console.log(response)
        });
    }
  }
}
</script>

<style scoped>

</style>
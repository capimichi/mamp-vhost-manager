<template>
  <div>
    <div v-for="vhost in vhosts" :key="vhost.Name" class="text-lg font-medium cursor-pointer">
      <div class="vhost-content py-2">
        <span class="vhost" @click="goToVhost(vhost.Name)">
          {{ vhost.Name }}
        </span>
        <button @click="deleteVhost(vhost.Name)" class="ml-4 text-red-500 cursor-pointer">
          <i class="fa fa-trash"></i>
        </button>
      </div>
      <hr>
    </div>
  </div>
</template>


<script>
import {GetVhosts} from '../../../wailsjs/go/controller/VhostController';
import {DeleteVhost} from '../../../wailsjs/go/controller/VhostController';

export default {
  name: "VhostList",
  data() {
    return {
      vhosts: []
    }
  },
  created: function () {
    this.fetchVhosts();
  },
  updated: function () {
    this.fetchVhosts();
  },
  methods: {
    goToVhost: function (vhostName){
      this.$router.push({ name: 'vhost', query: { name: vhostName } })
    },
    deleteVhost: function (vhostName){
      DeleteVhost(vhostName);
      this.$router.push({ name: 'home' })
    },
    fetchVhosts: function () {
      GetVhosts()
          .then((result) => {
            // Update result with data back from App.Greet()
            this.vhosts = result;
          })
          .catch((err) => {
            console.error(err);
          });
    }
  }
}
</script>

<style scoped>
.cursor-pointer {
  cursor: pointer;
}
</style>
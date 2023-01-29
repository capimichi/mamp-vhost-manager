<template>
  <default-layout>
    <template v-slot:sidebar-content>
      <vhost-list></vhost-list>
    </template>
    <template v-slot:main-content>
      <form class="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4" @submit.prevent="submitForm">
        <h1 class="text-3xl">Vhost</h1>
        <div class="mb-4">
          <label class="block text-gray-700 text-sm font-bold mb-2" for="name">
            Name
          </label>
          <input v-model="vhost.Name" class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" id="name" type="text" placeholder="Name">
        </div>
        <div class="mb-6">
          <label class="block text-gray-700 text-sm font-bold mb-2" for="documentRoot">
            ServerName
          </label>
          <input v-model="vhost.ServerName" class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" id="serverName" type="text" placeholder="ServerName">
          <p class="text-gray-600 text-xs italic">ServerName is the name in the url "example.locale"</p>
        </div>
        <div class="mb-6">
          <label class="block text-gray-700 text-sm font-bold mb-2" for="documentRoot">
            DocumentRoot
          </label>
          <input v-model="vhost.DocumentRoot" class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" id="documentRoot" type="text" :placeholder="this.baseDocumentRoot">
          <p class="text-gray-600 text-xs italic">DocumentRoot is the path to your web files</p>
        </div>
        <div class="flex items-center justify-between">
          <button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline" type="submit">
            Save
          </button>
        </div>
      </form>
    </template>
  </default-layout>
</template>

<script>
import DefaultLayout from '../components/layouts/DefaultLayout.vue'
import VhostList from '../components/list/VhostList.vue'
import {GetVhost} from "../../wailsjs/go/controller/VhostController";
import {CreateVhost} from "../../wailsjs/go/controller/VhostController";
import {GetGuessDocumentRoot} from "../../wailsjs/go/controller/VhostController";

export default {
  name: "SingleVhostView",
  data: function () {
    return {
      vhost: {},
      baseDocumentRoot: ''
    }
  },
  created: function () {
    this.fetchVhost();
    this.fetchBaseDocumentRoot();
  },
  updated: function () {
    this.fetchVhost();
    this.fetchBaseDocumentRoot();
  },
  methods: {
    fetchBaseDocumentRoot: function () {
      GetGuessDocumentRoot()
          .then((result) => {
            // Update result with data back from App.Greet()
            this.baseDocumentRoot = result;
          })
          .catch((err) => {
            console.error(err);
          });
    },
    fetchVhost: function () {
      var vhostName = this.$route.query.name;
      // Check if name is empty
      if (vhostName === undefined) {
        this.vhost = {};
      }
      GetVhost(vhostName)
          .then((result) => {
            // Update result with data back from App.Greet()
            this.vhost = result;
          })
          .catch((err) => {
            console.error(err);
          });
    },
    submitForm() {
      CreateVhost(this.vhost.Name, this.vhost.ServerName, '', this.vhost.DocumentRoot)
        .then((result) => {
          // Update result with data back from App.Greet()
          this.$router.push({ name: 'vhost', query: { name: result.Name } })
        });
    }
  },
  components: {
    'default-layout': DefaultLayout,
    'vhost-list': VhostList
  }
}
</script>

<style scoped>

</style>
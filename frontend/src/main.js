import {createApp} from 'vue'
import App from './App.vue'
import HomeView from './views/HomeView.vue'
import router from './router'
import '../node_modules/font-awesome/css/font-awesome.css'


const app = createApp(App)
app.use(router)
app.mount('#app')

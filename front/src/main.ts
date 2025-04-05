import { createApp } from 'vue';
import App from './App.vue';
import router from './router.js';
import { Buffer } from 'buffer';
import './utils/config_md.js';

globalThis.Buffer = Buffer;

const app = createApp(App);
app.use(router);
app.mount('#app');

// Vue and main app
import { createApp } from "vue";
import App from "./App.vue";

// Router
import router from "./router";

// PrimeVue and its components
import PrimeVue from "primevue/config";
import Panel from "primevue/panel";
import Menubar from 'primevue/menubar';
import Dialog from 'primevue/dialog';
import Dropdown from 'primevue/dropdown';
import Button from "primevue/button";

// CSS
import "./index.css";
import 'primeicons/primeicons.css' // Icons of PrimeVue

// Presets
import Lara from "./presets/lara";

// Create app
const app = createApp(App);

// Use plugins
app.use(PrimeVue, { pt: Lara, ripple: true });
app.use(router);

// Register components
app.component("Panel", Panel);
app.component('Menubar', Menubar);
app.component('Dialog', Dialog);
app.component('Dropdown', Dropdown);
app.component('Button', Button);

// Mount app
app.mount("#app");
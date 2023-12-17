import { createApp } from "vue";
import "./index.css";
import App from "./App.vue";
import PrimeVue from "primevue/config";
import Panel from "primevue/panel";
import Lara from "./presets/lara";

import Menubar from 'primevue/menubar';

const app = createApp(App);
app.component("Panel", Panel);
app.use(PrimeVue, { unstyled: true, pt: Lara });


app.component('Menubar', Menubar);

app.mount("#app");
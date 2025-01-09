import "./assets/main.css";

import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import PrimeVue from "primevue/config";
import Aura from "@primevue/themes/aura";
import ToastService from "primevue/toastservice";
import DialogService from "primevue/dialogservice";

import { VueRecaptchaPlugin } from "vue-recaptcha";
import { library } from "@fortawesome/fontawesome-svg-core";
// import { faPhone } from "@fortawesome/free-solid-svg-icons";
import { fas } from "@fortawesome/free-solid-svg-icons";
// import { far } from "@fortawesome/free-regular-svg-icons";
import { fab } from "@fortawesome/free-brands-svg-icons";
import store from "./store";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
//Add all icons to the library so you can use it in your page
library.add(fas, fab);
const app = createApp(App);

// app.component("font-awesome-icon", FontAwesomeIcon);
app.use(router);
app.use(store);
app.use(PrimeVue, {
  theme: {
    preset: Aura,
  },
});
app.use(ToastService);
app.use(DialogService);

// app.use(VueRecaptchaPlugin, {
//   v3SiteKey: "6LfuuYgqAAAAAOPnPbRKpJM3DWOyEy2rJagWTb0V",
// });
app.mount("#app");

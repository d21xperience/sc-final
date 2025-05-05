import "./assets/main.css";

import { createApp } from "vue";
import App from "./App.vue";
import PrimeVue from "primevue/config";
import Aura from "@primeuix/themes/aura";
import ToastService from "primevue/toastservice";
import DialogService from "primevue/dialogservice";
import Tooltip from "primevue/tooltip";
import Ripple from "primevue/ripple";
import { VueRecaptchaPlugin } from "vue-recaptcha";
import { library } from "@fortawesome/fontawesome-svg-core";
// import { faPhone } from "@fortawesome/free-solid-svg-icons";
import { fas } from "@fortawesome/free-solid-svg-icons";
// import { far } from "@fortawesome/free-regular-svg-icons";
import { fab } from "@fortawesome/free-brands-svg-icons";
import router from "./router";
import store from "./store";
//Add all icons to the library so you can use it in your page
library.add(fas, fab);
const app = createApp(App);
app.use(router);
app.use(store);
// app.use(PrimeVue, {
//   unstyled: true,
//   // theme: {
//   //   preset: Aura,
//   //   // options: {
//   //   //   prefix: "p",
//   //   //   darkModeSelector: "system",
//   //   //   cssLayer: false,
//   //   // },
//   // },
//   ripple: true,
// });
app.use(PrimeVue, {
  unstyled: true,
 
});
app.use(ToastService);
app.use(DialogService);
app.directive("tooltip", Tooltip);
app.directive("ripple", Ripple);
// app.use(VueRecaptchaPlugin, {
//   v3SiteKey: "6LfuuYgqAAAAAOPnPbRKpJM3DWOyEy2rJagWTb0V",
// });
app.mount("#app");
// window.store = store;

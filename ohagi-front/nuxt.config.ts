import { defineNuxtConfig } from "nuxt3";

// https://v3.nuxtjs.org/docs/directory-structure/nuxt.config
export default defineNuxtConfig({
  srcDir: "src/",
  meta: {
    script: [
      {
        hid: "google-sign-in",
        src: "https://accounts.google.com/gsi/client",
        async: true,
        defer: true,
      },
    ],
  },
});

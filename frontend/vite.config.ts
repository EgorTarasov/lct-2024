import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";
import { TanStackRouterVite } from "@tanstack/router-vite-plugin";
import path from "path";
import svgr from "vite-plugin-svgr";
import { VitePWA } from "vite-plugin-pwa";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    react(),
    TanStackRouterVite(),
    svgr({
      exportAsDefault: true,
      svgrOptions: {
        exportType: "default"
      }
    }),
    VitePWA({
      registerType: "autoUpdate",
      workbox: {
        runtimeCaching: [
          {
            urlPattern: /^https:\/\/api\.lct\.larek\.tech\/consumers\/q$/,
            handler: 'NetworkFirst',
            options: {
              cacheName: 'api-cache',
              expiration: {
                maxEntries: 50, // Number of entries to store in cache
                maxAgeSeconds: 24 * 60 * 60, // Cache for 1 day
              },
              cacheableResponse: {
                statuses: [0, 200], // Cache successful responses and opaque responses
              },
            },
          },
        ],
      }
    })
    // basicSsl()
  ],
  build: {
    target: "esnext"
  },
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
      api: path.resolve(__dirname, "./src/api")
    }
  },
  server: {
    proxy: {}
  }
});

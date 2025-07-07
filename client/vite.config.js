import path from "path";
import { fileURLToPath } from "url";
import react from "@vitejs/plugin-react";
import { defineConfig } from "vite";

const __dirname = path.dirname(fileURLToPath(import.meta.url));

export default defineConfig({
  plugins: [react()],

  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
    },
  },

  build: {
    outDir: "dist",
    sourcemap: true,
    minify: "esbuild",
    target: "esnext",
    rollupOptions: {
      output: {
        manualChunks: {
          vendor: ["react", "react-dom", "react-router-dom"],
          ui: ["sonner"],
        },
      },
    },
  },

  base: "/",

  server: {
    port: 3000,
    host: true,
    historyApiFallback: true,
  },

  preview: {
    port: 4173,
    host: true,
    historyApiFallback: true,
  },

  define: {
    global: "globalThis",
  },
});

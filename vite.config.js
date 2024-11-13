import { defineConfig } from 'vite'

export default defineConfig({
    assetsInlineLimit: 0,
    build: {
        sourcemap: true,
        outDir: 'dist',
        rollupOptions: {
            input: [
                'index.html',
                'pages/game.html',
            ],
        },
    },
})


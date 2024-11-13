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
                'style/tailwind.css',
                'style/hompage.css',
                'scripts/game_engine/board_generation.js',
                'scripts/game_engine/game.js',
            ],
        },
    },
})


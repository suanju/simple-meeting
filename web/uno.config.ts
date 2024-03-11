import { defineConfig } from 'unocss'
import presetUno from '@unocss/preset-uno'
import presetIcons from '@unocss/preset-icons'

export default defineConfig({
  // ...UnoCSS选项
  presets: [
    presetUno(),
    presetIcons({
      collections: {
        systemUicons: () => import('@iconify-json/system-uicons/icons.json').then((i) => i.default)
      }
    })
  ]
})

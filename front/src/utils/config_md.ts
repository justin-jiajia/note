import { config, XSSPlugin } from 'md-editor-v3';

config({
  markdownItPlugins(plugins) {
    return [
      ...plugins,
      {
        type: 'xss',
        plugin: XSSPlugin,
        options: {},
      },
    ];
  },
});
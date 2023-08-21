// webpack.config.js

module.exports = {
    // ...
    module: {
      rules: [
        // ...
        {
          test: /\.js$/,
          exclude: /node_modules\/(?!(timeago\.js)\/).*/,
          use: {
            loader: 'source-map-loader',
            options: {
              filterSourceMappingUrl: (url, resourcePath) => {
                if (/timeago\.js/.test(resourcePath)) {
                  return false;
                }
                return true;
              },
            },
          },
        },
        // ...
      ],
    },
    // ...
  };
  
module.exports = {
  mode: 'production',
  module: {
    rules: [
      {
        test: /\.scss/,
        use: [
          'style-loader',
          {
            loader: 'css-loader',
            options: {
              url: false,
              sourceMap: true,
              importLoaders: 2
            }
          },
          {
            loader: 'sass-loader',
            options: {
              includePaths: ['./node_modules'],
              sourceMap: true
            }
          }
        ]
      }
    ]
  }
}

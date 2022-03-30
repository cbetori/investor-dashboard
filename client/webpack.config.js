const { ModuleFederationPlugin } = require('webpack').container
const { MFLiveReloadPlugin } = require('@module-federation/fmr')
const HtmlWebPackPlugin = require('html-webpack-plugin')
const path = require('path')
const web = require('webpack')
const deps = require('./package.json').dependencies
const dotenv = require('dotenv').config({
  path: '../.env',
}).parsed

const mode = process.env.NODE_ENV || 'development'
const prod = mode === 'production'

module.exports = {
  entry: './src/index.js',
  mode,
  output: {
    path: __dirname + '/dist',
    filename: 'index.js',
    chunkFilename: '[name].[id].js',
    publicPath: prod
      ? 'https://aca-dashboard-go.herokuapp.com/'
      : 'http://localhost:3004/',
  },
  devtool: prod ? 'hidden-source-map' : 'hidden-source-map',
  resolve: {
    extensions: ['.js', '.json', '.css'],
  },
  module: {
    rules: [
      {
        test: /\.js$/,
        include: path.resolve(__dirname, 'src'),
        exclude: /node_modules/,
        loader: 'babel-loader',
        options: {
          presets: ['@babel/preset-react'],
        },
      },
      {
        test: /\.(png|svg|ico)$/i,
        type: 'asset/resource',
      },
      {
        test: /\.css$/,
        use: ['style-loader', 'css-loader'],
      },
    ],
  },
  devServer: {
    port: 3004,
    historyApiFallback: true,
  },
  plugins: [
    new ModuleFederationPlugin({
      name: 'dashboard',
      filename: 'remoteEntry.js',
      exposes: {
        './Dashboard': './src/Store.js',
      },
      shared: {
        react: {
          singleton: true,
          requiredVersion: deps.react,
        },
        'react-dom': {
          singleton: true,
          requiredVersion: deps['react-dom'],
        },
      },
    }),
    new HtmlWebPackPlugin({
      template: './public/index.html',
    }),
    ...(prod
      ? []
      : new MFLiveReloadPlugin({
          container: 'portfolio',
          port: 8080,
        })),
    new web.DefinePlugin({
      API_URL: prod
        ? JSON.stringify(process.env.API_URL)
        : JSON.stringify(dotenv.API_URL),
    }),
  ],
}

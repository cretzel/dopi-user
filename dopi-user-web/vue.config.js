const devServer = require('./src/mock/devServer')

module.exports = {
    publicPath: "/users",
    devServer: {
      before: devServer,
      proxy: {
        '/api/user/users': {
          target: 'http://localhost:8081'
        }
    }
  }
}
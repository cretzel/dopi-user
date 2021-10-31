

module.exports = {
    publicPath: "/users",
    devServer: {
        proxy: {
            '/api/user': {
                target: 'http://localhost:8081'
            }
        }
    }
}
const bodyParser = require('body-parser')

module.exports = app => {
  app.use(bodyParser.json())
  app.post('/api/user/refresh', (req, res) => {
    res.cookie('dopi_jwt',"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiUm9sZXMiOlsiUk9MRV9BRE1JTiIsIlJPTEVfVVNFUiIsIlJPTEVfQVBQIl0sImV4cCI6MTYzMzQ2OTI1MiwiaWF0IjoxNjAxOTMzMjUyLCJpc3MiOiJkb3BpIn0.ocEklJPavpgTirZMuVAEZmHJlkQk1wGrX6IAfvKqQKk");
    res.send(JSON.stringify({username: "admin"}))
  })
  app.get('/api/user/users', (req, res) => {
      res.send(JSON.stringify([
        {username:"admin",roles:["admin"]},
        {username:"jim",roles:["user"]},
        {username:"joe",roles:["user"]},
    ]
      ))
  })
}
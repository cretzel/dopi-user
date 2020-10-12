const bodyParser = require('body-parser')

const users = [
  {username:"admin",roles:["admin"]},
  {username:"jim",roles:["user"]},
  {username:"joe",roles:["user"]},
];

module.exports = app => {
  app.use(bodyParser.json())
  app.route('/api/user/refresh').post((req, res) => {
    res.cookie('dopi_jwt',"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiUm9sZXMiOlsiUk9MRV9BRE1JTiIsIlJPTEVfVVNFUiIsIlJPTEVfQVBQIl0sImV4cCI6MTYzMzQ2OTI1MiwiaWF0IjoxNjAxOTMzMjUyLCJpc3MiOiJkb3BpIn0.ocEklJPavpgTirZMuVAEZmHJlkQk1wGrX6IAfvKqQKk");
    res.send(JSON.stringify({username: "admin"}))
  });
  app.route('/api/user/users/:username').get((req, res) => {
    const username = req.params.username;
    console.log("get user " + req.params.username);
    const user = users.find(u => u.username == username);
    res.send(JSON.stringify(user));
  });
  app.route('/api/user/users').get((req, res) => {
    console.log("get users");
    res.send(JSON.stringify(users));
  });

}
const bodyParser = require('body-parser')

const users = [
  {username:"admin",roles:["admin"]},
  {username:"jim",roles:["user"]},
  {username:"joe",roles:["user"]},
];

module.exports = app => {
  app.use(bodyParser.json())
  app.route('/api/user/refresh').post((req, res) => {
    res.cookie('dopi_jwt',"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiUm9sZXMiOlsiYWRtaW4iXSwiZXhwIjoxNjMzOTAzNDA2LCJpYXQiOjE2MzM5MDI1MDYsImlzcyI6ImRvcGkifQ.1K-v9jOp1UsxNuFP32nEblhD-gCQwkMiof74cMLxoPU");
    res.send(JSON.stringify({username: "admin"}))
  });

}
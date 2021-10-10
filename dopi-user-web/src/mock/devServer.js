const bodyParser = require('body-parser')

const users = [
  {username:"admin",roles:["admin"]},
  {username:"jim",roles:["user"]},
  {username:"joe",roles:["user"]},
];

module.exports = app => {
  app.use(bodyParser.json())
  app.route('/api/user/refresh').post((req, res) => {
    res.cookie('dopi_jwt',"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiUm9sZXMiOlsiYWRtaW4iXSwiZXhwIjoxNjMzODg4ODY4LCJpYXQiOjE2MzM4ODc5NjgsImlzcyI6ImRvcGkifQ.d0GFM8MOiFtZwlrHEC4EzCnmfqJC4_it6AkcZ67e6d4");
    res.send(JSON.stringify({username: "admin"}))
  });

}
module.exports = app => {
  app.route('/api/user/refresh').post((req, res) => {
    res.cookie('dopi_jwt',"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiUm9sZXMiOlsiYWRtaW4iXSwiZXhwIjoyNzMzNTE2NjAzLCJpYXQiOjE2MzM5MDM2NzEsImlzcyI6ImRvcGkifQ.OexkEffWjesgZXUtzTxbontONA3YC8T-4ld_vv4sXvM");
    res.send(JSON.stringify({username: "admin"}))
  });

}
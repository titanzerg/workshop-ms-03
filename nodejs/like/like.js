const server = require('express')()

server.get('/:id', function (req, res) {
  // update feed's status
  // ...
  // response
  response = {
    message: "success",
    count: 100
  };
  res.end(JSON.stringify(response));
})

server.put('/like/:id', function (req, res) {
  // update feed's status
  // ...
  // response
  response = {
    message: "success"
  };
  res.end(JSON.stringify(response));
})

server.listen(process.env.PORT || 3000)

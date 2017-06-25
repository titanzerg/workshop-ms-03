const server = require('express')()

server.get('/search/:name', function (req, res) {
    response = [{
        user_id:"123456",
        first_name: req.params.name,
        last_name: "test"
    },{
        user_id:"111111",
        first_name: req.params.name+"xxx",
        last_name: "test xxx"
    }];
    res.writeHead(200, {"Content-Type": "application/json"});
    res.end(JSON.stringify(response));
})



server.listen(process.env.PORT || 3000)
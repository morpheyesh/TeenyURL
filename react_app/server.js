var express = require('express');
var app = express();

app.set('port', (process.env.PORT || 9000));
app.use(express.static(__dirname + '/dist'));


app.get('/app', function(req, res) {
  res.sendFile('./dist/index.html', {"root": __dirname});
});

app.listen(app.get('port'), function() {
  console.log("Node app is running at localhost:" + app.get('port'));
});

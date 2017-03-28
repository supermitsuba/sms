var args = process.argv.slice(2);

var request = require('request');
var util = require('util');
console.log(args[1])
console.log(args[0])
request(args[1], function (error, response, responsebody) {
  console.log('error:', error); // Print the error if one occurred 
  console.log('statusCode:', response && response.statusCode); // Print the response status code if a response was received 
  console.log('body:', responsebody); // Print the HTML for the Google homepage. 
  var body = JSON.parse(responsebody)
  var temperature = (body.main.temp * 9 / 5) - 459.17;
  var conditions = body.weather[0].main
  var body = { "text":util.format("Now:     %d F   %s", temperature.toFixed(0), conditions), "duration":30 }

  request({
      "uri":args[0], 
      "method":"POST",
      "body": JSON.stringify(body)
    }, function (error, response, body) {
        console.log('error:', error); // Print the error if one occurred 
        console.log('statusCode:', response && response.statusCode); // Print the response status code if a response was received 
        console.log('body:', body); // Print the HTML for the Google homepage. 
    })
});
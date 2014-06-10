'use strict';

var fs = require('fs');
var split = require('split');

var stream = fs.createReadStream(__dirname + '/data/allTraining.txt').pipe(split(' '));

var parts = {};
stream
  .on('data', function(d) {
    d = d.split('_');
    var word = d[0];
    var pos = d[1];
  })
  .on('end', function() {
    console.log('done');
  });
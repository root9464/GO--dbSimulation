import * as fs from 'fs';
let rawData = fs.readFileSync('./data.json', 'utf8');
let jsonData = JSON.parse(rawData);

console.log(jsonData);
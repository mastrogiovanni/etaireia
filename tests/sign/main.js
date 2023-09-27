var crypto = require('crypto');
var fs = require('fs');
const { generateKeyPairSync } = require('crypto'); 

const { publicKey, privateKey } = generateKeyPairSync('rsa', 
{   modulusLength: 2048,  // the length of your key in bits   
    publicKeyEncoding: {
      type: 'spki',       // recommended to be 'spki' by the Node.js docs
      format: 'pem'   
    },   
    privateKeyEncoding: {
      type: 'pkcs8',      // recommended to be 'pkcs8' by the Node.js docs
      format: 'pem',
      //cipher: 'aes-256-cbc',   // *optional*
      //passphrase: 'top secret' // *optional*   
  } 
}); 
console.log(privateKey); 

// var pem = fs.readFileSync('../asn/test.pem');
// var key = pem.toString('ascii');

var sign = crypto.createSign('RSA-SHA256');
sign.update('abcdef');  // data from your file would go here
var sig = sign.sign(privateKey, 'hex');

console.log(sig)
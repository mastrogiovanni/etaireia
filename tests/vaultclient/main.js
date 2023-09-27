var options = {
    apiVersion: 'v1',
    endpoint: 'http://localhost',
    token: 'hvs.oXoXejkgBNrrgaNyAyQwnlk3'
};

process.env.DEBUG = 'node-vault'; // switch on debug mode

// get new instance of the client
var vault = require("node-vault")(options);

//   // init vault server
//   vault.init({ secret_shares: 1, secret_threshold: 1 })
//   .then( (result) => {
//     var keys = result.keys;
//     // set token for all following requests
//     vault.token = result.root_token;
//     // unseal vault server
//     return vault.unseal({ secret_shares: 1, key: keys[0] })
//   })
//   .catch(console.error);

async function bootstrap() {
    console.log("start")

    // vault.write('secret/hello', { value: 'world', lease: '10h' })
    // .then(() => vault.read('secret/hello'))
    // .then((result) => vault.revoke({ lease_id: result.lease_id }))
    // .then(() => vault.revokePrefix({ path_prefix: 'secret' }))
    // .catch((err) => console.error(err.message));

    console.log(await vault.read('test/data/pippo'))
}

bootstrap().catch(console.log)

// vault.write('secret/hello', { value: 'world', lease: '1s' })
// .then( () => )
// .then( () => vault.delete('secret/hello'))
// .catch(console.error);
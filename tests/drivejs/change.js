const fs = require('fs');
const { GoogleAuth } = require('google-auth-library');
const { google } = require('googleapis');
const path = require('path');
const { authenticate } = require('@google-cloud/local-auth');

const TOKEN_PATH = path.join(process.cwd(), 'token.json');
const CREDENTIALS_PATH = path.join(process.cwd(), 'credentials.json');
const SCOPES = ['https://www.googleapis.com/auth/drive'];

/**
 * Reads previously authorized credentials from the save file.
 *
 * @return {Promise<OAuth2Client|null>}
 */
async function loadSavedCredentialsIfExist() {
    // try {
        console.log(TOKEN_PATH)
        const content = fs.readFileSync(TOKEN_PATH);
        const credentials = JSON.parse(content);
        return google.auth.fromJSON(credentials);
    // } catch (err) {
    //     console.log("Error:" + err)
    //     return null;
    // }
}

async function authorize() {
    console.log("Start reading...")
    let client = await loadSavedCredentialsIfExist();
    if (client) {
        console.log("Read existing credentials")
        return client;
    }
    client = await authenticate({
        scopes: SCOPES,
        keyfilePath: CREDENTIALS_PATH,
    });
    // if (client.credentials) {
    //     await saveCredentials(client);
    // }
    return client;
}

let requests = [
  {
    replaceAllText: {
      containsText: {
        text: '{{ nome }}',
        matchCase: true,
      },
      replaceText: "Michele",
    },
  },
  {
    replaceAllText: {
      containsText: {
        text: '{{ cognome }}',
        matchCase: true,
      },
      replaceText: "Mastrogiovanni",
    },
  },
  {
    replaceAllText: {
      containsText: {
        text: '{{ data }}',
        matchCase: true,
      },
      replaceText: "07/08/2023",
    },
  },
];

let apiKey = "AIzaSyC1LHa53fdX-idpjSDERcNd0_kZXz8Hpwk"

async function bootstrap() {

    let auth = await authorize()

    console.log(auth)

    google.options({auth: auth});

    google.discoverAPI(`https://docs.googleapis.com/$discovery/rest?version=v1&key=${apiKey}`)
    .then(function(docs) {
      docs.documents.batchUpdate(
          {
            documentId: '1nkwmnhmK1DXGuU_B4Zkq2AERPm48m46qibhBZBJxyJ8',
            resource: {
              requests,
            },
          },
          (err, {data}) => {
            if (err) return console.log('The API returned an error: ' + err);
            console.log(data);
          });
    });

}

bootstrap().catch(console.log)
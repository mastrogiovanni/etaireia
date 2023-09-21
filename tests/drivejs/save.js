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
    try {
        console.log(TOKEN_PATH)
        const content = fs.readFileSync(TOKEN_PATH);
        const credentials = JSON.parse(content);
        return google.auth.fromJSON(credentials);
    } catch (err) {
        console.log("Error:" + err)
        return null;
    }
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

/**
 * Insert new file.
 * @return{obj} file Id
 * */
async function uploadBasic() {
    console.log("Start upload")

    // Get credentials and build service
    // TODO (developer) - Use appropriate auth mechanism for your app
    const auth = await authorize();
    const service = google.drive({ version: 'v3', auth });
    const fileMetadata = {
        name: 'Ciao',
        mimeType: 'application/vnd.google-apps.document',
    };
    const media = {
        mimeType: 'text/html',
        body: fs.createReadStream('index.html'),
    };
    try {
        const file = await service.files.create({
            resource: fileMetadata,
            media: media,
            fields: 'id'
        });
        console.log('File Id:', file.data.id);
        return file.data.id;
    } catch (err) {
        // TODO(developer) - Handle error
        throw err;
    }
}

uploadBasic().catch(console.log)
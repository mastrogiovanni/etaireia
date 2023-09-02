import { generateMnemonic, mnemonicToSeedSync, setDefaultWordlist } from 'bip39';
import { ed25519 } from '@noble/curves/ed25519'

export 	function loadCredentials() {
    let ds = localStorage.getItem('digital.signature');
    if (!ds) {
        return [];
    }
    return JSON.parse(ds);
}

/**
 * Get private key from mnemonic and password
 * 
 * @param mnemonic 
 * @param password 
 * @returns 
 */
function getPrivateKey(mnemonic: string, password: string) : Buffer {
    let seed = mnemonicToSeedSync(mnemonic, password);
    let privateKey = seed.subarray(0, 32)
    return privateKey
}

export function addCredential(data: any, mnemonic: string, publicKey: string) {
    let credentials = loadCredentials()
    credentials.push({
        status: "created",
        data,
        mnemonic,
        publicKey,
    })
    localStorage.setItem('digital.signature', JSON.stringify(credentials));
    return credentials
}

export function deleteCredential(index: number) {
    let credentials = loadCredentials()
    credentials.splice(index, 1)
    localStorage.setItem('digital.signature', JSON.stringify(credentials));
    return credentials
}

export async function requestCredential(mnemonic: string, data: any, password: string) {
    let privateKey = getPrivateKey(mnemonic, password)
    const publicKey = await ed25519.getPublicKey(privateKey);

    // TODO send documents and public key generated 

    return addCredential(data, mnemonic, Buffer.from(publicKey).toString('hex'))
}

/**
 * Recover a credential 
 * 
 * @param mnemonic List of words
 * @param data Data associated to the 
 * @param password 
 */
export async function recoverCredential(mnemonic: string, data: any, password: string) {
    let privateKey = getPrivateKey(mnemonic, password)
    const publicKey = await ed25519.getPublicKey(privateKey);
    return addCredential(data, mnemonic, Buffer.from(publicKey).toString('hex'))
}

/**
 * Get private and public Key starting from a credential and a password
 * 
 * @param index 
 * @param password 
 * @returns 
 */
export async function getKeysFromCredentials(index: number, password: string) {
    let credentials = loadCredentials()
    let credential = credentials[index]
    let privateKey = getPrivateKey(credential.mnemonic, password)
    const publicKey = await ed25519.getPublicKey(privateKey);
    return {
        privateKey: privateKey.toString('hex'),
        publicKey: Buffer.from(publicKey).toString('hex')
    }
}
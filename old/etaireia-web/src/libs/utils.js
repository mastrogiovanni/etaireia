import { credential } from './stores';

// padd with leading 0 if <16
function i2hex(i) {
    return ("0" + i.toString(16)).slice(-2);
}

export function toHex(bytes) {
    return Array.from(bytes).map(i2hex).join("");
}

export function fromHex(string) {
    var bytes = new Uint8Array(Math.ceil(string.length / 2));
    for (var i = 0; i < bytes.length; i++)
        bytes[i] = parseInt(string.substr(i * 2, 2), 16);
    return bytes
}

export function clearCredential() {
    localStorage.removeItem("digital.signature");
    credential.update(null);
}

export function loadCredential() {
    let ds = localStorage.getItem("digital.signature");
    if (!ds) {
        return undefined
    }
    return JSON.parse(ds);
}

export function saveCredential({ privateKey, publicKey, name, surname }) {
    localStorage.setItem(
      "digital.signature",
      JSON.stringify({
        privateKey,
        publicKey,
        name,
        surname,
      })
    );
}
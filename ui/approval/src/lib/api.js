
export async function getSigners() {
    let response = await fetch("/signer/api/v1/credential", {
        method: "GET"
    })
    return await response.json()
}

/**
 * @param {string} hexPublicKey
 */
export async function approveCredential(hexPublicKey) {
    let response = await fetch(`/signer/api/v1/credential/${hexPublicKey}/approved`, {
        method: "PUT",
        body: JSON.stringify({
            hexPublicKey
        })
    })
    return await response.json()
}

/**
 * @param {string} hexPublicKey
 */
export async function rejectCredential(hexPublicKey) {
    let response = await fetch(`/signer/api/v1/credential/${hexPublicKey}/rejected`, {
        method: "PUT",
        body: JSON.stringify({
            hexPublicKey
        })
    })
    return await response.json()
}

/**
 * @param {string} hexPublicKey
 */
export async function deleteCredential(hexPublicKey) {
    let response = await fetch(`/signer/api/v1/credential/${hexPublicKey}/deleted`, {
        method: "PUT",
        body: JSON.stringify({
            hexPublicKey

        })
    })
    return await response.json()
}

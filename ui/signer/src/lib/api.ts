
export async function createCredential(
    data: any, 
    front: File | undefined | null, 
    back: File | undefined | null) {

    const formData = new FormData()
    for (const [key, value] of Object.entries(data)) {
        formData.append(key, value as string)
    }
    if (front) {
        formData.append("front", front)
    }
    if (back) {
        formData.append("back", back)
    }
    let response = await fetch("/signer/api/v1/credential", {
        method: "POST",
        body: formData
    })
    return await response.json()
}



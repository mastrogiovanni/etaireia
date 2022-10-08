<script>
  import UploadImage from "../lib/UploadImage.svelte";
  import * as ed from "@noble/ed25519";
  import { clearCredential, toHex } from "../libs/utils";

  export let credentials;
  let image;
  let message;

  function deleteSignature() {
    clearCredential()
    location.reload();
  }

  async function signAndSend() {

    if (!image) {
      return;
    }

    let data = new Uint8Array(await readFile(image));

    const signature = await ed.sign(data, credentials.privateKey.toUpperCase());

    const formData = new FormData();
    formData.append("document", image);
    formData.append("publicKey", credentials.publicKey);
    formData.append("signature", toHex(signature));

    const response = await fetch("/api/v1/sign", {
      method: "POST",
      body: formData,
    });

    // TODO manage error

    image = undefined;
    message = "Documento inviato firmato con successo"

  }

  function readFile(file) {
    return new Promise((resolve, reject) => {
      // Create file reader
      let reader = new FileReader();
      // Register event listeners
      reader.addEventListener("loadend", (e) => resolve(e.target.result));
      reader.addEventListener("error", reject);
      // Read file
      reader.readAsArrayBuffer(file);
    });
  }
</script>

<form style="padding-top: 10px;">
  <div class="mb-3">
    <span class="badge rounded-pill bg-success"
      >Firma Digitale Riconosciuta</span
    >
    <button class="btn btn-sm btn-primary" on:click={deleteSignature}>Elimina Firma</button>
  </div>
  <div class="mb-3">
    <span class="badge rounded-pill bg-info text-dark"
      >{credentials.publicKey.slice(0, 32)}</span
    >
  </div>
  {#if message}
  <div class="mb-3">
    <span class="badge rounded-pill bg-warning text-dark"
      >{message}</span
    >
  </div>
  {/if}
  <div class="mb-3 form-check">
    <UploadImage bind:image={image} title="Da Firmare" />
  </div>
  <button disabled={!image} type="button" class="btn btn-primary" on:click={signAndSend}>Firma e Invia</button>
</form>

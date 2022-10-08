<script>
  import * as ed from "@noble/ed25519";
  import UploadImage from "../lib/UploadImage.svelte";
  import { saveCredential, toHex } from '../libs/utils';

  export let done = () => {}

  let image;
  let name;
  let surname;

  async function generate() {
    if (!image || !name || !surname) {
      return;
    }
    const privateKey = ed.utils.randomPrivateKey();
    const publicKey = await ed.getPublicKey(privateKey);

    const formData = new FormData();
    formData.append("document", image);
    formData.append("name", name);
    formData.append("surname", surname);
    formData.append("publicKey", toHex(publicKey));

    const response = await fetch("/api/v1/upload", {
      method: "POST",
      body: formData,
    });

    // TODO check errors

    saveCredential({
      privateKey: toHex(privateKey), 
      publicKey: toHex(publicKey), 
      name, 
      surname 
    })

    done();
  }

</script>

<form style="padding-top: 10px;">
  <div class="mb-3">
    <label for="name" class="form-label">Nome</label>
    <input type="text" class="form-control" id="name" bind:value={name} />
  </div>
  <div class="mb-3">
    <label for="surname" class="form-label">Cognome</label>
    <input type="text" class="form-control" id="surname" bind:value={surname} />
  </div>
  <div class="mb-3">
    <UploadImage bind:image={image} title="Documento Identità" />
  </div>
  <button disabled={!image} type="button" class="btn btn-primary" on:click={generate}>Iscrizione</button>
</form>

<script>
  import UploadImage from "../lib/UploadImage.svelte";
  import * as ed from "@noble/ed25519";
  import { clearCredential, toHex } from "../libs/utils";
  import { Col, Container, Row } from "sveltestrap";
  import { credential } from '../libs/stores';
  import { Button, Modal } from "sveltestrap";
  import ShowSignatureModal from "../modals/ShowSignatureModal.svelte";
  import DocumentSentModal from "../modals/DocumentSentModal.svelte";
  import { PAGE_HOME, PAGE_SUBSCRIBE } from "../libs/constants";
  import { page } from '../libs/stores';

  import {
    Card,
    CardBody,
    CardFooter,
    CardHeader,
    CardSubtitle,
    CardText,
    CardTitle,
  } from "sveltestrap";

  let image;
  let message;
  let header;

  function deleteSignature() {
    clearCredential();
    $page = PAGE_SUBSCRIBE;
  }

  async function signAndSend() {
    if (!image) {
      return;
    }

    let data = new Uint8Array(await readFile(image));

    const signature = await ed.sign(data, $credential.privateKey.toUpperCase());

    const formData = new FormData();
    formData.append("document", image);
    formData.append("publicKey", $credential.publicKey);
    formData.append("signature", toHex(signature));

    const response = await fetch("/api/v1/sign", {
      method: "POST",
      body: formData,
    });

    // TODO manage error

    showDocumentSent = true;
    image = undefined;
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

  let showSignature = false
  let showDocumentSent = false;
  
</script>

<ShowSignatureModal bind:isOpen={showSignature} ></ShowSignatureModal>

<DocumentSentModal bind:isOpen={showDocumentSent}></DocumentSentModal>

<Card class="mb-3">
  <CardFooter>
    <Row>
      <Col xs="5" sm="5">La tua firma digitale:</Col>
      <Col xs="5" sm="5">
        <span on:click={() => { showSignature = true }} class="badge rounded-pill bg-info text-dark">
          {$credential.publicKey.slice(0, 8)}
        </span>
      </Col>
      <Col xs="2" sm="2">
        <Button class="delete-signature" on:click={deleteSignature}><i class="bi bi-trash3" /></Button>
      </Col>
    </Row>
  </CardFooter>
</Card>

<Card class="mb-3">
  <CardHeader>
    <CardTitle>Invia Documento</CardTitle>
  </CardHeader>
  <CardBody>
    <CardSubtitle />
    <CardText>
      <UploadImage bind:image title={image ? '' : 'Carica Documento'} />
    </CardText>
    <div style="text-align: center">
      <Button color="primary" disabled={!image} on:click={signAndSend}>Firma e Invia</Button>
    </div>
  </CardBody>
</Card>

<style>
  :global(.delete-signature) {
    background-color: red !important;
  }
</style>

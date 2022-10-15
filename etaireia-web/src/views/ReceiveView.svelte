<script>
  import { onMount } from "svelte";
  import * as QRCode from "qrcode";
  import { loadCredential } from "../libs/utils";
  import { HOST_NAME } from "../libs/constants";

  onMount(() => {
    let credentials = loadCredential();
    let canvas = document.getElementById("canvas");
    QRCode.toCanvas(canvas, HOST_NAME + credentials.publicKey.slice(0, 32), function (error) {
      if (error) console.error(error);
      console.log("success!");
    });
  });
</script>

<h1>Ricevi un documento da firmare</h1>

<p>
    Per ricevere un documento da firmare mostra al mittente il QR code qui sotto: 
</p>

<canvas id="canvas" />

<script>
  import { onMount } from "svelte";
  import svelteLogo from "./assets/svelte.svg";
  import NavBar from "./lib/NavBar.svelte";
  import UploadImage from "./lib/UploadImage.svelte";
  import { loadCredential, toHex } from "./libs/utils";
  import GenerateKeysView from "./views/GenerateKeysView.svelte";
  import SignDocumentView from "./views/SignDocumentView.svelte";

  let credentials;
  
  onMount(() => {
    credentials = loadCredential();
    if (credentials) {
      console.log(credentials)
    }
  });

  function done() {
    credentials = loadCredential();
    if (credentials) {
      console.log(credentials)
    }
  }

</script>

<NavBar></NavBar>

<div class="container">

  {#if !credentials}
  <GenerateKeysView {done}></GenerateKeysView>
  {:else}
  <SignDocumentView {credentials}></SignDocumentView>
  {/if}


  <!--

  <p>
    Check out <a href="https://github.com/sveltejs/kit#readme" target="_blank"
      >SvelteKit</a
    >, the official Svelte app framework powered by Vite!
  </p>

  <p class="read-the-docs">Click on the Vite and Svelte logos to learn more</p>
  -->

</div>

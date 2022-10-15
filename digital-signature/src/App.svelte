<script>
  import { onMount } from "svelte";
  import svelteLogo from "./assets/svelte.svg";
  import NavBar from "./lib/NavBar.svelte";
  import UploadImage from "./lib/UploadImage.svelte";
  import { loadCredential, toHex } from "./libs/utils";
  import GenerateKeysView from "./views/GenerateKeysView.svelte";
  import ReceiveView from "./views/ReceiveView.svelte";
  import SignDocumentView from "./views/SignDocumentView.svelte";
  import { page } from './libs/stores';
  import { credential } from "./libs/stores";
  import {
    PAGE_HOME,
    PAGE_RECEIVE,
    PAGE_SIGNED,
    PAGE_SUBSCRIBE,
  } from "./libs/constants";
  import SignedDocumentsView from "./views/SignedDocumentsView.svelte";
  import Logo from "./lib/Logo.svelte";

  onMount(() => {
    $page = PAGE_HOME;
    $credential = loadCredential();
    if ($credential) {
      $page = PAGE_HOME;
    } else {
      $page = PAGE_SUBSCRIBE;
    }
    // $page = PAGE_RECEIVE
  });

  function done() {
    $credential = loadCredential();
    if ($credential) {
      $page = PAGE_HOME;
    }
  }

</script>

<NavBar />

Ciao

<div class="container">
  {#if $page == PAGE_SUBSCRIBE}
    <GenerateKeysView {done} />
  {:else if $page == PAGE_HOME}
    <SignDocumentView {$credential}></SignDocumentView>
  {:else if $page == PAGE_RECEIVE}
    <ReceiveView />
  {:else if $page == PAGE_SIGNED}
    <SignedDocumentsView />
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

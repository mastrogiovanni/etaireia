<script lang="ts">
	import { onMount } from 'svelte';
	import { setDefaultWordlist } from 'bip39';
	import { browser } from '$app/environment';
	import { Card } from 'flowbite-svelte';
	import { Button } from 'flowbite-svelte';
	import { deleteCredential, loadCredentials } from '$lib/credential';
	import CredentialCard from '../../components/CredentialCard.svelte';

	let credentials: any[] = [];

	function _deleteCredential(index: number) {
		credentials = deleteCredential(index)
	}

	onMount(async () => {
		if (browser) {
			const { Buffer } = await import('buffer');
			window.Buffer = Buffer;
		}
		setDefaultWordlist('italian');
		credentials = loadCredentials();
	});

</script>

<div class="grid place-items-center">

	<Card href="#!" class="w-full">
		<Button color="blue" pill href="credenziali/new" on:click={() => {}}>Crea Nuova Credenziale</Button>
	</Card>

	<Card href="#!" class="w-full">
		<Button color="blue" pill href="credenziali/import" on:click={() => {}}>Importa Credenziale</Button>
	</Card>

	{#if credentials && credentials.length > 0}

		{#each credentials as credential, index}

			<CredentialCard credential={credential} index={index} deleteCredential={_deleteCredential}></CredentialCard>

		{/each}

	{:else}
	
		<p>Non possiedi nessuna credenziale</p>

	{/if}
	
</div>
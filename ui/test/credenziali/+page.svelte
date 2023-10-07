<script lang="ts">
	import { onMount } from 'svelte';
	import { setDefaultWordlist } from 'bip39';
	import { browser } from '$app/environment';
	import { Card } from 'flowbite-svelte';
	import { Badge } from 'flowbite-svelte';
	import { Button } from 'flowbite-svelte';
	// import { Label, Input } from 'flowbite-svelte';
	// import { Modal } from 'flowbite-svelte';
	// import { Drawer, CloseButton, A } from 'flowbite-svelte';
	// import { Icon } from 'flowbite-svelte-icons';
	// import { sineIn } from 'svelte/easing';

	import { deleteCredential, loadCredentials, recoverCredential, requestCredential } from '$lib/credential';

	// let hidden7 = true;
	// let transitionParamsTop = {
	// 	y: -320,
	// 	duration: 200,
	// 	easing: sineIn
	// };

	let credentials: any[] = [];

	// function addCredential(data: any, mnemonic: string, publicKey: string) {
	// 	let current = loadCredentials()
	// 	current.push({
	// 		status: "created",
	// 		data,
	// 		mnemonic,
	// 		publicKey,
	// 	})
	// 	localStorage.setItem('digital.signature', JSON.stringify(current));
	// 	credentials = current;
	// }

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

<!-- Nuova Credenziale -->
<!--
<Drawer placement="top" width="w-full" transitionType="fly" transitionParams={transitionParamsTop} bind:hidden={hidden7} id="sidebar7">

	<div class="flex items-center">
		<div class="container"><p>Metti qui i tuoi dati e fai richiesta</p></div>
		<h5 id="drawer-label" class="inline-flex items-center mb-4 text-base font-semibold text-gray-500 dark:text-gray-400">
			<Icon name="info-circle-solid" class="w-4 h-4 mr-2.5" />Top drawer
		</h5>
	  <CloseButton on:click={() => (hidden7 = true)} class="mb-4 dark:text-white" />
	</div>
	<form>
		<UserData bind:name={name} bind:surname={surname}></UserData>
		<div class="mb-6">
			<Label for="mnemonic" class="block mb-2">Password</Label>
			<Input id="mnemonic" type="password" />
			<Helper class="text-sm mt-2">
				Viene utilizzata per decifrare la tua credenziale, assicurati che sia la stessa usata quando hai generato la credenziale
			</Helper>
		</div>
		<Button color="green" pill on:click={_requestCredential}>Richiedi Credenziale</Button>
	</form>
	<p class="max-w-lg mb-6 text-sm text-gray-500 dark:text-gray-400">
	  Supercharge your hiring by taking advantage of our <A href="/" class="text-primary-600 underline dark:text-primary-500 hover:no-underline">limited-time sale</A> for Flowbite Docs + Job Board. Unlimited access to over 190K top-ranked candidates and the #1 design job board.
	</p>

	<Button color="light" href="/">Learn more</Button>
	<Button href="/" class="px-4">Get access <Icon name="arrow-right-outline" class="w-3.5 h-3.5 ml-2" /></Button>
</Drawer>
-->

<div class="grid place-items-center">

	<Card href="#!" class="w-full">
		<Button color="blue" pill href="credenziali/new" on:click={() => {}}>Crea Nuova Credenziale</Button>
	</Card>

	<Card href="#!" class="w-full">
		<Button color="blue" pill href="credenziali/import" on:click={() => {}}>Importa Credenziale</Button>
	</Card>

	{#if credentials && credentials.length > 0}

	{#each credentials as credential, index}

	<Card href="#!">
		<h5 class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">{credential?.data?.name} {credential?.data?.surname}</h5>
		<p class="font-normal text-gray-700 dark:text-gray-400 leading-tight">
			{#each credential.mnemonic.split(" ") as word}
				<Badge color="green">{word}</Badge>
				&nbsp;
			{/each}
		</p>
		<br/>
		<Button color="red" pill on:click={() => { _deleteCredential(index) }}>Elimina</Button>
	</Card>

	{/each}

	{:else}
	<p>Non possiedi nessuna credenziale</p>

	<!-- <CloseButton on:click={() => (hidden7 = false)} class="mb-4 dark:text-white" /> -->

	{/if}
	
</div>
<script lang="ts">
	import { Label, Input, Helper, Alert, Button, Badge, Modal } from 'flowbite-svelte';
	import UserData from '$lib/../components/UserData.svelte';
	import { generateMnemonic, setDefaultWordlist } from 'bip39';
	import { onMount } from 'svelte';
	import { browser } from '$app/environment';
	import { requestCredential } from '$lib/credential';

	let password: string;
	let name: string;
	let surname: string;
	let email: string;
	let fiscalCode: string;
	let dateOfBirth: Date;
	let placeOfBirth: string;
	let documentType: string;
	let documentId: string;
	let documentExpireTime: string;
	let mnemonic: string;

	let callback: () => void;

	onMount(async () => {
		if (browser) {
			const { Buffer } = await import('buffer');
			window.Buffer = Buffer;
		}
		setDefaultWordlist('italian');
	});

	let defaultModal = false

	function _requestCredential() {
		mnemonic = generateMnemonic();
		console.log(mnemonic);
		callback = async () => {
			let payload = {
				name,
				surname,
				fiscalCode,
				dateOfBirth,
				placeOfBirth,
				documentType,
				documentId,
				documentExpireTime 
			}
			console.log(payload)
			await requestCredential(mnemonic, payload, password)
			
		}
		defaultModal = true
	}

</script>

<Modal title="Nuova Credenziale" bind:open={defaultModal} autoclose>
	<p class="text-base leading-relaxed text-gray-500 dark:text-gray-400">
		Cliccando "Accetta" creerai una credenziale sul tuo cellulare. 
	</p>
	<p class="text-base leading-relaxed text-gray-500 dark:text-gray-400">
		La credenziale sarà protetta dalla tua password e non ci sarà modo per noi di falsificarla.
	</p>
	<p class="text-base leading-relaxed text-gray-500 dark:text-gray-400">
		La tua credenziale potrà essere ripristinata in qualunque momento utilizzando la <b>stessa</b> password e questa serie di 12 parole: scrivile su un pezzo di carta in modo da non perderle più
	</p>
	<p class="text-base leading-relaxed text-gray-500 dark:text-gray-400">
		{#each mnemonic.split(" ") as word}
			<Badge color="green">{word}</Badge>
			&nbsp;
		{/each}
	</p>
	<p class="text-base leading-relaxed text-gray-500 dark:text-gray-400">
		Ricorda: solo la stessa password permetterà di firmare correttamente il documento. Memorizza anche quella.
	</p>
	<svelte:fragment slot="footer">
	  <Button on:click={callback}>Accetto</Button>
	  <Button color="alternative">Annulla</Button>
	</svelte:fragment>
</Modal>

<div class="container mb-3"><p>Metti qui i tuoi dati per richiedere una credenziale</p></div>
<form>
	<UserData 
		bind:name={name} 
		bind:surname={surname}
		bind:email={email}
		bind:fiscalCode={fiscalCode}
		bind:dateOfBirth={dateOfBirth}
		bind:placeOfBirth={placeOfBirth}
		bind:documentType={documentType}
		bind:documentId={documentId}
		bind:documentExpireTime={documentExpireTime}>
	</UserData>

	<div class="mt-6">
		<Alert class="block mb-2 text-sm mt-2">
			<span class="font-medium">Attenzione!</span>
			La password che inserirai qui sotto verrà utilizzata per rendere 
			sicura la tua credenziale. Questa sarà presente solo su questo dispositivo 
			e nemmno noi potremmo leggerla o decifrarla.
			Devi ricordare questa password quindi ti consigliamo di appuntarla in un 
			luogo sicuro.
		</Alert>
	</div>
	
	<div class="mt-6">
		<Label for="mnemonic" class="block mb-2">Password</Label>
		<Input id="mnemonic" type="password" />
		<Helper class="text-sm mt-2">
			Utilizzata per decifrare la tua credenziale
		</Helper>
	</div>
	<Button class="mt-6" color="green" pill on:click={_requestCredential}>Richiedi Credenziale</Button>
</form>


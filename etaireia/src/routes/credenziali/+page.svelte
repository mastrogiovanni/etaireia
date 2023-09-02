<script lang="ts">
	import { onMount } from 'svelte';
	import { generateMnemonic, setDefaultWordlist } from 'bip39';
	import { browser } from '$app/environment';

	import { Tabs, TabItem, Helper } from 'flowbite-svelte';
	import { Card } from 'flowbite-svelte';
	import { Badge } from 'flowbite-svelte';
	import { Button } from 'flowbite-svelte';
	import { Label, Input } from 'flowbite-svelte';
	import { Modal } from 'flowbite-svelte';
	import { Drawer, CloseButton, A } from 'flowbite-svelte';
	import { Icon } from 'flowbite-svelte-icons';
	import { sineIn } from 'svelte/easing';

	import UserData from '$lib/../components/UserData.svelte';
	import { deleteCredential, loadCredentials, recoverCredential, requestCredential } from '$lib/credential';

	let callback: () => void;

	let hidden7 = true;
	let transitionParamsTop = {
		y: -320,
		duration: 200,
		easing: sineIn
	};

	let defaultModal = false;

	let password: string;
	let name: string;
	let surname: string;
	let mnemonic: string;

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

	function _requestCredential() {
		mnemonic = generateMnemonic();
		callback = async () => {
			credentials = await requestCredential(mnemonic, {
				name,
				surname
			}, password)
		}
		defaultModal = true
	}

	async function _recoverCredential() {
		credentials = await recoverCredential(mnemonic, {
			name,
			surname
		}, password)
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

<div class="grid place-items-center">

<Tabs>

	<TabItem open title="Credenziali">

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

		<CloseButton on:click={() => (hidden7 = false)} class="mb-4 dark:text-white" />

		{/if}
	
	</TabItem>

	<TabItem title="Nuova">

		<div class="container"><p>Metti qui i tuoi dati e fai richiesta</p></div>
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
	
	</TabItem>

	<TabItem title="Importa">

		<h1>Importa una credenziale</h1>

		<div class="mb-6">
			<Label for="mnemonic" class="block mb-2">Codice Mnemonico</Label>
			<Input id="mnemonic" placeholder="Mnemonic" />
			<Helper class="text-sm mt-2">
				Serie di parole che identifica una credenziale: (fanfara malgrado nevrosi...)
			</Helper>
		</div>
		<UserData bind:name={name} bind:surname={surname}></UserData>
		<div class="mb-6">
			<Label for="password" class="block mb-2">Password</Label>
			<Input id="password" type="password" />
			<Helper class="text-sm mt-2">
				Viene utilizzata per decifrare la tua credenziale, assicurati che sia la stessa usata quando hai generato la credenziale
			</Helper>
		</div>
		<Button color="green" pill on:click={_recoverCredential}>Importa</Button>
					
	</TabItem>

</Tabs>

</div>
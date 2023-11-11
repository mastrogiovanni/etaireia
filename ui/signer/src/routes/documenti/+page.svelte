<script lang="ts">

	import { onMount } from 'svelte';

	import { Button, TabItem, Tabs, Modal, Label } from 'flowbite-svelte';
	import DrawerSign from '../../components/DrawerSign.svelte';
	import CardToSign from '../../components/CardToSign.svelte';
	import DrawerDocument from '../../components/DrawerDocument.svelte';
	import CardToView from '../../components/CardToView.svelte';
	import { Select } from 'flowbite-svelte';
	import { loadCredentials } from '$lib/credential';
	import { getDocumentsToSign } from '$lib/api';

	// Selected document
	let selectedDocument: any;

	// All documents
	let allDocuments: any[] = []

	// Documents to be signed
	let documentsToSign: any[] = []

	// Credentials that can be used
	let credentials: any[] = []

	let credentialSelected: any;

	let hideSignDrawer = true
	let hideDocumentDrawer = true
	let signedModal = false;

	function show(document: any) {
		selectedDocument = document
		hideSignDrawer = true
		hideDocumentDrawer = false
	}

	function sign(document: any) {
		selectedDocument = document
		hideDocumentDrawer = true
		hideSignDrawer = false
	}

	async function signSent(publicKey: string) {
		signedModal = true
		hideSignDrawer = true
		hideDocumentDrawer = true
	}

	async function changedCredential() {
		console.log(credentialSelected)
		if (!credentialSelected) {
			documentsToSign = []
			return
		}
		let response = await getDocumentsToSign(credentialSelected.fiscalCode, credentialSelected.email)
		if (response?.success) {
			documentsToSign = response.data
			console.log(documentsToSign)
		}
	}
	
	onMount(async () => {
		let result = loadCredentials();
		credentials = result.map((x: { data: { fiscalCode: any; email: any; name: any; }; }) => { return {
			value: {
				fiscalCode: x.data.fiscalCode,
				email: x.data.email
			},
			name: `${x.data.name} (${x.data.fiscalCode})`
		}})

		// documentsToSign = await getDocumentsToSign()
		// resp = await fetch(/json/api/v1/documents")
		// allDocuments = await resp.json()
	})


</script>

<Label>
	Select an option
	<Select on:change={changedCredential} class="mt-2" items={credentials} bind:value={credentialSelected} />
</Label>

{#if selectedDocument}

<DrawerSign 
	bind:hidden={hideSignDrawer} 
	title={selectedDocument.title}
	sign={signSent} />

<DrawerDocument 
	bind:hidden={hideDocumentDrawer}
	documentUrl={selectedDocument.documentUrl}
	title={selectedDocument.title} />

{/if}

<Modal title="Documento Firmato" bind:open={signedModal} autoclose>
	<p class="text-base leading-relaxed text-gray-500 dark:text-gray-400">
		Grazie per aver firmato il documento: "{selectedDocument.title}"
	</p>
	<p class="text-base leading-relaxed text-gray-500 dark:text-gray-400">
		Don Mario saprà...
	</p>
	<svelte:fragment slot="footer">
	  <Button color="alternative" on:click={() => { signedModal = false }}>Chiudi</Button>
	</svelte:fragment>
</Modal>

<div class="grid place-items-center">

{#each documentsToSign as document}
<CardToSign 
	title={document.title}
	description={document.description}
	img={document.img}
	href={document.href} 
	show={() => show(document)}
	sign={() => sign(document)} />
{/each}

</div>

<!--
<div class="grid place-items-center">

<Tabs>

	<TabItem open title="Da Firmare">
		{#each documenti as documento}
			<CardToSign 
				title={documento.title}
				description={documento.description}
				img={documento.img}
				href={documento.href} 
				show={() => show(documento)}
				sign={() => sign(documento)} />
		{/each}
	</TabItem>

	<TabItem title="Tutti">
		
		{#each allDocuments as documento}
		<CardToView 
			title={documento.title}
			description={documento.description}
			img={documento.img}
			href={documento.href} 
			show={() => show(documento)}
		/>
	{/each}

	</TabItem>

</Tabs>

</div>
-->


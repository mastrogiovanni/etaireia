<script lang="ts">

	import { Button, TabItem, Tabs, Modal } from 'flowbite-svelte';
	import { onMount } from 'svelte';
	import DrawerSign from '../../components/DrawerSign.svelte';
	import CardToSign from '../../components/CardToSign.svelte';
	import DrawerDocument from '../../components/DrawerDocument.svelte';
	import CardToView from '../../components/CardToView.svelte';
	import { baseUrl } from '$lib/constants';

	let selectedDocument: any;

	let allDocuments: any[] = []
	let documenti: any[] = []

	let hideSignDrawer = true
	let hideDocumentDrawer = true
	let signedModal = false;

	function show(document: any) {
		console.log("Show", document)
		selectedDocument = document
		hideSignDrawer = true
		hideDocumentDrawer = false
	}

	function sign(document: any) {
		selectedDocument = document
		hideDocumentDrawer = true
		hideSignDrawer = false
	}

	onMount(async () => {
		let resp = await fetch(baseUrl + "/json/api/v1/waiting")
		documenti = await resp.json()

		resp = await fetch(baseUrl + "/json/api/v1/documents")
		allDocuments = await resp.json()
	})

	async function signSent(publicKey: string) {
		console.log("Signed", document)
		signedModal = true
		hideSignDrawer = true
		hideDocumentDrawer = true
	}

</script>

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


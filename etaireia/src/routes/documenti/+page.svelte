<script lang="ts">

	import { Table, TableBody, TableBodyCell, TableBodyRow, TableHead, TableHeadCell, Checkbox, TableSearch, TabItem, Tabs, Modal } from 'flowbite-svelte';
  	import { Card, Button } from 'flowbite-svelte';
  	import { Icon } from 'flowbite-svelte-icons';
	import { Drawer, CloseButton, A } from 'flowbite-svelte';
	import { sineIn } from 'svelte/easing';
	import { Listgroup, ListgroupItem } from 'flowbite-svelte';
	import { onMount } from 'svelte';
	import { loadCredentials } from '$lib/credential';
	import { Label, Select } from 'flowbite-svelte';
	import PdfViewer from 'svelte-pdf';
	import DrawerSign from '../../components/DrawerSign.svelte';
	import CardToSign from '../../components/CardToSign.svelte';
	import DrawerDocument from '../../components/DrawerDocument.svelte';

	let selectedDocument: any;

	let documenti: any[] = [
		{
			title: "Modulo Adesione Cerchi D'Onda",
			description: `Modulo per parte dell'associazione Cerchi D'Onda.`,
			img: "/imgs/iscrizione.jpg",
			documentUrl: "/docs/note.pdf",
			info: {
				name: "Ulteriori Info",
				href: "https://apple.com"
			}
		},
		{
			title: "Modulo Iscrizione SFAF 2024",
			description: `
				Modulo per l'iscrizione alla SFAF: questo modulo è richiesto per far parte dell'associazione
				Cerchi D'Onda e da diritto alla copertura assicurativa per event associativi che verrà
				stipulata.`,
			img: "/imgs/sfaf.png",
			documentUrl: "/docs/note.pdf",
			info: {
				name: "Per saperne di più",
				href: "https://apple.com"
			}
		}
	]

	let hideSignDrawer = true
	let hideDocumentDrawer = true
	let signedModal = false;

	function show(document: any) {
		console.log("Show")
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
		let resp = await fetch("/api/v1/documents")
		documenti = await resp.json()
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
		Mostra qui tutti i documenti firmati e non
	</TabItem>

</Tabs>

</div>


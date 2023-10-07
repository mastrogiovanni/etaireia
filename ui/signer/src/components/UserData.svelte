<script lang="ts">
	import { Select, Label, Input } from 'flowbite-svelte';
	import { onMount } from 'svelte';

	export let name: string;
	export let surname: string;
	export let email: string;
	export let fiscalCode: string;
	export let dateOfBirth: string;
	export let placeOfBirth: string;
	export let documentType: string;
	export let documentId: string;
	export let documentExpireTime: string;

	let documentTypes = [
		{ value: 'CI', name: 'Carta di Identità' },
		{ value: 'PT', name: 'Patente' },
		{ value: 'PS', name: 'Passaporto' }
	];

	function setFakeData() {
		name = "Michele"
		surname = "Mastrogiovanni"
		email = "michele.mastrogiovanni@gmail.com"
		fiscalCode = "MSTMHL80E19D708X"
		dateOfBirth = "1980-05-19"
		placeOfBirth = "Formia"
		documentType = "CI"
		documentId = "AX123456"
		documentExpireTime = "2024-05-19"
	}

	onMount(() => {
		setFakeData()
	})

</script>

<div class="mb-1">
	<Label for="name" class="block mb-2">Nome</Label>
	<Input bind:value={name} id="name" placeholder="Mario" />
</div>

<div class="mb-1">
	<Label for="surname" class="block mb-2">Cognome</Label>
	<Input bind:value={surname} id="surname" placeholder="Rossi" />
</div>

<div class="mb-1">
	<Label for="email" class="block mb-2">Email</Label>
	<Input bind:value={email} id="email" placeholder="mario.rossi@gmail.com" />
</div>

<div class="mb-1">
	<Label for="fiscalCode" class="block mb-2">Codice Fiscale</Label>
	<Input bind:value={fiscalCode} id="fiscalCode" placeholder="RSSMRA80A01H501U" />
</div>

<div class="mb-1">
	<Label for="dateOfBirth" class="block mb-2">Data di Nascita</Label>
	<Input type="date" bind:value={dateOfBirth} id="dateOfBirth" />
</div>

<div class="mb-1">
	<Label for="placeOfBirth" class="block mb-2">Luogo di Nascita</Label>
	<Input bind:value={placeOfBirth} id="placeOfBirth" />
</div>

<div class="mb-1">
	<Label for="documentType" class="block mb-2">Tipo Documento</Label>
	<Select class="mt-2" items={documentTypes} bind:value={documentType} />
</div>

<div class="mb-1">
	<Label for="documentId" class="block mb-2">Numero Documento</Label>
	<Input bind:value={documentId} id="documentId" />
</div>

<div class="mb-1">
	<Label for="documentExpireTime" class="block mb-2">Data Scadenza Documento</Label>
	<Input type="date" bind:value={documentExpireTime} id="documentExpireTime" />
</div>

<script lang="ts">
	import { Label, Input } from 'flowbite-svelte';

    export let name: string;
    export let surname: string;
</script>

<div class="mb-6">
    <Label for="name" class="block mb-2">Nome</Label>
    <Input bind:value={name} id="name" placeholder="Mario" />
</div>

<div class="mb-6">
    <Label for="surname" class="block mb-2">Cognome</Label>
    <Input bind:value={surname} id="surname" placeholder="Rossi" />
</div>

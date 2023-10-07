<script lang="ts">
	import { Card, Button, Badge } from "flowbite-svelte";

    export let index: number;
    export let credential: any;
    export let deleteCredential = (index: number) => {};

</script>

<!-- svelte-ignore missing-declaration -->
<Card href="#!">
    <h5 class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">{credential?.data?.name} {credential?.data?.surname}</h5>
    <p class="font-normal text-gray-700 dark:text-gray-400 leading-tight">
        {#each credential.mnemonic.split(" ") as word}
            <Badge color="green">{word}</Badge>
            &nbsp;
        {/each}
    </p>
    <br/>
    <Button color="red" pill on:click={() => { deleteCredential(index) }}>Elimina</Button>
</Card>

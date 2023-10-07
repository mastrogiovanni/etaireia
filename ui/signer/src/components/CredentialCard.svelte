<script lang="ts">
	import { getCredentialStatus } from "$lib/api";
	import { Card, Button, Badge } from "flowbite-svelte";
	import { onMount } from "svelte";

    export let index: number;
    export let credential: any;
    export let deleteCredential = (index: number) => {};

    let currentStatus = 0

    onMount(async () => {
        let resp = await getCredentialStatus(credential.publicKey)
        console.log(resp)
        if (resp?.success) {
            currentStatus = resp?.data?.status;
        }
    })

    /**
     * @param {number} status
     */
     function getApprovalLabel(status) {
        switch (status) {
            case 0: return "Da Approvare"
            case 1: return "Approvato"
            case 2: return "Rigettato"
            case 3: return "Eliminato"
        }
        return status
    }

    function getApprovalColor(status: number) {
        switch (status) {
            case 0: return "yellow"
            case 1: return "green"
            case 2: return "red"
            case 3: return "dark"
        }
        return "none"
    }

</script>

<!-- svelte-ignore missing-declaration -->
<Card href="#!">
    <h5 class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">{credential?.data?.name} {credential?.data?.surname}</h5>
    <p class="font-normal text-gray-700 dark:text-gray-400 leading-tight">
        {#each credential.mnemonic.split(" ") as word, index}
            <Badge color="green">{word}</Badge>
            &nbsp;
        {/each}
    </p>
    <p class="mt-4">
        Stato: <Badge large color={getApprovalColor(currentStatus)}>{getApprovalLabel(currentStatus)}</Badge>
    </p>
    <br/>
    <Button color="red" pill on:click={() => { deleteCredential(index) }}>Elimina</Button>
</Card>

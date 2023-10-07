<script>
    import {
        Badge,
        Button,
        Modal,
        Table,
        TableBody,
        TableBodyCell,
        TableBodyRow,
        TableHead,
        TableHeadCell,
        TableSearch,
    } from "flowbite-svelte";
    import {
        Navbar,
        NavBrand,
        NavLi,
        NavUl,
        NavHamburger,
    } from "flowbite-svelte";
    import { onMount } from "svelte";
    import DocumentApproval from "../components/DocumentApproval.svelte";
    import { deleteCredential, getSigners } from "$lib/api";

    let searchTerm = "";

    /**
     * @type {any[]}
     */
    let items = []

    $: filteredItems = items.filter(
        (item) =>
            item.name.toLowerCase().indexOf(searchTerm.toLowerCase()) !== -1
    );

    let selectedItem = items[0]
    
    let defaultModal = false

    /**
     * @param {any} item
     */
    function open(item) {
        selectedItem = item
        defaultModal = true
    }

    async function reloadData() {
        let result = await getSigners()
        if (result?.success) {
            items = result?.data
            items.sort((a, b) => {
                if (a.id < b.id) {
                    return -1
                }
                if (a.id > b.id) {
                    return 1
                }
                return 0
            })
        }
    }

    onMount(reloadData)

    /**
     * @param {number} status
     */
    function getApprovalColor(status) {
        switch (status) {
            case 0: return "yellow"
            case 1: return "green"
            case 2: return "red"
            case 3: return "dark"
        }
        return "none"
    }

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

    async function approvedOrRejected() {
        await reloadData()
    }

    async function _deleteItem() {
        let resp = await deleteCredential(itemToDelete.hexPublicKey)
        console.log(resp);
        itemToDelete = undefined
        confirmModalOpen = false
        await reloadData();
    }

    /**
     * @type {{ hexPublicKey: any; } | undefined}
     */
    let itemToDelete;

    /**
     * @param {{ hexPublicKey: string; }} item
     */
     async function _delete(item) {
        itemToDelete = item;
        confirmModalOpen = true
    }

    let confirmModalOpen = false;

</script>

{#if selectedItem}
    <DocumentApproval command={approvedOrRejected} personToValidate={selectedItem} bind:open={defaultModal} />
{/if}

<Modal title="Conferma Eliminazione" bind:open={confirmModalOpen} autoclose size="lg">

    <p class="text-base leading-relaxed text-gray-500 dark:text-gray-400">
        Sei sicuro di voler eliminare questa credenziale?<br/>
        L'utente collegato non potrà più firmare documenti con questa credenziale.
    </p>

    <div class="m-12 mb-2 mt-2">
    </div>

    <svelte:fragment slot="footer">
        <Button on:click={_deleteItem}>Approva Eliminazione</Button>
        <Button color="alternative">Cancella</Button>
    </svelte:fragment>

</Modal>
  
<Navbar let:hidden let:toggle>
    <NavBrand href="/approval">
        <span
            class="self-center whitespace-nowrap text-xl font-semibold dark:text-white"
            >Cerchi D'Onda</span
        >
    </NavBrand>
    <NavHamburger on:click={toggle} />
    <NavUl {hidden}>
        <NavLi href="/approval">Home</NavLi>
    </NavUl>
</Navbar>

<div class="container mx-auto mt-16">
    <TableSearch
        placeholder="Nome o cognome"
        hoverable={true}
        bind:inputValue={searchTerm}
    >
    <caption class="p-5 text-lg font-semibold text-left text-gray-900 bg-white dark:text-white dark:bg-gray-800">
    Documenti da approvare
    <p class="mt-1 text-sm font-normal text-gray-500 dark:text-gray-400">
        Seleziona una richiesta da approvare e verifica i dati immessi
    </p>
  </caption>
        <TableHead>
            <TableHeadCell>ID</TableHeadCell>
            <TableHeadCell>Nome</TableHeadCell>
            <TableHeadCell></TableHeadCell>
            <TableHeadCell></TableHeadCell>
        </TableHead>
        <TableBody tableBodyClass="divide-y">
            {#each filteredItems as item}
                <TableBodyRow>
                    <TableBodyCell>{item?.id}</TableBodyCell>
                    <TableBodyCell>{item?.name}</TableBodyCell>
                    <TableBodyCell><Badge color={getApprovalColor(item.status)}>{getApprovalLabel(item?.status)}</Badge></TableBodyCell>
                    <TableBodyCell>
                        {#if item?.status === 0}
                            <a href="#!" class="font-medium text-primary-600 hover:underline dark:text-primary-500" on:click={() => open(item)}>Edit</a>
                        {:else}
                            {#if item?.status === 1}
                                <Button on:click={() => { _delete(item) }}>Elimina</Button>
                            {/if}
                        {/if}
                      </TableBodyCell>
                </TableBodyRow>
            {/each}
        </TableBody>
    </TableSearch>
</div>

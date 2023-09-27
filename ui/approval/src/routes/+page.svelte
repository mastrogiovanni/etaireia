<script>
    import {
        Table,
        TableBody,
        TableBodyCell,
        TableBodyRow,
        TableHead,
        TableHeadCell,
        TableSearch,
    } from "flowbite-svelte";
    import { Button, Modal } from 'flowbite-svelte';
    import { Alert } from "flowbite-svelte";
    import { Gallery } from 'flowbite-svelte';
    import { Card } from 'flowbite-svelte';
    import { Checkbox } from 'flowbite-svelte';
    import {
        Navbar,
        NavBrand,
        NavLi,
        NavUl,
        NavHamburger,
    } from "flowbite-svelte";
    import { onMount } from "svelte";
    import DocumentApproval from "../components/DocumentApproval.svelte";
    let searchTerm = "";

    const michele_imgs = [
        { alt: 'michele-fronte', src: '/imgs/michele/fronte.png' },
        { alt: 'michele-retro', src: '/imgs/michele/retro.png' },
    ]

    const micaela_imgs = [
        { alt: 'micaela-fronte', src: '/imgs/micaela/fronte.png' },
        { alt: 'micaela-retro', src: '/imgs/micaela/retro.png' },
    ]

    const tommaso_imgs = [
        { alt: 'tommaso-fronte', src: '/imgs/tommaso/fronte.jpeg' },
    ]

    let items = [
        { id: 1, name: "Michele", surname: "Mastrogiovanni", fiscalCode: "MSTMHL80E19D708X", birthPlace: "Formia (LT)", birthDate: "19/05/1980", cardEmission: "16/06/20", cardExpire: "19/05/2031", images: michele_imgs },
        { id: 2, name: "Micaela", surname: "Spinosa", fiscalCode: "MSTMHL80E19D708X", birthPlace: "Gaeta (LT)", birthDate: "19/05/1980", cardEmission: "16/06/20", cardExpire: "19/05/2031", images: micaela_imgs },
        { id: 3, name: "Tommaso", surname: "Spinosa", fiscalCode: "MSTMHL80E19D708X", birthPlace: "Gaeta (LT)", birthDate: "19/05/1980", cardEmission: "16/06/20", cardExpire: "19/05/2031", images: tommaso_imgs },
    ];

    $: filteredItems = items.filter(
        (item) =>
            item.name.toLowerCase().indexOf(searchTerm.toLowerCase()) !== -1
    );

    let selectedItem = items[0]
    
    let defaultModal = false

    function open(item) {
        console.log("test")
        selectedItem = item
        defaultModal = true
    }

</script>

<DocumentApproval personToValidate={selectedItem} bind:open={defaultModal} />

<Navbar let:hidden let:toggle>
    <NavBrand href="/">
        <!--
        <img
            src="/images/flowbite-svelte-icon-logo.svg"
            class="mr-3 h-6 sm:h-9"
            alt="Flowbite Logo"
        />
        -->
        <span
            class="self-center whitespace-nowrap text-xl font-semibold dark:text-white"
            >Cerchi D'Onda</span
        >
    </NavBrand>
    <NavHamburger on:click={toggle} />
    <NavUl {hidden}>
        <NavLi href="/">Home</NavLi>
        <!--
        <NavLi href="/about">About</NavLi>
        <NavLi href="/docs/components/navbar">Navbar</NavLi>
        <NavLi href="/pricing">Pricing</NavLi>
        <NavLi href="/contact">Contact</NavLi>
        -->
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
        </TableHead>
        <TableBody class="divide-y">
            {#each filteredItems as item}
                <TableBodyRow>
                    <TableBodyCell>{item.id}</TableBodyCell>
                    <TableBodyCell>{item.name} {item.surname}</TableBodyCell>
                    <TableBodyCell>
                        <a href="#!" class="font-medium text-primary-600 hover:underline dark:text-primary-500" on:click={() => open(item)}>Edit</a>
                      </TableBodyCell>
                </TableBodyRow>
            {/each}
        </TableBody>
    </TableSearch>
</div>

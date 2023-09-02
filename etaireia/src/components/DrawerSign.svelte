<script lang="ts">
	import { loadCredentials } from "$lib/credential";
	import { Button, CloseButton, Drawer, Label, Select } from "flowbite-svelte";
	import { Icon } from "flowbite-svelte-icons";
	import { onMount } from "svelte";
	import { sineIn } from 'svelte/easing';

    export let hidden = true
    export let title = ""
	export let sign = async (publicKey: string) => {}

	let transitionParamsTop = {
		y: -320,
		duration: 200,
		easing: sineIn
	};

    let credenziali: any[] = []

    let credenzialiSelect : any[] = []

    $: {
        credenzialiSelect = credenziali.map(credenziale => {
            return {
                value: credenziale.publicKey,
                name: credenziale.data.name + " " + credenziale.data.surname
            }
        })
    }

    let selected: string = "";

    async function _sign() {
        hidden = true
		await sign(selected)
    }

    onMount(() => {
        credenziali = loadCredentials()
    }) 

</script>

<Drawer placement="top" width="w-full" transitionType="fly" transitionParams={transitionParamsTop} bind:hidden={hidden} id="sidebar7">
	<div class="flex items-center">
		<div class="container"><p>{title}</p></div>
		<h5 id="drawer-label" class="inline-flex items-center mb-4 text-base font-semibold text-gray-500 dark:text-gray-400">
			<Icon name="info-circle-solid" class="w-4 h-4 mr-2.5" />Firma
		</h5>
	  <CloseButton on:click={() => (hidden = true)} class="mb-4 dark:text-white" />
	</div>
	<p class="max-w-lg mb-6 text-sm text-gray-500 dark:text-gray-400">
	  Cliccando su Firma, apponi la tua firma digitale e invii il documento all'associazione 
	</p>
	<Label class="mb-4">
		Seleziona una credenziale
		<Select class="mt-2" items={credenzialiSelect} bind:value={selected} />
	</Label>
	<Button on:click={() => (hidden = true)} color="light" href="/documenti">Annulla</Button>
	<Button on:click={_sign} class="px-4">Firma <Icon name="arrow-right-outline" class="w-3.5 h-3.5 ml-2" /></Button>
</Drawer>

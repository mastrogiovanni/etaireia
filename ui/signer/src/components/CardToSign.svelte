<script lang="ts">
	import { Card, Button } from "flowbite-svelte";
	import { Icon } from "flowbite-svelte-icons";

    export let title = ""
    export let description = ""
    export let img = ""
    export let href = ""

    export let show = () => {}
    export let sign = () => {}

</script>

<Card img={img} class="mb-4">
	<h5 class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">
        {title}
	</h5>
	<p class="mb-3 font-normal text-gray-700 dark:text-gray-400 leading-tight">
		{description}
	</p>
	<div class="mb-3 flex mt-4 space-x-3 lg:mt-6">
		<Button on:click={show} color="blue">Visualizza</Button>
		<Button on:click={sign} color="green" class="dark:text-white">Firma</Button>
	</div>
	{#if href}
	<a href={href} class="inline-flex items-center text-primary-600 hover:underline">
		Informazioni di iscrizione sul sito
		<Icon name="arrow-up-right-from-square-outline" class="w-3 h-3 ml-2.5" />
	</a>
	{/if}
</Card>

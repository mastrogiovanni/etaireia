<script lang="ts">
	import { Gallery } from 'flowbite-svelte';
	import { onMount } from 'svelte';

    let contacts: any[] = []

    onMount(async () => {
        let resp = await fetch("/json/api/v1/contacts")
		contacts = await resp.json()
    })
</script>

<div class="container mx-auto px-2">
    <Gallery class="gap-4 grid-cols-2">
    {#each contacts as image}
        <figure class="max-w-lg">
            <img src="{image.src}" alt="{image.alt}" class="h-auto space-x-4 rounded-lg" />
            <figcaption class="mt-2 text-sm text-center text-gray-500 dark:text-gray-400"><span class="text-gray-900">{image.alt}</span><br/>{image.ruolo}</figcaption>
        </figure>
    {/each}
    </Gallery>
</div>



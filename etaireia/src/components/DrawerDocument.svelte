<script>
	import { Drawer, CloseButton } from "flowbite-svelte";
	import PdfViewer from "svelte-pdf";
	import { sineIn } from 'svelte/easing';

    export let hidden = true
    export let title = ""
    export let documentUrl = ""

	let transitionParamsTop = {
		y: -320,
		duration: 200,
		easing: sineIn
	};

</script>

<Drawer placement="top" width="w-full" transitionType="fly" transitionParams={transitionParamsTop} bind:hidden={hidden} id="sidebar7">
	<div class="flex items-center">
		<div class="container"><p>{title}</p></div>
	  <CloseButton on:click={() => (hidden = true)} class="mb-4 dark:text-white" />
	</div>
	<PdfViewer data={null} url={documentUrl} showButtons={["navigation", "zoom", "download", "autoflip", "timeInfo", "pageInfo"]} />
</Drawer>

<script lang="ts">
	import '../app.postcss';

	import { Card, BottomNav, BottomNavItem, Listgroup, Indicator } from 'flowbite-svelte';
	import { Icon } from 'flowbite-svelte-icons';
	import { ButtonGroup, Button, GradientButton } from 'flowbite-svelte';
	import { Tabs, TabItem } from 'flowbite-svelte';
	import { Navbar, NavBrand, NavLi, NavUl, NavHamburger } from 'flowbite-svelte';
	import { page } from '$app/stores';
	import { Badge } from 'flowbite-svelte';
	import { onMount } from 'svelte';
	import { baseUrl } from '$lib/constants';

	$: activeUrl = $page.url.pathname;
	let activeClass = 'text-white bg-green-700 md:bg-transparent md:text-green-700 md:dark:text-white dark:bg-green-600 md:dark:bg-transparent';
  	let nonActiveClass = 'text-gray-700 hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-green-700 dark:text-gray-400 md:dark:hover:text-white dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent';

	let company: {name: string} = { name: "-" }

	onMount(async () => {
		let resp = await fetch(baseUrl + "/api/v1/company")
		company = await resp.json()
	})

</script>

<Navbar let:hidden let:toggle>
	<NavBrand href="/">
		<span class="self-center whitespace-nowrap text-xl font-semibold dark:text-white">{company?.name}</span>
	</NavBrand>
	<NavHamburger on:click={toggle} />
	<NavUl {activeUrl} {hidden} {activeClass} {nonActiveClass}>
		<NavLi on:click={toggle} href="/signer" actve={true}>
			Home
		</NavLi>
		<NavLi on:click={toggle} href="/signer/credenziali">Credenziali</NavLi>
		<NavLi on:click={toggle} href="/signer/documenti">
			Documenti
			&nbsp;
			<Indicator color="blue" border size="xl" placement="right" class="text-xs font-bold">18</Indicator>
		</NavLi>
		<NavLi on:click={toggle} href="/signer/whoweare">Chi siamo</NavLi>
	</NavUl>
</Navbar>

<slot></slot>


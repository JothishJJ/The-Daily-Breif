<script>
	import { onMount } from 'svelte';

	/** @type {any | undefined} */
	let dataObj;
	onMount(async () => {
		const response = await fetch('http://localhost:8080');
		const data = await response.json();
		dataObj = data;
	});
</script>

<svelte:head>
	<title>World | The Daily Brief</title>
</svelte:head>

<h1>World</h1>
{#if dataObj}
	<div>{dataObj.message}</div>
{/if}

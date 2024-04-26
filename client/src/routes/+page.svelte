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

<h1>Hello World</h1>
{#if dataObj}
	<div>{dataObj.message}</div>
{/if}

<script>
	export let name;
	let title = "no title";
	let gameName = "";

	function handleClick() {
		title = gameName;
	}
</script>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>

<svelte:head>
	<title>{title}</title>
</svelte:head>

<main>
	<h1>Hello {name}!</h1>
	<p>
		Visit the
		<a href="https://svelte.dev/tutorial">Svelte tutorial</a>
		to learn how to build Svelte apps.
	</p>
	<input
		id="gameName"
		type="text"
		data-qa="game-name"
		bind:value={gameName} />
	<input type="button" data-qa="start-game" on:click={handleClick} />
</main>

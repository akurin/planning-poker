<script>
    import { onMount } from "svelte";

    export let id;

    let title = "not loaded";

    onMount(async () => {
        const res = await fetch(`/api/games/${id}`);
        const game = await res.json();
        title = game.title;
    });
</script>

<svelte:head>
    <title>{title}</title>
</svelte:head>

<main>
    <p>Game</p>
    <p>{title}</p>
</main>

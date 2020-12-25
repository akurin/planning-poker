<script>
    import {onMount} from "svelte";

    export let id;

    let game = null;
    let pageTitle = game === null ? "Loading..." : game.title;

    onMount(async () => {
        const res = await fetch(`/api/games/${id}`);
        game = await res.json();
        pageTitle = game.title;
    });
</script>

<svelte:head>
    <title>{pageTitle}</title>
</svelte:head>

<main data-qa="game">
    <p>Game</p>
    {#if game !== null}
        <p data-qa="game-title">{game.title}</p>
    {/if}
</main>

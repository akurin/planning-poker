<script>
    // import Link from "svelte-routing/src/Link.svelte";
    import { onMount } from "svelte";

    async function handleClick() {
        const res = await fetch(`http://localhost:8080/api/games`, {
            method: "POST",
            body: JSON.stringify({
                title: "some",
            }),
        });

        const responseBody = await res.json();
        window.location.href = `http://localhost:8080/games/${responseBody.id}`;
    }
</script>

<main>
    <p>New game</p>
    <input type="text" data-qa="game-name" />
    <button on:click={handleClick}>New Game</button>
</main>

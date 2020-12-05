<script>
    import { navigate } from "svelte-routing";

    let title = "";

    async function handleClick() {
        const res = await fetch(`/api/games`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                title: title,
            }),
        });

        const responseBody = await res.json();
        navigate(`/games/${responseBody.id}`);
    }
</script>

<main>
    <p>New game</p>
    <input type="text" data-qa="game-name" bind:value={title} />
    <button on:click={handleClick}>New Game</button>
</main>

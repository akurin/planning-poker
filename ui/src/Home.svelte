<script>
    import {navigate} from "svelte-routing";

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
    <input bind:value={title} data-qa="game-name" type="text"/>
    <button data-qa="start-game" on:click={handleClick}>New Game</button>
</main>

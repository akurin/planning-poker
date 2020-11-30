<script>
    import Link from "svelte-routing/src/Link.svelte";
    import { onMount } from "svelte";

    onMount(async () => {
        const res = await fetch(`http://localhost:5000/api/`, {
            method: "POST",
            body: JSON.stringify({
                
            }),
        });
        const photos = await res.json();
        console.log(photos);
    });
</script>

<main>
    <p>New game</p>
    <Link to="games/123"><input type="button" data-qa="start-game" /></Link>
    <input type="text" data-qa="game-name" />
</main>

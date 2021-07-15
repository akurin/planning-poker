<script>
    import {onMount} from "svelte";

    export let id;

    let playerName = "";

    let game = null;
    let pageTitle = game === null ? "Loading..." : game.title;

    let player = null;

    onMount(async () => {
        const res = await fetch(`/api/games/${encodeURIComponent(id)}`);
        game = await res.json();
        pageTitle = game.title;
    });


    async function handleOkClick() {
        const res = await fetch(`/api/games/${encodeURIComponent(id)}/players`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                name: playerName,
            }),
        });

        player = await res.json();
    }
</script>

<style>
    .overlay {
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        position: fixed;
        background-color: rgba(52, 74, 94, 0.8);
        z-index: 999;
        overflow-x: hidden;
        overflow-y: auto;
        outline: none;
    }

    .dialog {
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        position: fixed;
        z-index: 1000;
        width: 100%;
        height: 100%;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    .dialog__content {
        position: relative;
        max-width: 520px;
        margin-bottom: 5%;
        background-color: white;
        padding: 30px;
        color: #344a5e;
        border-radius: 5px;
        box-shadow: 0 30px 30px 0 rgba(52, 74, 94, 0.8);
    }

    .dialog h2 {
        margin: 0;
        margin-bottom: 10px;
    }

    .dialog__ok-btn {
        cursor: pointer;
        height: 30px;
        padding: 6px 15px;
        margin-right: 15px;
        min-width: 60px;
        transition-duration: .25s;
        transition-property: background-color, color;
        text-align: center;
        border-width: 1px;
        border-style: solid;
        border-radius: 3px;
        box-shadow: 0 1px 2px 0 rgba(64, 61, 4, 0.44);
        font-size: 14px;
        color: white;
        border-color: transparent;
        background-color: #0279c0;
    }

    .dialog__cancel-btn {
        cursor: pointer;
        color: #0279c0;
        background-color: transparent;
        border-color: transparent;
        box-shadow: none;
        font-family: -apple-system, BlinkMacSystemFont, Ubuntu, 'Segoe UI', Roboto, Oxygen, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;

    }

    .dialog__cancel-btn:hover {
        text-decoration: underline;
    }

    .dialog__close-btn {
        border: 0;
        background: none;
        position: absolute;
        top: 20px;
        right: 20px;
        width: 15px;
        height: 15px;
        cursor: pointer;
    }

    .dialog__close-btn::before, .dialog__close-btn::after {
        position: absolute;
        top: -2px;
        right: 6px;
        display: block;
        width: 2px;
        height: 19px;
        content: '';
        border-radius: 3px;
        background-color: #d9d9d9;
    }

    .dialog__close-btn::before {
        transform: rotate(45deg);
    }

    .dialog__close-btn::after {
        transform: rotate(-45deg);
    }
</style>

<svelte:head>
    <title>{pageTitle}</title>
</svelte:head>

// loading game info -> showing popup -> game entered
<main data-qa="game">
    <p>Game</p>
    {#if game !== null}
        <p data-qa="game-title">{game.title}</p>
    {/if}

    {#if player === null}
        <div class="dialog">
            <div class="dialog__content">
                <header>
                    <h2>Choose your display name</h2>
                </header>
                <main>
                    <label for="name">Your display name</label>
                    <input id="name" bind:value={playerName} data-qa="game-name" type="text"/>
                </main>
                <footer>
                    <button class="dialog__ok-btn" on:click={handleOkClick}>OK</button>
                </footer>
            </div>
        </div>
        <div class="overlay"></div>
    {/if}
</main>

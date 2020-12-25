<script>
    import {Route, Router} from "svelte-routing";
    import Game from "./Game.svelte";
    import Home from "./Home.svelte";

    export let url = "";
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

<!-- <svelte:head>
	<title>{title}</title>
</svelte:head> -->

<main>
    <Router {url}>
        <div>
            <Route let:params path="games/:id">
                <Game id={params.id}/>
            </Route>

            <Route component={Home} path="/"/>
        </div>
    </Router>
</main>

<script>
    import { onMount } from "svelte";
    import { querystring } from "svelte-spa-router";
    import AnimeCard from "../components/AnimeCard.svelte";

    let animes = [];
    let loading = false;
    let title = "Browse Anime";
    let activeFilter = "";

    const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ".split("");
    const genres = [
        "Action",
        "Adventure",
        "Comedy",
        "Drama",
        "Fantasy",
        "Horror",
        "Mystery",
        "Romance",
        "Sci-Fi",
        "Slice of Life",
        "Sports",
        "Supernatural",
    ];

    // React to querystring changes
    $: {
        const params = new URLSearchParams($querystring);
        const q = params.get("q");
        // If q is present, fetch data. providing a default fallback for initial load if needed
        if (q) {
            fetchData(q);
            activeFilter = q;
            title = `Browse: ${q}`;
        } else {
            // Default load (e.g. Action) or clear
            activeFilter = "";
            title = "Browse Anime";
            animes = [];
        }
    }

    async function fetchData(query) {
        loading = true;
        animes = [];
        try {
            // Using a limit of 50 as planned
            const url = `http://localhost:8080/api/anime/search?q=${encodeURIComponent(query)}&limit=50`;

            const res = await fetch(url);
            if (!res.ok) throw new Error("Failed to fetch");
            const data = await res.json();
            animes = data.data.map((item) => item.node);
        } catch (err) {
            console.error(err);
        } finally {
            loading = false;
        }
    }

    function handleFilterClick(filter) {
        // Navigate to same page but update query param
        // This will trigger the reactive statement above
        window.location.hash = `#/browse?q=${encodeURIComponent(filter)}`;
    }
</script>

<div class="container mx-auto px-4 py-8 font-body">
    <h1
        class="text-4xl font-heading font-black mb-8 text-center text-transparent bg-clip-text bg-gradient-to-r from-primary to-accent"
    >
        BROWSE ANIME
    </h1>

    <!-- A-Z Filter -->
    <div class="mb-8">
        <h3
            class="text-xl text-white font-heading mb-4 border-l-4 border-primary pl-3"
        >
            By Letter
        </h3>
        <div class="flex flex-wrap gap-2 justify-center">
            {#each alphabet as letter}
                <button
                    on:click={() => handleFilterClick(letter)}
                    class="w-10 h-10 rounded-lg flex items-center justify-center font-bold transition duration-300 border border-white/10
          {activeFilter === letter
                        ? 'bg-primary text-white shadow-[0_0_15px_#a855f7]'
                        : 'bg-dark/50 text-gray-400 hover:bg-white/10 hover:text-white'}"
                >
                    {letter}
                </button>
            {/each}
        </div>
    </div>

    <!-- Genre Filter -->
    <div class="mb-12">
        <h3
            class="text-xl text-white font-heading mb-4 border-l-4 border-accent pl-3"
        >
            By Genre
        </h3>
        <div class="flex flex-wrap gap-3 justify-center">
            {#each genres as genre}
                <button
                    on:click={() => handleFilterClick(genre)}
                    class="px-4 py-2 rounded-full text-sm font-bold uppercase tracking-wider transition duration-300 border border-white/10
          {activeFilter === genre
                        ? 'bg-accent text-white shadow-[0_0_15px_#f472b6]'
                        : 'bg-dark/50 text-gray-400 hover:bg-white/10 hover:text-white'}"
                >
                    {genre}
                </button>
            {/each}
        </div>
    </div>

    <!-- Results -->
    <div class="min-h-[400px]">
        <h2
            class="text-2xl font-bold mb-6 text-white tracking-widest drop-shadow-lg flex items-center gap-3"
        >
            {#if activeFilter}
                <span
                    class="w-2 h-8 bg-gradient-to-b from-primary to-accent rounded-full"
                ></span>
                Results for "{activeFilter}"
            {/if}
        </h2>

        {#if loading}
            <div class="flex justify-center items-center h-64">
                <div
                    class="animate-spin rounded-full h-16 w-16 border-t-4 border-b-4 border-primary shadow-[0_0_20px_#a855f7]"
                ></div>
            </div>
        {:else if animes.length === 0 && activeFilter}
            <div
                class="text-center p-12 bg-dark/50 backdrop-blur rounded-2xl border border-white/5"
            >
                <p class="text-gray-400 text-xl font-heading">
                    No anime found for "{activeFilter}".
                </p>
            </div>
        {:else if animes.length === 0 && !activeFilter}
            <div
                class="text-center p-12 bg-dark/50 backdrop-blur rounded-2xl border border-white/5"
            >
                <p class="text-gray-400 text-xl font-heading">
                    Select a letter or genre to start browsing.
                </p>
            </div>
        {:else}
            <div
                class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-6"
            >
                {#each animes as anime (anime.id)}
                    <AnimeCard {anime} />
                {/each}
            </div>
        {/if}
    </div>
</div>

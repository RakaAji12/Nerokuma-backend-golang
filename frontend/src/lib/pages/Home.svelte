<script>
    import { onMount } from "svelte";
    import { querystring } from "svelte-spa-router";
    import AnimeCard from "../components/AnimeCard.svelte";
    import SearchBar from "../components/SearchBar.svelte";

    let animes = [];
    let loading = false;
    let title = "Trending Anime";
    let activeTab = "all";

    const tabs = [
        { id: "all", label: "Top Anime" },
        { id: "airing", label: "Top Airing" },
        { id: "upcoming", label: "Top Upcoming" },
        { id: "movie", label: "Top Movies" },
    ];

    // Reactive statement to fetch data when querystring changes
    $: {
        const params = new URLSearchParams($querystring);
        const q = params.get("q");
        fetchData(q, activeTab);
    }

    async function fetchData(query, type) {
        loading = true;
        animes = [];
        try {
            let url = `http://localhost:8080/api/anime/trending?type=${type}`;
            if (query) {
                // If searching, ignore tabs
                url = `http://localhost:8080/api/anime/search?q=${encodeURIComponent(query)}`;
                title = `Search Results for "${query}"`;
            } else {
                title =
                    tabs.find((t) => t.id === type)?.label || "Trending Anime";
            }

            const res = await fetch(url);
            if (!res.ok) throw new Error("Failed to fetch");
            const data = await res.json();
            // The API returns { data: [ { node: ... } ] }
            animes = data.data.map((item) => item.node);
        } catch (err) {
            console.error(err);
        } finally {
            loading = false;
        }
    }

    function switchTab(tabId) {
        activeTab = tabId;
        // Reset search (remove query param) if switching tabs?
        // Or just re-fetch. Ideally clear search if browsing top lists.
        // For simplicity, we just refetch, assuming no search query active or we override it.
        // But the reactive block depends on querystring.
        // Let's just call fetchData directly if not searching.
        const params = new URLSearchParams($querystring);
        if (!params.get("q")) {
            fetchData(null, activeTab);
        }
    }
</script>

<div class="container mx-auto px-4 py-8">
    <div class="mb-12 text-center">
        <h1
            class="text-5xl md:text-7xl font-heading font-black mb-8 text-transparent bg-clip-text bg-gradient-to-r from-primary via-accent to-primary animate-pulse drop-shadow-[0_0_20px_rgba(168,85,247,0.5)]"
        >
            NEROKUMA
        </h1>
        <SearchBar />
    </div>

    <div class="flex items-center gap-4 mb-8 justify-between flex-wrap">
        <div class="flex items-center gap-4">
            <div
                class="h-10 w-2 bg-primary rounded-full shadow-[0_0_15px_#a855f7]"
            ></div>
            <h2
                class="text-3xl font-heading font-bold text-white tracking-widest drop-shadow-lg"
            >
                {title}
            </h2>
        </div>

        <!-- Tabs -->
        <div
            class="flex gap-2 bg-dark/50 backdrop-blur rounded-lg p-1 border border-white/5 overflow-x-auto"
        >
            {#each tabs as tab}
                <button
                    on:click={() => switchTab(tab.id)}
                    class="px-4 py-2 rounded-md text-xs sm:text-sm font-bold uppercase tracking-wider transition duration-300 whitespace-nowrap
                {activeTab === tab.id
                        ? 'bg-primary text-white shadow-md'
                        : 'text-gray-400 hover:text-white hover:bg-white/5'}"
                >
                    {tab.label}
                </button>
            {/each}
        </div>
    </div>

    {#if loading}
        <div class="flex justify-center items-center h-64">
            <div
                class="animate-spin rounded-full h-16 w-16 border-t-4 border-b-4 border-primary shadow-[0_0_20px_#a855f7]"
            ></div>
        </div>
    {:else if animes.length === 0}
        <div
            class="text-center p-12 bg-dark/50 backdrop-blur rounded-2xl border border-white/5"
        >
            <p class="text-gray-400 text-xl font-heading">
                No anime found matching your criteria.
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

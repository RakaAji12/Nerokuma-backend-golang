<script>
    import { onMount } from "svelte";
    import AnimeCard from "../components/AnimeCard.svelte";

    let animes = [];
    let loading = false;

    // Calculate current season
    const date = new Date();
    const month = date.getMonth(); // 0-11
    const currentYear = date.getFullYear();
    let currentSeason = "";

    if (month >= 0 && month <= 2) currentSeason = "winter";
    else if (month >= 3 && month <= 5) currentSeason = "spring";
    else if (month >= 6 && month <= 8) currentSeason = "summer";
    else currentSeason = "fall";

    let selectedYear = currentYear;
    let selectedSeason = currentSeason;

    const seasons = ["winter", "spring", "summer", "fall"];

    $: {
        fetchSeasonal(selectedYear, selectedSeason);
    }

    async function fetchSeasonal(year, season) {
        loading = true;
        animes = [];
        try {
            const url = `http://localhost:8080/api/anime/season/${year}/${season}`;
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

    function capitalize(s) {
        return s.charAt(0).toUpperCase() + s.slice(1);
    }

    function prevSeason() {
        let idx = seasons.indexOf(selectedSeason);
        if (idx === 0) {
            selectedSeason = "fall";
            selectedYear--;
        } else {
            selectedSeason = seasons[idx - 1];
        }
    }

    function nextSeason() {
        let idx = seasons.indexOf(selectedSeason);
        if (idx === 3) {
            selectedSeason = "winter";
            selectedYear++;
        } else {
            selectedSeason = seasons[idx + 1];
        }
    }
</script>

<div class="container mx-auto px-4 py-8 font-body">
    <h1
        class="text-4xl font-heading font-black mb-8 text-center text-transparent bg-clip-text bg-gradient-to-r from-primary to-accent drop-shadow-[0_0_10px_rgba(168,85,247,0.5)]"
    >
        SEASONAL ANIME
    </h1>

    <div class="flex justify-center items-center gap-6 mb-12">
        <button
            on:click={prevSeason}
            class="p-2 rounded-full bg-dark/50 hover:bg-primary/20 hover:text-primary transition border border-white/10"
            aria-label="Previous Season"
        >
            <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-6 w-6"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
                ><path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M15 19l-7-7 7-7"
                /></svg
            >
        </button>

        <div class="text-center">
            <h2 class="text-3xl font-bold text-white uppercase tracking-widest">
                {selectedSeason}
                {selectedYear}
            </h2>
        </div>

        <button
            on:click={nextSeason}
            class="p-2 rounded-full bg-dark/50 hover:bg-primary/20 hover:text-primary transition border border-white/10"
            aria-label="Next Season"
        >
            <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-6 w-6"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
                ><path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M9 5l7 7-7 7"
                /></svg
            >
        </button>
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
                No anime found for this season.
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

<script>
    import { onMount } from "svelte";
    import { pop } from "svelte-spa-router";

    export let params = {};

    let anime = null;
    let loading = true;
    let error = null;

    onMount(async () => {
        try {
            const res = await fetch(
                `http://localhost:8080/api/anime/${params.id}`,
            );
            if (!res.ok) throw new Error("Failed to fetch anime details");
            anime = await res.json();
        } catch (err) {
            error = err.message;
        } finally {
            loading = false;
        }
    });
</script>

<div class="container mx-auto px-4 py-8 font-body">
    <button
        on:click={() => pop()}
        class="mb-8 flex items-center gap-3 px-6 py-3 bg-dark/50 backdrop-blur hover:bg-white/10 border border-white/5 rounded-full text-primary font-heading tracking-wider transition duration-300 group hover:border-primary/50"
    >
        <span class="transform group-hover:-translate-x-1 transition"
            >&larr;</span
        > Back to Search
    </button>

    {#if loading}
        <div class="flex justify-center items-center h-64">
            <div
                class="animate-spin rounded-full h-16 w-16 border-t-4 border-b-4 border-primary shadow-[0_0_20px_#a855f7]"
            ></div>
        </div>
    {:else if error}
        <div
            class="bg-red-900/50 backdrop-blur border border-red-500/50 text-red-200 px-6 py-4 rounded-xl shadow-[0_0_20px_rgba(239,68,68,0.3)]"
            role="alert"
        >
            <strong class="font-bold">Error!</strong>
            <span class="block sm:inline">{error}</span>
        </div>
    {:else if anime}
        <div
            class="bg-dark/70 backdrop-blur-md rounded-3xl overflow-hidden shadow-[0_0_50px_-10px_rgba(0,0,0,0.5)] border border-white/5"
        >
            <div class="md:flex">
                <div class="md:shrink-0 md:w-1/3 relative group">
                    <img
                        class="h-full w-full object-cover transition duration-700 group-hover:scale-105"
                        src={anime.image_url ||
                            "https://via.placeholder.com/300x450?text=No+Image"}
                        alt={anime.title}
                    />
                    <div
                        class="absolute inset-0 bg-gradient-to-t from-darker via-transparent to-transparent opacity-60"
                    ></div>
                </div>
                <div class="p-8 md:p-12 md:w-2/3 flex flex-col justify-center">
                    <div
                        class="uppercase tracking-widest text-sm font-heading text-accent mb-2"
                    >
                        Score: <span class="text-white text-lg font-bold"
                            >{anime.mean ? anime.mean.toFixed(2) : "N/A"}</span
                        >
                    </div>

                    <h1
                        class="text-4xl md:text-5xl font-heading font-black tracking-tight text-white mb-6 drop-shadow-md"
                    >
                        {anime.title}
                    </h1>

                    <div class="flex flex-wrap gap-2 mb-8">
                        {#each anime.genres as genre}
                            <span
                                class="px-4 py-1.5 rounded-md text-xs font-bold uppercase tracking-widest bg-white/5 text-gray-300 border border-white/10 hover:border-primary/50 hover:text-white transition cursor-default"
                            >
                                {genre.name}
                            </span>
                        {/each}
                    </div>

                    <div
                        class="prose prose-invert prose-lg max-w-none text-gray-300 leading-relaxed bg-black/20 p-6 rounded-xl border border-white/5"
                    >
                        {anime.synopsis}
                    </div>
                </div>
            </div>
        </div>
    {/if}
</div>

<script>
    import { link } from "svelte-spa-router";
    export let anime;
</script>

<div
    class="group relative bg-dark/60 backdrop-blur-sm rounded-xl overflow-hidden border border-white/5 hover:border-primary/50 transition duration-500 transform hover:-translate-y-2 hover:shadow-[0_10px_30px_-10px_rgba(168,85,247,0.3)]"
>
    <!-- Image Container -->
    <div class="relative overflow-hidden aspect-[3/4]">
        <img
            src={anime.image_url ||
                "https://via.placeholder.com/225x320?text=No+Image"}
            alt={anime.title}
            class="w-full h-full object-cover transition duration-700 group-hover:scale-110"
        />
        <!-- Overlay Gradient -->
        <div
            class="absolute inset-0 bg-gradient-to-t from-darker via-transparent to-transparent opacity-80"
        ></div>

        <!-- Score Badge -->
        <div
            class="absolute top-2 right-2 bg-black/70 backdrop-blur text-white text-xs font-bold px-3 py-1 rounded-full border border-primary/30 flex items-center gap-1"
        >
            <span class="text-accent">★</span>
            {anime.mean ? anime.mean.toFixed(2) : "N/A"}
        </div>
    </div>

    <!-- Content -->
    <div class="p-4 relative">
        <h3
            class="text-lg font-bold text-white truncate group-hover:text-primary transition duration-300 font-heading"
        >
            <a use:link href={`/anime/${anime.id}`}>{anime.title}</a>
        </h3>
        <div
            class="mt-3 flex flex-wrap gap-1.5 h-16 overflow-hidden content-start"
        >
            {#each anime.genres.slice(0, 3) as genre}
                <span
                    class="text-[10px] uppercase tracking-wider bg-white/5 hover:bg-white/10 text-gray-300 px-2 py-1 rounded border border-white/5 transition"
                >
                    {genre.name}
                </span>
            {/each}
        </div>
    </div>
</div>

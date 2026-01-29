<script>
    import { push } from "svelte-spa-router";

    let query = "";

    function handleSearch() {
        if (query.trim()) {
            push(`/?q=${encodeURIComponent(query)}`);
        }
    }

    function handleKeydown(event) {
        if (event.key === "Enter") {
            handleSearch();
        }
    }
</script>

<div class="relative group w-full max-w-xl mx-auto">
    <div
        class="absolute -inset-1 bg-gradient-to-r from-primary to-accent rounded-full blur opacity-40 group-hover:opacity-100 transition duration-500"
    ></div>
    <div
        class="relative flex items-center bg-dark/80 backdrop-blur-md rounded-full border border-white/10 overflow-hidden shadow-2xl"
    >
        <input
            type="text"
            placeholder="Search Anime..."
            bind:value={query}
            on:keydown={handleKeydown}
            class="pl-6 pr-4 py-4 w-full bg-transparent text-white placeholder-gray-500 focus:outline-none font-body text-lg"
        />
        <button
            on:click={handleSearch}
            class="pr-6 pl-4 py-4 text-primary hover:text-accent transition duration-300 transform group-hover:scale-110"
            aria-label="Search"
        >
            <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-6 w-6"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
            >
                <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
                />
            </svg>
        </button>
    </div>
</div>

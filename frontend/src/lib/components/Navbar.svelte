<script>
    import { link } from "svelte-spa-router";
</script>

<nav
    class="sticky top-0 z-40 bg-darker/70 backdrop-blur-lg border-b border-white/5 shadow-lg"
>
    <div class="container mx-auto px-6 py-4 flex justify-between items-center">
        <a
            use:link
            href="/"
            class="text-3xl font-heading font-black bg-clip-text text-transparent bg-gradient-to-r from-primary to-accent hover:opacity-80 transition duration-300 drop-shadow-[0_0_10px_rgba(168,85,247,0.5)]"
        >
            NEROKUMA
        </a>
        <div class="flex space-x-6">
            <a
                use:link
                href="/"
                class="text-sm uppercase tracking-widest text-gray-300 hover:text-primary transition duration-300 font-heading"
                >Home</a
            >
            <a
                use:link
                href="/browse"
                class="text-sm uppercase tracking-widest text-gray-300 hover:text-primary transition duration-300 font-heading"
                >Browse</a
            >
            <a
                use:link
                href="/seasonal"
                class="text-sm uppercase tracking-widest text-gray-300 hover:text-primary transition duration-300 font-heading"
                >Seasonal</a
            >
        </div>
    </div>
</nav>

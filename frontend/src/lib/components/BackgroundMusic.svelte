<script>
    import { onMount } from "svelte";
    export let src;

    let audio;
    let isPlaying = false;
    let volume = 0.3;

    function togglePlay() {
        if (audio) {
            if (isPlaying) {
                audio.pause();
            } else {
                audio
                    .play()
                    .catch((e) => console.error("Autoplay prevented:", e));
            }
            isPlaying = !isPlaying;
        }
    }

    onMount(() => {
        if (audio) {
            audio.volume = volume;
            // Try autoplay
            const playPromise = audio.play();
            if (playPromise !== undefined) {
                playPromise
                    .then((_) => {
                        isPlaying = true;
                    })
                    .catch((error) => {
                        console.log(
                            "Autoplay was prevented, user interaction required.",
                        );
                        isPlaying = false;
                    });
            }
        }
    });
</script>

<div class="fixed bottom-4 left-4 z-50">
    <audio bind:this={audio} {src} loop></audio>

    <button
        on:click={togglePlay}
        class="flex items-center justify-center w-12 h-12 rounded-full bg-black/60 border border-primary/50 text-primary hover:bg-primary hover:text-white transition-all duration-300 shadow-[0_0_15px_rgba(168,85,247,0.4)] backdrop-blur-md"
        aria-label={isPlaying ? "Pause Music" : "Play Music"}
    >
        {#if isPlaying}
            <!-- Pause Icon -->
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
                    d="M10 9v6m4-6v6m7-3a9 9 0 11-18 0 9 9 0 0118 0z"
                />
            </svg>
        {:else}
            <!-- Play Icon -->
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
                    d="M14.752 11.168l-3.197-2.132A1 1 0 0010 9.87v4.263a1 1 0 001.555.832l3.197-2.132a1 1 0 000-1.664z"
                />
                <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
                />
            </svg>
        {/if}
    </button>
</div>

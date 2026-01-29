<script>
    export let src;
</script>

<div class="fixed inset-0 w-full h-full overflow-hidden -z-50 bg-black">
    <video
        autoplay
        muted
        loop
        playsinline
        class="absolute top-0 left-0 w-full h-full object-cover opacity-50 contrast-125 saturate-150"
    >
        <source {src} type="video/mp4" />
        Your browser does not support the video tag.
    </video>
    <div
        class="absolute inset-0 bg-darker/60 bg-gradient-to-t from-darker via-transparent to-darker"
    ></div>
</div>

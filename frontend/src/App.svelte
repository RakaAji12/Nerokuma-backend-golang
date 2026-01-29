<script>
  import Router from "svelte-spa-router";
  import Navbar from "./lib/components/Navbar.svelte";
  import Home from "./lib/pages/Home.svelte";
  import Browse from "./lib/pages/Browse.svelte";
  import Seasonal from "./lib/pages/Seasonal.svelte";
  import AnimeDetail from "./lib/pages/AnimeDetail.svelte";
  import BackgroundVideo from "./lib/components/BackgroundVideo.svelte";
  import BackgroundMusic from "./lib/components/BackgroundMusic.svelte";

  const routes = {
    "/": Home,
    "/browse": Browse,
    "/seasonal": Seasonal,
    "/anime/:id": AnimeDetail,
  };
</script>

<div
  class="min-h-screen text-gray-100 font-body antialiased selection:bg-accent selection:text-white"
>
  <BackgroundVideo src="/assets/video/background.mp4" />
  <BackgroundMusic src="/assets/audio/music.mp3" />

  <div class="relative z-10 flex flex-col min-h-screen">
    <Navbar />

    <main class="flex-grow container mx-auto px-4 py-6">
      <Router {routes} />
    </main>

    <footer
      class="text-center py-6 text-gray-400 text-sm font-heading tracking-widest bg-darker/80 backdrop-blur-md border-t border-white/5"
    >
      <p>&copy; 2026 NEROKUMA. CRAFTED WITH SVELTE & TAILWIND.</p>
    </footer>
  </div>
</div>

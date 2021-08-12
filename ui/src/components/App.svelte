<script>
  import Home from "./Home.svelte";
  import About from "./About.svelte";
  import page from "page";

  let currentPage;

  page("/", () => (currentPage = Home));
  page("/about", () => (currentPage = About));
  page("*", "/");
  page();

  $: homeActive = currentPage === Home;
  $: aboutActive = currentPage === About;
</script>

<main class="container text-center">
  <ul class="nav nav-pills my-2">
    <li class="nav-item">
      <a class="nav-link" class:active={homeActive} href="/">HOME</a>
    </li>
    <li class="nav-item">
      <a class="nav-link" class:active={aboutActive} href="/about">ABOUT</a>
    </li>
  </ul>
  <svelte:component this={currentPage} />
</main>

<script>
  import page from "page";
  import Home from "./Home.svelte";
  import Chat from "./Chat.svelte";

  let currentPage;

  page("/", () => (currentPage = Home));
  page("/chat", () => (currentPage = Chat));
  page("*", "/");
  page();

  $: homeActive = currentPage === Home;
  $: chatActive = currentPage === Chat;
</script>

<main class="container vh-100 d-flex flex-column">
  <ul class="nav nav-pills my-2">
    <li class="nav-item">
      <a class="nav-link" class:active={homeActive} href="/">HOME</a>
    </li>
    <li class="nav-item">
      <a class="nav-link" class:active={chatActive} href="/chat">CHAT</a>
    </li>
  </ul>
  <svelte:component this={currentPage} />
</main>

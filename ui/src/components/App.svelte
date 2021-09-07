<script>
  import page from "page";
  import chatStore from "../stores/chatStore";
  import Home from "./Home.svelte";
  import Chat from "./Chat.svelte";

  let currentPage;

  page("/", () => (currentPage = Home));
  page("/chat", () => (currentPage = Chat));
  page("*", "/");
  page();

  chatStore.connect();

  $: homeActive = currentPage === Home;
  $: chatActive = currentPage === Chat;
</script>

<main class="container text-center vh-100 d-flex flex-column">
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

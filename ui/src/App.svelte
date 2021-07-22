<script>
  import Counter from "./Counter.svelte";
  import Message from "./Message.svelte";
</script>

<main>
  <h1>Go Embed Demo</h1>
  <Counter />
  <Message />
</main>

<style>
  :root {
    font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen,
      Ubuntu, Cantarell, "Open Sans", "Helvetica Neue", sans-serif;
  }

  main {
    text-align: center;
    margin: 0 auto;
    padding: 1em;
  }

  h1 {
    color: #ff3e00;
    text-transform: uppercase;
    font-size: 4rem;
    font-weight: 100;
    line-height: 1.1;
    margin: 2rem auto;
  }
</style>

import { component$, Slot } from "@builder.io/qwik";
import Header from "../components/header/header";

export default component$(() => {
  return (
    <div class={`whole`}>
      <main>
        <Header />
        <section
          style={{
            minHeight: "500px",
          }}
        >
          <Slot />
        </section>
      </main>
      <footer>
        <a href="https://github.com/hominsu" target="_blank">
          Made By HominSu
        </a>
      </footer>
    </div>
  );
});

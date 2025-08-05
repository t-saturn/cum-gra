import { component$ } from "@builder.io/qwik";
import type { DocumentHead } from "@builder.io/qwik-city";

export default component$(() => {
  return (
    <>
      <div>sso frontend </div>
    </>
  );
});

export const head: DocumentHead = {
  title: "sso auth client",
  meta: [
    {
      name: "description",
      content: "Qwik site description",
    },
  ],
};

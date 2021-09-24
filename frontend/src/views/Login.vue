<template>
  <div class="Login">
    <div>
      <div class="mb-2">
        <h1 class="text-white text-3xl font-semibold mb-1">Sign in</h1>
        <span class="text-gray-500 text-base"
          >Don't have an account?
          <a href="/" class="Link">Sign up here</a></span
        >
      </div>
      <form>
        <input type="email" class="LoginInput" placeholder="Enter your email" />
        <input
          type="email"
          class="LoginInput"
          placeholder="Enter your password"
        />
        <button class="PrimaryButton">
          Login
        </button>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  name: "Login",
};
</script>

<style style="scss">
.Login {
  @apply w-screen h-screen flex flex-col text-center justify-center items-center;
}

.LoginInput {
  @apply w-full p-5 rounded-lg outline-none focus:outline-none text-gray-400 placeholder-gray-400 text-xs my-3;
  background-color: #1c1f27;
  transition: all 400ms ease-in-out;
}

.LoginInput:focus {
  background-color: #171a20;
  filter: drop-shadow(0 0 0.5rem #3d4457);
}

.PrimaryButton {
  @apply text-white font-bold py-3 rounded-full w-full mt-3 outline-none focus:outline-none;
  background-color: #80c44c;
  transition: all 400ms ease-in-out;
}

.PrimaryButton:hover {
  background-color: #629739;
}

.Link {
  @apply text-gray-400 underline;
  text-underline-offset: 0.3em;
}
</style>

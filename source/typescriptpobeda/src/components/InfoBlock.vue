<script setup lang="ts">

const props = defineProps({
  title: String,
  mainText: String,
});

</script>

<template>
  <div class="card flex border border-slate-400 lg:max-w-[400px]  rounded-3xl p-5 gap-4 hover:-translate-y-2 transition">
    <slot class="icon" />
    <div class="flex flex-col gap-2">
      <h1 class="main-text text-2xl">{{ props.title }}</h1>
      <p>{{ props.mainText }}</p>
    </div>
  </div>
</template>

<style scoped>
  .card{

  }

  @media (max-width: 800px) {
    .card{
      max-width: calc(100vw - 64px);
      height: 120px;
    }
  }

  @media(max-width: 400px) {
    .card{
      max-width: 100vw;
      height: 140px;
    }
  }
</style>
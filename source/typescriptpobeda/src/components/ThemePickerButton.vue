<script setup lang="ts">
import {MoonIcon , Sun} from "lucide-vue-next";
import {ref, watch} from "vue";

const emit = defineEmits<{
  (event: 'themeChange'): void;
}>();

const dark = ref(true);

watch(dark, ()=>{
  console.log('wathing');
  if (!dark.value) {
    document.getElementById('btn')?.classList.add('light')
    return
  }
  document.getElementById('btn')?.classList.remove('light')
});


const handleClick = () =>{
  emit('themeChange');
  dark.value = !dark.value;
  console.log('dark.value:', dark.value);

}

</script>

<template>
  <button @click="handleClick" id="btn" class="absolute flex items-center justify-center lg:ml-8 sm:ml-4 sm:mt-4 lg:mt-8 mt-4 ml-8 size-12 rounded-2xl border border-white/50">
    <MoonIcon v-if="dark"/>
    <Sun v-else />
  </button>
</template>

<style scoped>
#btn{
  background-color: black;
  color: white
}

  #btn.light{
    background-color: white;
    color: black;
    border-color: black;
  }
</style>
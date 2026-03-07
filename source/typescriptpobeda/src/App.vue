<script lang="ts" setup>
import {ref, watch} from "vue";
import NumberFlow from '@number-flow/vue'
import InfoBlock from "./components/InfoBlock.vue";
import { Zap, Brain, Code, Bug, Shield, AlertCircle } from "lucide-vue-next";
import MemesCarousel from "@/components/MemesCarousel.vue";
import ThemePickerButton from "@/components/ThemePickerButton.vue";


const infoBlocks = [

  {
    title: "Типизация",
    mainText: "Прекрасная и всеми любимая статическая типизация",
    icon: Brain,
  },
  {
    title: "Баги JS",
    mainText: "undefined is not a function...",
    icon: Bug,
  },
  {
    title: "Безопасность",
    mainText: "TS ловит ошибки до запуска, а не в проде",
    icon: Shield,
  },
  {
    title: "Код читабельный",
    mainText: "TS-код = сразу понятно, что происходит",
    icon: Code,
  },
  {
    title: "ХЗ что писать",
    mainText: "",
    icon: AlertCircle,
  },
  {
    title: "Blazingly fast",
    mainText: "And memory effective",
    icon: Zap,
  },
];

const gayTheme = ref(false);
const tsAge = ref(0);
const darkTheme = ref(true);

function calculateTypescriptAge(): number {
  const creationDate = new Date(2012, 9, 1);
  const currentDate = new Date();
  const diffMs = currentDate.getTime() - creationDate.getTime();
  const msPerYear = 1000 * 60 * 60 * 24 * 365.25;

  const years = diffMs / msPerYear;
  return Number(years.toFixed(10));
}


watch(darkTheme, ()=>{
  console.log('wathing');
  if (!darkTheme.value) {
    document.getElementById('app-container')?.classList.add('light')
    return
  }
  document.getElementById('app-container')?.classList.remove('light')
});

watch(gayTheme, ()=>{
  if (gayTheme.value) {
    document.getElementById('tswins')?.classList.add('gay')
    return
  }
  document.getElementById('tswins')?.classList.remove('gay')

})

setInterval(() => {
  tsAge.value = calculateTypescriptAge();
},50);
</script>


<template>
  <div id="app-container"  class="app-container">
    <ThemePickerButton @theme-change="darkTheme = !darkTheme"/>
    <button
        @click="gayTheme = !gayTheme"
        class="fixed bottom-4 left-4 text-sm p-2 rounded-2xl border border-slate-400 bg-gray-800/10 backdrop-blur-md z-20 transition"
    >
      {{ !gayTheme ? 'Enable' : 'Disable' }} furry theme
    </button>
    <div class="flex sm:p-16 p-4 mt-20 flex-col items-center align-center gap-10">
      <h1 id="tswins" class="tswins text-center font-bold text-5xl">TS победа уже</h1>
      <div>
        <NumberFlow
            class="sm:text-3xl lg:text-6xl text-4xl"
            :value="tsAge"
            :format="{ minimumFractionDigits: 8, maximumFractionDigits: 8 }"
        />
        <span class="text-3xl opacity-45">лет</span>
      </div>
      <div class="grid mt-10 md:grid-cols-2 lg:grid-cols-3 sm:grid-cols-1 gap-4">
        <InfoBlock
            v-for="(block, index) in infoBlocks"
            :key="index"
            :title="block.title"
            :mainText="block.mainText"
        >
          <component :is="block.icon" />
        </InfoBlock>

      </div>

      <MemesCarousel class=""/>
      <a href="https://t.me/laynexx" class="underline pb-10">Предложить свой мем</a>

    </div>
  </div>

</template>

<style scoped>

.tswins.gay{
  animation: colorFlow 3s infinite alternate;
}

.app-container.light{
  background: #fff;
  color: black
}


@keyframes colorFlow {
  0% { color: rgb(255, 70, 70); }
  20% { color: #ffbb3d; }
  40% { color: #ffff4f; }
  60% { color: #3fff3f; }
  80% { color: #5454ff; }
  100% { color: purple; }
}

</style>
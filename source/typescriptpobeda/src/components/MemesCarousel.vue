<script setup lang="ts">
import {
  Carousel,
  CarouselContent,
  CarouselItem,
  CarouselNext,
  CarouselPrevious,
} from '@/components/ui/carousel'
import { ref } from "vue";

const images = ref(
    Array.from({ length: 7 }, (_, i) => `/meme-images/img_${i}.png`)
);
</script>

<template>
  <Carousel class="relative w-full">
    <CarouselContent>
      <CarouselItem
          v-for="(item, index) in images"
          :key="index"
          class="flex items-center justify-center transition-transform duration-500 ease-in-out"
      >
        <img
            :src="item"
            class="w-full h-auto sm:max-w-[320px] md:max-w-[480px] lg:max-w-[600px] xl:max-w-[600px]"
            alt="item image"
        />
      </CarouselItem>
    </CarouselContent>

    <CarouselPrevious
        class="absolute opacity-50 left-4 top-1/2 transform -translate-y-1/2 bg-black p-2 rounded-full shadow-md hover:bg-gray-200 transition"
    />
    <CarouselNext
        class="absolute opacity-50 right-4 top-1/2 transform -translate-y-1/2 bg-black  p-2 rounded-full shadow-md hover:bg-gray-200 transition"
    />
  </Carousel>
</template>

<style scoped>

.carousel-content {
  overflow: hidden;
}

.carousel-item img {
  object-fit: cover;
}

.carousel-previous,
.carousel-next {
  z-index: 10;
}

.carousel-next:hover,
.carousel-previous:hover {
  background-color: #f3f4f6;
}

</style>

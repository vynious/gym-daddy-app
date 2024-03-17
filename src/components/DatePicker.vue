<template>
    <a-space direction="vertical" :size="12">
      <a-date-picker v-model:value="value1" />
    </a-space>
  </template>
  <script setup>
  import { ref } from 'vue';
  const value1 = ref();
  </script>
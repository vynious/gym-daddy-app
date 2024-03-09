<template>
    <a-result
      status="error"
      title="Error Encountered"
      sub-title="Sorry for the inconveniece. Please try again."
    >
      <template #extra>
        <a-button key="console" type="primary">Logout</a-button>
        <a-button key="buy">Try Again</a-button>
      </template>

    </a-result>
  </template>
  <script setup>
//   export {};
  </script>
  <style scoped>
  .desc p {
    margin-bottom: 1em;
  }
  </style>
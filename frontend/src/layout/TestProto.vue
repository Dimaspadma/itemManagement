<script setup lang="ts">

import axios from 'axios';
import { decodeResponse, formatTimestamp } from '../lib/ResponseService'
import Response from '../lib/response'
import { onMounted, ref } from 'vue';

const response = ref<Response.awesomeProject.Response>()

onMounted(async () => {
  try {
    const result = await axios.get("http://localhost/users/", {
      responseType: "arraybuffer",
    })
    response.value = decodeResponse(result.data)
    
    response.value.data.forEach(user => {
      console.log(formatTimestamp(user.createdAt))
    });
  } catch (error) {
    console.error("Failed to fetch user data:", error);
  }
})

</script>

<template>

</template>
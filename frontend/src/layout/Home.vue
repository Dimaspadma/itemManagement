<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { fetchPaketData, type Paket } from '../api/paketData';
import StackedList from '../components/StackedList.vue';

const pakets = ref<Paket[]>()
const loading = ref(false)

onMounted(async () => {
  loading.value = true
  try {
    const res = await fetchPaketData()
    pakets.value = res
  } catch (error) {
    console.log("Error bro: " + error)
  } finally {
    loading.value = false
  }
})

</script>
  
<template>
  <div>
    <div v-if="loading">Loading...</div>

    <div v-if="pakets">
       <StackedList :pakets="pakets" />
    </div>
  </div>
</template>
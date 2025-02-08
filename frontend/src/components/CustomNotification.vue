<script setup lang="ts">
import { ref } from 'vue';
import Toast from './Toast.vue';

interface Notification {
  message: string,
  type: "success" | "danger" | "warning",
}
const notifications = ref<Notification[]>([])

// Fungsi untuk nambah notifikasi baru
const addNotification = (message: string, type: Notification["type"]) => {
  console.log("add notification")
  notifications.value.push({ message, type });
  
  // Auto-remove setelah 5 detik
  setTimeout(() => {
    notifications.value.shift();
  }, 5000);
};

// Fungsi untuk hapus notifikasi yang diklik
const removeNotification = (index: number) => {
  notifications.value.splice(index, 1);
};

// Expose `addNotification` biar bisa dipake di tempat lain
defineExpose({ addNotification });

</script>

<template>
  <div class="fixed bottom-0 right-0">
    <Toast v-for="(notification, index) in notifications" :key="index" 
    :notification="notification" :index="index" :remove-func="removeNotification" />
  </div>
</template>
<script setup lang="ts">

import axios from 'axios';
import { create, decodeResponse, formatTimestamp, formatTimestamp2 } from '../lib/ResponseService'
import Response from '../lib/response'
import { onBeforeUnmount, onMounted, ref } from 'vue';
import CustomNotification from '../components/CustomNotification.vue'
import { Metadata, User } from '../response';
import { Timestamp } from '../google/protobuf/timestamp';

const response = ref<Response.awesomeProject.Response>()

const notificationRef = ref<InstanceType<typeof CustomNotification> | null>(null)

const message = ref("")
let eventSource: EventSource

const fetchUsers = async () => {
  console.log("Fetch Users")
  try {
    const result = await axios.get("http://localhost/users/", {
      responseType: "arraybuffer",
    })
    response.value = decodeResponse(result.data)
  } catch (error) {
    console.error("Failed to fetch user data:", error)
  }
}

const addUser = async () => {
  let user: User = {
    id: -1,
    username: "jhonchina",
    createdAt: Timestamp.now()
  }
  let bytes = User.toBinary(user)
  console.log(formatTimestamp2(user.createdAt))

  try {
    const response = await axios.post("http://localhost/users/test", bytes, {
      headers: {
        "Content-Type": "application/protobuf",
      }
    })
    console.log(response.data)
  } catch (error){
    console.error(error)
  }
}

const connectToSSE = () => {
  console.log("Connect to SSE")
  // Menghubungkan ke server SSE
  eventSource = new EventSource('http://localhost/events/')

  eventSource.onopen = () => {
    message.value = "Open EventSource"
  }

  // Menerima data setiap kali server mengirimkan pesan
  eventSource.onmessage = (event) => {
    message.value = event.data
    const data = JSON.parse(message.value)
    if (data.key === "modify"){
      notificationRef.value?.addNotification(data.action, "success")
      fetchUsers()
    }
  }

  // Handle error jika koneksi terputus
  eventSource.onerror = (err) => {
    console.log("Error occurred:", err)
    eventSource.close()
  };
}

onMounted(() => {
  connectToSSE()
  fetchUsers()

  addUser()
})

onBeforeUnmount(() => {
  if (eventSource) {
    eventSource.close()
  }
})

</script>

<template>
  <CustomNotification ref="notificationRef" />
  <h1>Event: {{ message }}</h1>
  <ul v-if="response">
    <li v-for="(user, index) in response.data" :key="index">
      <p>{{ user.id }} - {{ user.username }}</p>
    </li>
  </ul>
</template>
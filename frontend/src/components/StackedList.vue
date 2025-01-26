<script setup lang="ts">

import type { Paket } from '../api/paketData';

defineProps<{ pakets: Paket[] }>()

</script>

<template>
  <ul role="list" class="divide-y divide-gray-100">
    <li v-for="paket in pakets" :key="paket.name" class="flex justify-between gap-x-6 py-5">
      <div class="flex min-w-0 gap-x-4">
        <img class="size-12 flex-none rounded-full bg-gray-50" :src="paket.photo" alt="Photo" />
        <div class="min-w-0 flex-auto">
          <p class="text-sm/6 font-semibold text-gray-900">{{ paket.name }}</p>
          <p class="mt-1 truncate text-xs/5 text-gray-500">email</p>
        </div>
      </div>
      <div class="hidden shrink-0 sm:flex sm:flex-col sm:items-end">
        <p class="text-sm/6 text-gray-900">Status</p>
        <div class="mt-1 flex items-center gap-x-1.5">
          <p class="text-xs/5 text-gray-500">
            <div v-if="paket.status === 'active'">
              <span class="inline-flex items-center rounded-md bg-green-50 px-2 py-1 text-xs font-medium text-green-700 ring-1 ring-inset ring-green-600/20">Active</span>
            </div>
            <div v-else-if="paket.status === 'pending'">
              <span class="inline-flex items-center rounded-md bg-yellow-50 px-2 py-1 text-xs font-medium text-yellow-800 ring-1 ring-inset ring-yellow-600/20">Pending</span>
            </div>
            <div v-else>
              <span class="inline-flex items-center rounded-md bg-red-50 px-2 py-1 text-xs font-medium text-red-700 ring-1 ring-inset ring-red-600/10">Fail</span>
            </div>
          </p>
        </div>
      </div>
    </li>
  </ul>
</template>
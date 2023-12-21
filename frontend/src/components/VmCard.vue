<template>
    <div
    class="bg-surface-600 w-full rounded-xl border shadow-md overflow-hidden"
    data-v0-t="card"
  >
    <div class="flex">
      <div class="flex-none w-24 h-24 bg-surface-600">
        <img
          src=""
          alt="VM"
          class="object-cover w-full h-full"
          width="96"
          height="96"
          style="aspect-ratio: 96 / 96; object-fit: cover;"
        />
      </div>
      <div class="flex-1 p-3">
        <div class="flex justify-between items-center">
          <h5 class="text-lg font-semibold">VM Name</h5>
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
            class="text-green-500 float-right"
          >
            <circle cx="12" cy="12" r="10"></circle>
          </svg>
        </div>
        <div class="flex justify-between text-sm mt-6 text-primary-600">
          <span>Uptime: 24 days</span>
          <span>RAM: 8GB</span>
          <span>CPU: 4 cores</span>
        </div>
      </div>
    </div>
  </div>
</template>


<script>

export default {

}

</script>
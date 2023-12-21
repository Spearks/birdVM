<script>

import Navbar from './components/Navbar.vue';


export default {
    components: { Navbar }
};

</script>


<template>
  <div class="card">
    <Navbar></Navbar>

    <div class="p-2 w-full h-full">
      <router-view/>
    </div>

    <div class="card fixed z-90 bottom-10 right-8">
      <Button type="button" class="rounded-3xl" icon="pi pi-plus" :loading="loading" @click="load" />
    </div>

  </div>
</template>


<template>
    <div class="card w-full rounded-lg">
        <SetupUi></SetupUi>
        <div class="grid grid-rows-1 grid-cols-1 gap-4">
          
          <VmCard></VmCard>

        </div>
    </div>
</template>

<script>
import { useStateStore } from '../stores/stateStore'
import SetupUi from '../components/SetupUi.vue'
import VmCard from '../components/VmCard.vue'

const stateStore = useStateStore()

export default {
    setup() {
      return {
        isConnected: stateStore.isConnected,
      }
    },
    components: { SetupUi, VmCard }
}
</script>
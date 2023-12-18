<script>
import { CheckIfIsInstalled } from '../../wailsjs/go/backend/App'
import { GenerateConfig } from '../../wailsjs/go/backend/App'
import { TestConnection } from '../../wailsjs/go/backend/App';
import { useStateStore } from '../stores/stateStore'



export default {
    setup() {
      
      const stateStore = useStateStore()

      return {
        isConnected: stateStore.isConnected,
      }
    },
    data() {
        return {
          showInstall: false,
          providers: [ 
            { name: "Libvirt (Linux)", code: "libvirt"},
            { name: "Virtualbox (Windows) ", code: "vbox"}
          ],
          selectedProvider: null,
        };
    },
    created() {

        CheckIfIsInstalled().then(result => {
          this.showInstall = result
        })
    },

    methods: {
      async writeConfig() {

        this.showInstall = false 
        
        GenerateConfig(this.selectedProvider.code)

        let err, connected = await TestConnection("/var/run/libvirt/libvirt-sock")

        this.isConnected = connected

        console.log(this.isConnected)
        
      }
    }
}

</script>

<template>
  <div class="card flex justify-center">
    <Dialog v-model:visible="showInstall" modal header="Setup Agent" :style="{ width: '50rem' }"
      :breakpoints="{ '1199px': '75vw', '575px': '90vw' }">

      <template #header>
        <div class="inline-flex items-center justify-center gap-2">
            <span class="pi pi-star"></span>
            <p>{{ stateStore.isConnected }}</p>
            <span class="font-bold white-space-nowrap">Setup your host system</span>
        </div>
      </template>

      <p class="font-semibold">
        Select an provider
      </p>
      
      <div class="card flex justify-center mt-5">
        <Dropdown v-model="selectedProvider" :options="providers" optionLabel="name" placeholder="Select" class="w-full md:w-14rem" />
      </div>

      <template #footer>
        <Button label="Ok" icon="pi pi-check" @click="writeConfig()" autofocus />
      </template>

    </Dialog>
  </div>
</template>
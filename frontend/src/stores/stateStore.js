import { defineStore } from 'pinia'

export const useStateStore = defineStore({
  id: 'state',
  state: () => ({
    isConnected: null
  })
})
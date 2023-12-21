import { defineStore } from 'pinia'

export const useStateStore = defineStore({
  id: 'state',
  state: () => ({
    isConnected: false,
    provider: null
  })
})
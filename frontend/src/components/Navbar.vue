<template>
    <div id="nav" class="p-1">
        <Menubar :model="items">
            <template #item="{ item, props, hasSubmenu }">
                <router-link v-if="item.route" v-slot="{ href, navigate }" :to="item.route" custom>
                    <a :href="href" v-bind="props.action" @click="navigate">
                        <span :class="item.icon" />
                        <span class="ml-2">{{ item.label }}</span>
                    </a>
                </router-link>
                <a v-else :href="item.url" :target="item.target" v-bind="props.action">
                    <span :class="item.icon" />
                    <span class="ml-2">{{ item.label }}</span>
                    <span v-if="hasSubmenu" class="pi pi-fw pi-angle-down ml-2" />
                </a>
            </template>
            <template #end>
                <ProviderWidget>
                    Libvirt 
                </ProviderWidget>
            </template>
        </Menubar>
    </div>
</template>

<script>
import ProviderWidget from './ProviderWidget.vue';


export default {
    data() {
        return {
            items: [
                {
                    label: 'Home',
                    icon: 'pi pi-home',
                    command: () => {
                        this.$router.push('/');
                    },
                },
                {
                    label: 'Stats',
                    icon: 'pi pi-star',
                    command: () => {
                        this.$router.push('/stats');
                    },
                },
            ],
            visible: true
        };
    },
    components: { ProviderWidget }
}

</script>
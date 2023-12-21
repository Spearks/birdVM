<template>
    <div class="flex items-center text-center bg-primary-400 h-full w-52 rounded-lg p-2">
        <InputText placeholder="Search" type="text" class="w-[8rem] sm:w-auto" />

        <span class="pi pi-check text-white mr-1"></span>
        <p class="text-white uppercase text-sm font-extrabold"> <slot></slot> </p>
    </div>
</template>

<script>

export default {

}

</script>
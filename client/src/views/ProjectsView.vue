<template>
    <h1>{{ currentPage }}</h1>
</template>

<script>
export default {
    data() {
        return {
            currentPage: this.$route.query.page,
        }
    }
}
</script>


<style scoped>

</style>
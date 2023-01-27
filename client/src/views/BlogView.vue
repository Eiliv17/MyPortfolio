<template>
    <h1>{{ currentPage }}</h1>
</template>

<script>
export default {
    props: ["page"],
    data() {
        return {
            currentPage: this.$route.query.page,
        }
    }
}
</script>


<style scoped>

</style>
<template>
    <div class="container">
        <h1>Error 404</h1>
        <p>Page Not Found</p>
        <RouterLink :to="{name: 'home'}" >Return Home</RouterLink>
    </div>
</template>

<script>
import { RouterLink } from 'vue-router'
</script>

<style scoped>
h1 {
    margin-top: 15rem;
}

p {
    margin-bottom: 4rem;
}

div {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: flex-start;
    gap: 1rem;
}
</style>
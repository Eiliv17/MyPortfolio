<script setup>
import { RouterLink } from 'vue-router'
</script>

<template>
    <article class="project">
        <div class="project-body" :class="{swap: theme === 'swap'}">
            <h3 class="project-title">{{ title }}</h3>
            <p class="project-description">{{ description }}</p>
            <ul class="project-techstack">
                <li v-for="tech in techstack">{{ tech }}</li>
            </ul>
            <div class="project-links">
                <a target="_blank" :href="githublink" class="github">GitHub</a>
                <RouterLink :to="{ name: 'projects-article', params: { id: id }}">Learn More &rarr;</RouterLink>
            </div>
        </div>
        <figure class="project-image">
            <img :src="image" :alt="`${title} Project Image`">
        </figure>
    </article>
</template>

<script>
export default {
    props: [
        'id', 
        'title', 
        'description', 
        'techstack', 
        'githublink', 
        'image', 
        'theme'
    ],
}
</script>

<style scoped>

.project {
    display: grid;
    grid-template-columns: 1fr 1fr;
    height: 50vh;
    gap: 4rem;
}

.project-body{
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: center;
    gap: 2rem;
}

.project-description {
    font-size: 3.6rem;
    text-overflow: ellipsis;
}

.project-techstack{
    list-style: none;
    display: flex;
    align-items: center;
    justify-content: flex-start;
    gap: 4rem;
    font-weight: 400;
}

.project-links{
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: flex-start;
    gap: 4rem;
    font-weight: 700;
}

.project-image{
    height: 50vh;
}

.project-image img{
    width: 100%;
    height: 100%;
    object-fit: cover;
}

.swap {
    order: 2;
}


</style>
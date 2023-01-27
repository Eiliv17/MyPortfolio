<script setup>
import { RouterLink } from 'vue-router'
import Avatar from '../components/icons/Avatar.vue'
</script>

<template>
    <section class="hero-section container">
        <p>I'm a</p>
        <h1>Front-end Web <br> Developer</h1>
        <a href="https://google.com" class="btn btn-primary" target="_blank">Download CV</a>
    </section>
    <section class="about-section container">
        <Avatar />
        <div class="about-body">
            <h2 class="about-heading">About Me</h2>
            <p class="about-text">Hi, My name is Alberto Fabro and I am a junior web developer with knowledge of HTML, CSS, and JavaScript. I am passionate about creating websites that are both visually appealing and easy to use. I am always looking to improve my skills and take on new challenges.</p>
            <RouterLink class="btn btn-primary" :to="{name: 'about'}">Get to Know Me</RouterLink>
        </div>
    </section>
    <section class="projects-section">
        <div class="container">

        </div>
    </section>
    <section class="blog-section">
        
    </section>
    <div class="container contact-background">
        <section class="contact-section">
            <p>If you are interested in working with me or have any questions about my skills and experience, please do not hesitate to get in touch. You can reach me by email, or via my social media profiles. I would be more than happy to discuss any potential projects or opportunities with you.</p>
            <RouterLink class="btn btn-dark" :to="{name: 'contact'}">Get in Touch &rarr;</RouterLink>
        </section>
    </div>
</template>

<script>

</script>


<style scoped>

/* Hero Section */

.hero-section{
    min-height: calc(100vh - var(--header-height) - var(--around-padding) * 2);
}

.hero-section h1{
    margin-bottom: 4rem;
}

.hero-section p{
    margin-top: 8rem;
    font-weight: 600;
}

/* About Section */

.about-section {
    min-height: 100vh;
    display: grid;
    grid-template-columns: 1fr 1fr;
    align-items: center;
}

.about-body {
    display: flex;
    flex-direction: column;
    gap: 2rem;
    align-items: flex-start;
}

.about-text {
    font-size: 3.6rem;
}

/* Projects Section */

.projects-section{
    min-height: 200vh;
    width: 100%;
    margin: 10rem auto;
    position: relative;
}

.projects-section::before{
    min-height: 200vh;
    width: 100%;
    position: absolute;
    transform: skewY(5deg);
    background: var(--color-background-soft);
    z-index: 10;
}


/* Contact Section */

.contact-section {
    background: var(--color-primary);
    color: var(--color-background-soft);
    padding: 6rem;
    max-width: 120rem;
    margin: 0 auto;
    font-weight: 600;
    position: relative;
    font-size: 3.6rem;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    gap: 4rem;
}

.contact-background::before {
    content: "";
    background: var(--color-background-soft);
    position: absolute;
    width: 100%;
    height: 10rem;
    bottom: 0;
    z-index: -1;
    left: 0;
}

</style>
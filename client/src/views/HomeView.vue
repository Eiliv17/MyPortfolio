<script setup>
import { RouterLink } from 'vue-router'
import Avatar from '../components/icons/Avatar.vue'
import ProjectShort from '../components/ProjectShort.vue';
import BlogShort from '../components/BlogShort.vue';
</script>

<template>
    <section class="hero-section container">
        <p>I'm a</p>
        <h1>Front-end Web <br> Developer</h1>
        <a href="https://google.com" class="btn btn-primary" target="_blank">Download CV</a>
    </section>
    <section class="about-section container">
        <Avatar class="avatar" />
        <div class="about-body">
            <h2 class="about-heading">About Me</h2>
            <p class="text-description">Hi, My name is Alberto Fabro and I am a junior web developer with knowledge of HTML, CSS, and JavaScript. I am passionate about creating websites that are both visually appealing and easy to use. I am always looking to improve my skills and take on new challenges.</p>
            <RouterLink class="btn btn-primary-outline" :to="{name: 'about'}">Get to Know Me</RouterLink>
        </div>
    </section>
    <section class="projects-section">
        <div class="container">
            <h2 class="projects-heading">Projects</h2>
            <div class="projects-showcase">
                <ProjectShort :v-if="projects.length !== 0" v-for="project, index in projects"
                :id="project.id"
                :title="project.title"
                :description="project.description"
                :techstack="project.techStack"
                :githublink="project.gitHubLink"
                :image="project.image"
                :theme=" index % 2 !== 0 ? 'swap' : ''"
                />
            </div>
            <RouterLink class="btn btn-primary-outline" :to="{name: 'projects'}">View More</RouterLink>
        </div>
    </section>
    <section class="blog-section">
        <div class="container">
            <h2 class="blog-heading">Blog</h2>
            <div class="blog-showcase">
                <BlogShort :v-if="posts.length !== 0" v-for="post in posts"
                :id="post.id"
                :title="post.title"
                :image="post.image"
                :views="post.views"
                :createdAt="post.createdAt"
                />
            </div>
            <RouterLink class="btn btn-primary-outline" :to="{name: 'blog'}">View More</RouterLink>
        </div>
    </section>
    <div class="container contact-background">
        <section class="contact-section">
            <p class="text-description">If you are interested in working with me or have any questions about my skills and experience, please do not hesitate to get in touch. You can reach me by email, or via my social media profiles. I would be more than happy to discuss any potential projects or opportunities with you.</p>
            <RouterLink class="btn btn-dark" :to="{name: 'contact'}">Get in Touch &rarr;</RouterLink>
        </section>
    </div>
</template>

<script>
export default {
  components: {
    ProjectShort,
    BlogShort
  },
  data() {
    return {
        projects: [],
        posts: [],
    }
  },
  methods: {
    
  },
  mounted() {
    fetch('http://localhost:3000/api/projects?' + new URLSearchParams({
        offset: '0',
        limit: '2',
    }))
    .then(res => res.json())
    .then(data => this.projects = data)
    .catch(err => console.log(err.message)); 

    fetch('http://localhost:3000/api/posts?' + new URLSearchParams({
        offset: '0',
        limit: '4',
    }))
    .then(res => res.json())
    .then(data => this.posts = data)
    .catch(err => console.log(err.message)); 
  }
}
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
    gap: 4rem;
}

.about-body {
    display: flex;
    flex-direction: column;
    gap: 2rem;
    align-items: flex-start;
}

/* Projects Section */

.projects-section{
    width: 100%;
    background: var(--color-background-soft);
}

.projects-section .container{
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    gap: 8rem;
}

.projects-heading{
    text-align: center;
    margin-top: 4rem;
}

.projects-showcase {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    gap: 8rem;
}

.projects-section > .container > .btn {
    margin-bottom: 4rem;
}


/* Blog */

.blog-section .container{
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    gap: 8rem;
}

.blog-heading{
    text-align: center;
    margin-top: 4rem;
}

.blog-showcase {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 4rem;
    width: 100%;
    justify-items: stretch;
}

.blog-section > .container > .btn {
    margin-bottom: 4rem;
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
    height: 25rem;
    bottom: 0;
    z-index: -1;
    left: 0;
}

/* For screen sizes lower than 600px  */
@media screen and (max-width: 37.5em) {
    .about-section {
        grid-template-columns: 1fr;
        margin-bottom: 8rem;
    }

    .blog-showcase{
        grid-template-columns: 1fr;
    }
}



</style>
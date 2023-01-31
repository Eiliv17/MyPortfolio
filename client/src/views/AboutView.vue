<script setup>
import Avatar from '../components/icons/Avatar.vue';
import IconCssTechstack from '../components/icons/techstack/IconCssTechstack.vue';
import IconFigmaTechstack from '../components/icons/techstack/IconFigmaTechstack.vue';
import IconGitTechstack from '../components/icons/techstack/IconGitTechstack.vue';
import IconGithubTechstack from '../components/icons/techstack/IconGithubTechstack.vue';
import IconGolangTechstack from '../components/icons/techstack/IconGolangTechstack.vue';
import IconHtmlTechstack from '../components/icons/techstack/IconHtmlTechstack.vue';
import IconJavascriptTechstack from '../components/icons/techstack/IconJavascriptTechstack.vue';
import IconMongodbTechstack from '../components/icons/techstack/IconMongodbTechstack.vue';
import IconNodeTechstack from '../components/icons/techstack/IconNodeTechstack.vue';
import IconPythonTechstack from '../components/icons/techstack/IconPythonTechstack.vue';
import IconSassTechstack from '../components/icons/techstack/IconSassTechstack.vue';
import IconVscodeTechstack from '../components/icons/techstack/IconVscodeTechstack.vue';
import IconVueTechstack from '../components/icons/techstack/IconVueTechstack.vue';
</script>

<template>
    <div class="container">
        <div class="about-main">
            <h1 class="about-heading">About Me</h1>
            <p class="about-text text-description">
                Hi, I'm Alberto Fabro and I'm a Junior Web Developer located in Italy. I've a passion for Web Development and Web Design. I'm a well-organized person and a problem solver with an attention to detail. I'm a fan of Tech and Science since when I was a child. <br> 
                Interested in working on any project even if I don't know where to start I'll find how in one way or another.
            </p>
            <Avatar class="about-avatar avatar" />
        </div>
        <div class="techstack">
            <h2>Tech Stack</h2>
            <div class="techstack-languages">
                <h3>Languages & Frameworks</h3>
                <ul class="techstack-list">
                    <li class="techstack-list-item">
                        <IconHtmlTechstack class="tech-icon" />
                    </li>
                    <li class="techstack-list-item">
                        <IconCssTechstack class="tech-icon" />
                    </li>
                    <li class="techstack-list-item">
                        <IconSassTechstack class="tech-icon" />
                    </li>
                    <li class="techstack-list-item">
                        <IconJavascriptTechstack class="tech-icon" />
                    </li>
                    <li class="techstack-list-item">
                        <IconVueTechstack class="tech-icon" />
                    </li>
                    <li class="techstack-list-item">
                        <IconGolangTechstack class="tech-icon" />
                    </li>
                    <li class="techstack-list-item">
                        <IconPythonTechstack class="tech-icon" />
                    </li>
                </ul>
            </div>
            <div class="techstack-tools">
                <h3>Tools</h3>
                <ul class="techstack-list">
                    <li class="techstack-list-item">
                        <IconGithubTechstack class="tech-icon" />
                    </li>
                    <li class="techstack-list-item">
                        <IconGitTechstack class="tech-icon" />
                    </li>
                    <li class="techstack-list-item">
                        <IconNodeTechstack class="tech-icon" />
                    </li>
                    <li class="techstack-list-item">
                        <IconMongodbTechstack class="tech-icon" />
                    </li>
                    <li class="techstack-list-item">
                        <IconVscodeTechstack class="tech-icon" />
                    </li>
                    <li class="techstack-list-item">
                        <IconFigmaTechstack class="tech-icon" />
                    </li>
                </ul>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    components: {
        Avatar,
    }
}
</script>

<style scoped>
.container {
    margin-bottom: 8rem;
}

.about-main {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 4rem;
    margin-bottom: 8rem;
}

.about-heading{
    grid-column: 1 / -1;
}

.about-text{
    align-self: center;
}

.about-avatar {
    align-self: start;
}

.techstack{
    grid-column: 1 / -1;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: center;
    gap: 4rem;
}

.techstack-list {
    list-style: none;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-wrap: wrap;
    gap: 1rem;
}

.techstack-icon{
    width: 10rem;
    height: 10rem;
}

.techstack-icon svg{
    width: 10rem;
    height: 10rem;
}

@media screen and (max-width: 37.5em) {
    .about-main{
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
    }

    .about-text{
        order: 3;
    }

    .techstack{
        align-items: center;
        text-align: center;
    }
}


</style>

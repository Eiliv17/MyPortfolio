<template>
    <div class="container">
        <div class="article" v-if="article">
            <h1 class="article-heading">{{ title }}</h1>
            <div class="article-info">
                <a target="_blank" :href="githublink" class="article-github">GitHub</a>
                <ul class="techstack-list">
                    <li v-for="tech in techstack">{{ tech }}</li>
                </ul>
            </div>
            <picture v-if="image !== ''">
                <img :src="image" :alt="title">
            </picture>
            <div v-html="textBody" class="article-text">
                
            </div>
        </div>
    </div>
</template>

<script>
export default {
    props: ['id'],
    data() {
        return {
            article: {},
        }
    },
    computed: {
        title() {
            return this.article["title"];
        },
        textBody() {
            return this.article["textBody"];
        },
        views() {
            return this.article["views"];
        },
        githublink() {
            return this.article["gitHubLink"];
        },
        techstack() {
            return this.article["techStack"];
        },
        image() {
            return this.article["image"];
        }
    },
    methods: {
        loadArticle() {
            fetch('http://localhost:3000/api/projects/' + String(this.id))
            .then(res => res.json())
            .then(data => this.article = data)
            .catch(err => console.log(err)); 
        }
    },
    mounted () {
        this.loadArticle();
    }
}
</script>


<style scoped>
.article {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: center;
    gap: 4rem;
    margin-bottom: 8rem;
}

.article-heading {
    width: 100%;
}

.article-info{
    display: flex;
    align-items: center;
    justify-content: space-between;
    width: 100%;
}

.techstack-list {
    list-style: none;
    display: flex;
    align-items: center;
    justify-content: flex-start;
    gap: 4rem;
    font-weight: 400;
    text-transform: uppercase;
}

:deep(p + h2), :deep(p + h3), :deep(p + h4){
    margin-top: 4rem;
}

:deep(p) {
    line-height: 1.4;
}

picture img{
    width: 100%;
    height: 100%;
    object-fit: cover;
}

picture{
    width: 100%;
    height: 50rem;
}


</style>
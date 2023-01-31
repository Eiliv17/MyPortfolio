<template>
    <div class="container">
        <div class="article" v-if="article">
            <h1 class="article-heading">{{ title }}</h1>
            <div class="article-info">
                <time class="article-date">{{ publishDate }}</time>
                <p class="article-views">{{ views }} Views</p>
            </div>
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
            return this.article["title"]
        },
        textBody() {
            return this.article["textBody"]
        },
        views() {
            return this.article["views"]
        },
        publishDate() {
            const date = new Date(this.article["createdAt"])
            const options = { month: "long", day: "numeric", year: "numeric" };
            const formattedDate = date.toLocaleDateString("en-US", options);
            
            return formattedDate
        }
    },
    methods: {
        loadArticle() {
            fetch('http://localhost:3000/api/posts/' + String(this.id))
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

:deep(p + h2), :deep(p + h3), :deep(p + h4){
    margin-top: 4rem;
}

:deep(p) {
    line-height: 1.4;
}


</style>
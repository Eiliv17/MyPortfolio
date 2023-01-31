<script setup>
import BlogShort from '../components/BlogShort.vue';
</script>

<template>
    <div class="container">
        <h2 class="blog-heading">Blog</h2>
        <div class="blog-showcase">
            <BlogShort :v-if="posts" v-for="article in articles"
            :id="article.id"
            :title="article.title"
            :image="article.image"
            :views="article.views"
            :createdAt="article.createdAt"
            />
        </div>
        <div class="navigation">
            <button @click="previousPage" v-if="page > 1" class="btn btn-primary-outline">Previous Page</button>
            <button @click="nextPage" v-if="!lastPage" class="btn btn-primary-outline">Next Page</button>
        </div>
    </div>
</template>

<script>
export default {
    components: {
        BlogShort,
    },
    data() {
        return {
            articles: [],
            page: 1,
            articlesPerPage: 10,
        }
    },
    computed: {
        lastPage() {
            if (this.articles.length < this.articlesPerPage){
                return true;
            } else {
                return false
            }
        }
    },
    methods: {
        scrollToTop () {
            window.scrollTo(0, 0);
        },
        previousPage() {
            this.page = this.page - 1;
            this.loadArticles()
            this.scrollToTop()
        },
        nextPage() {
            if (!this.lastPage) {
                this.page = this.page + 1;
                this.loadArticles()
                this.scrollToTop()
            }
        },
        loadArticles() {
            fetch('http://localhost:3000/api/posts?' + new URLSearchParams({
                offset: String( this.page * 10 - this.articlesPerPage),
                limit: this.articlesPerPage,
            }))
            .then(res => res.json())
            .then(data => this.articles = data)
            .catch(err => console.log(err)); 
        }
    },
    mounted () {
        this.loadArticles();
    }
}
</script>


<style scoped>
.container{
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
    gap: 8rem;
    width: 100%;
}

.navigation{
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 4rem;
}

.btn {
    outline: none;
    cursor: pointer;
}

@media screen and (max-width: 37.5em) {
    .blog-showcase {
        grid-template-columns: 1fr;
    }
}

</style>
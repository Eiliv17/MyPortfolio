<template>
    <article @click="postClick" class="post">
        <h3 class="post-title">{{ title }}</h3>
        <div class="post-info">
            <time class="post-date">{{ publishDate }}</time>
            <p class="views">{{ views }} Views</p>
        </div>
    </article>
</template>

<script>
export default {
    props: [
        'id', 
        'title', 
        'views',
        'createdAt'
    ],
    methods: {
        postClick() {
            this.$router.push({ name: 'blog-article', params: { id: this.id } })
        }
    },
    computed: {
        publishDate() {
            const date = new Date(this.createdAt)
            const options = { month: "long", day: "numeric", year: "numeric" };
            const formattedDate = date.toLocaleDateString("en-US", options);
            
            return formattedDate
        }
    }
}
</script>

<style scoped>

.post {
    height: 30rem;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: space-between;
    background: var(--color-background-soft);
    padding: 4rem;
    cursor: pointer;
    width: 100%;
}

.post-title {
    display: -webkit-box;
    -webkit-line-clamp: 3;
    -webkit-box-orient: vertical;  
    overflow: hidden;
    width: 100%;
}

.post-info {
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: space-between;
}

</style>
<template>
    <article @click="postClick" class="post">
        <figure class="post-image">
            <img :src="image" :alt="`${title} Post Image`">
        </figure>
        <div class="post-body">
            <h3 class="post-title">{{ title }}</h3>
            <div class="post-info">
                <time class="post-date">{{ createdAt }}</time>
                <p class="views">{{ views }} Views</p>
            </div>
        </div>
        
    </article>
</template>

<script>
export default {
    props: [
        'id', 
        'title', 
        'image',
        'views',
        'createdAt'
    ],
    methods: {
        postClick() {
            this.$router.push({ name: 'blog-article', params: { id: this.id } })
        }
    }
}
</script>

<style scoped>

.post {
    background: var(--color-background-soft);
    cursor: pointer;
}

.post-image {
    width: 100%;
    height: 30rem;
}

.post-image img{
    width: 100%;
    object-fit: cover;
    height: 100%;
}

.post-title{
    margin-bottom: 4rem;
    display: -webkit-box;
    text-overflow: ellipsis;
    overflow: hidden;
    max-height: 20rem;
    -webkit-line-clamp: 3;
    -webkit-box-orient: vertical;
}

.post-body{
    padding: 4rem;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
}

.post-info {
    display: flex;
    align-items: center;
    justify-content: space-between;
}

</style>
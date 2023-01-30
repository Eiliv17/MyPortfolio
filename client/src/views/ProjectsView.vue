<script setup>
import ProjectShort from '../components/ProjectShort.vue';
</script>

<template>
    <div class="container">
        <h2 class="projects-heading">Projects</h2>
        <div class="projects-showcase">
            <ProjectShort :v-if="projects" v-for="project, index in projects"
            :id="project.id"
            :title="project.title"
            :description="project.description"
            :techstack="project.techStack"
            :githublink="project.gitHubLink"
            :image="project.image"
            :theme=" index % 2 !== 0 ? 'swap' : ''"
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
        ProjectShort,
    },
    data() {
        return {
            projects: [],
            page: 1,
            projectsPerPage: 10,
        }
    },
    computed: {
        lastPage() {
            if (this.projects.length < this.projectsPerPage){
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
            this.loadProjects()
            this.scrollToTop()
        },
        nextPage() {
            if (!this.lastPage) {
                this.page = this.page + 1;
                this.loadProjects()
                this.scrollToTop()
            }
        },
        loadProjects() {
            fetch('http://localhost:3000/api/projects?' + new URLSearchParams({
                offset: String( this.page * 10 - this.projectsPerPage),
                limit: this.projectsPerPage,
            }))
            .then(res => res.json())
            .then(data => this.projects = data)
            .catch(err => console.log(err)); 
        }
    },
    mounted () {
        this.loadProjects();
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
</style>
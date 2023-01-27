<script setup>
import { RouterLink, RouterView } from 'vue-router'
import IconSocialCodepen from './components/icons/IconSocialCodepen.vue'
import IconSocialFacebook from './components/icons/IconSocialFacebook.vue'
import IconSocialGithub from './components/icons/IconSocialGithub.vue'
import IconSocialInstagram from './components/icons/IconSocialInstagram.vue'
import IconSocialLinkedin from './components/icons/IconSocialLinkedin.vue'
import IconSocialTwitter from './components/icons/IconSocialTwitter.vue'
import Logo from './components/icons/Logo.vue'
</script>

<template>
  <div class="container">
      <header>
        <RouterLink class="logo-link" :to="{name: 'home'}">
            <Logo class="logo"/>
        </RouterLink>
        <nav>
          <RouterLink class="router-link" :to="{name: 'home'}">Home</RouterLink>
          <RouterLink class="router-link" :to="{name: 'about'}">About</RouterLink>
          <RouterLink class="router-link" :to="{name: 'projects'}">Projects</RouterLink>
          <RouterLink class="router-link" :to="{name: 'blog'}">Blog</RouterLink>
          <RouterLink class="router-link" :to="{name: 'contact'}">Contact</RouterLink>
        </nav>
      </header>
  </div>


  <main>
    <RouterView />
  </main>
  
  <footer class="grid">
    <div class="container">
        <!-- <IconSocialCodepen/>
        <IconSocialFacebook/>
        <IconSocialGithub/>
        <IconSocialInstagram/>
        <IconSocialLinkedin/>
        <IconSocialTwitter/> -->
    </div>
  </footer>
</template>




<style scoped>

/* Header */

header {
    height: var(--header-height);
    display: flex;
    align-items: center;
    justify-content: space-between;
}

.logo-link {
    width: auto;
    height: 100%;
}

.logo {
    fill: var(--color-primary);
    width: auto;
    height: 100%;
}

nav {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 3rem;
    font-weight: 600;
}

.router-link:link {
    color: var(--color-text)
}

.router-link-active.router-link-exact-active{
    color: var(--color-primary);
}


/* Footer */

footer {
    width: 100%;
    height: 50vh;
    background: var(--color-background-soft);
    display: grid;
    grid-template-columns: 1.5fr 1.5fr 2fr;
}


</style>

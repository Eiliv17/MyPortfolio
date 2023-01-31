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
        <nav class="header-navigation" :class="{ active : hamburgerActive}">
          <RouterLink @click="hamburgerActive = false" class="router-link" :to="{name: 'home'}">Home</RouterLink>
          <RouterLink @click="hamburgerActive = false" class="router-link" :to="{name: 'about'}">About</RouterLink>
          <RouterLink @click="hamburgerActive = false" class="router-link" :to="{name: 'projects'}">Projects</RouterLink>
          <RouterLink @click="hamburgerActive = false" class="router-link" :to="{name: 'blog'}">Blog</RouterLink>
          <RouterLink @click="hamburgerActive = false" class="router-link" :to="{name: 'contact'}">Contact</RouterLink>
        </nav>
        <div @click="hamburgerActive = !hamburgerActive" class="hamburger" :class="{ active : hamburgerActive}">
            <span class="bar"></span>
            <span class="bar"></span>
            <span class="bar"></span>
        </div>
      </header>
  </div>


  <main @click="hamburgerActive = false">
    <RouterView />
  </main>
  
  <footer>
    <div class="container grid">
        <div class="footer-logo">
            <RouterLink class="logo-link" :to="{name: 'home'}">
                <Logo class="logo"/>
            </RouterLink>
            <p>Alberto Fabro</p>
        </div>
        <nav class="footer-navigation">
          <RouterLink class="router-link" :to="{name: 'home'}">Home</RouterLink>
          <RouterLink class="router-link" :to="{name: 'about'}">About</RouterLink>
          <RouterLink class="router-link" :to="{name: 'projects'}">Projects</RouterLink>
          <RouterLink class="router-link" :to="{name: 'blog'}">Blog</RouterLink>
          <RouterLink class="router-link" :to="{name: 'contact'}">Contact</RouterLink>
        </nav>

        <ul class="social-list">
            <li class="social-item">
                <a target="_blank" href="https://github.com/Eiliv17" class="social-link"><IconSocialGithub class="social-icon"/></a>
            </li>
            <li class="social-item">
                <a target="_blank" href="https://www.linkedin.com/in/alberto-fabro-475b74207/" class="social-link"><IconSocialLinkedin class="social-icon"/></a>
            </li>
            <li class="social-item">
                <a target="_blank" href="https://codepen.io/Eiliv17" class="social-link"><IconSocialCodepen class="social-icon"/></a>
            </li>
            <li class="social-item">
                <a target="_blank" href="https://twitter.com/albertofabro_tw" class="social-link"><IconSocialTwitter class="social-icon"/></a>
            </li>
            <li class="social-item">
                <a target="_blank" href="https://www.facebook.com/albertofabro17" class="social-link"><IconSocialFacebook class="social-icon"/></a>
            </li>
            <li class="social-item">
                <a target="_blank" href="https://www.instagram.com/albertofabrodev/" class="social-link"><IconSocialInstagram class="social-icon"/></a>
            </li>
        </ul>

    </div>
    <div class="copyright-policy">
        <div class="container">
            <p class="copyright">Copyright © 2023. All rights reserved.</p>
            <RouterLink class="router-link" :to="{name: 'cookies'}">Cookie Policy</RouterLink>
        </div>
    </div>
  </footer>
</template>


<script>
export default {
    data() {
        return {
            hamburgerActive: false,
        }
    }
}
</script>


<style scoped>

/* Header */

header {
    height: var(--header-height);
    display: flex;
    align-items: center;
    justify-content: space-between;
    position: relative;
}

.logo-link {
    width: auto;
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
}

.logo {
    fill: var(--color-primary);
    width: auto;
    height: 100%;
}

.header-navigation {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 3rem;
    font-weight: 600;
}

.router-link {
    color: var(--color-text)
}

.router-link-active.router-link-exact-active{
    color: var(--color-primary);
}


/* Hamburger */

.hamburger {
    display: none;
    cursor: pointer;
}

.bar {
    display: block;
    width: 30px;
    height: 2px;
    background: var(--color-text);
    transition: all 0.3s ease-in-out;
}

@media screen and (max-width: 43.75em) {

    .header-navigation{
        position: absolute;
        left: -100%;
        top: var(--header-height);
        flex-direction: column;
        width: 100%;
        background: var(--color-background);
        z-index: 10;
        gap: 2rem;
        padding: 2rem;
        transition: all 0.3s ease-in-out;
    }

    .header-navigation.active{
        left: 0;
    }

    .header-navigation a{
        padding: 2rem;
    }

    .hamburger {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        gap: 7px;
    }

    .hamburger.active .bar:first-child {
        transform: translateY(9px) rotate(45deg);
    }

    .hamburger.active .bar:nth-child(2) {
        opacity: 0;
    }

    .hamburger.active .bar:last-child {
        transform: translateY(-9px) rotate(-45deg);
    }


}



/* Footer */

footer {
    width: 100%;
    height: 50vh;
    background: var(--color-background-soft);
}

footer > .grid{
    display: grid;
    grid-template-columns: 1.5fr 1.5fr 2fr;
    height: 100%;
    align-items: center;
    justify-content: center;
    align-content: center;
    justify-items: center;
}

.footer-logo{
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    gap: 2rem;
}

.footer-logo .logo {
    fill: var(--color-text);
    max-width: 10rem;
}

.footer-navigation{
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: flex-start;
    gap: 2rem;
}

.social-list {
    display: grid;
    grid-template-columns: 1fr 1fr 1fr;
    list-style: none;
    align-items: center;
    justify-content: center;
    align-content: center;
    justify-items: center;
    gap: 2rem;
}


.social-link {
    display: inline-flex;
    justify-content: center;
    align-items: center;
    border: 1px solid var(--color-text);
    padding: 1rem;
}

.social-icon {
    max-width: 7.5rem;
    height: 7.5rem;
    fill: var(--color-text);
}

/* Copyright */

.copyright-policy{
    border-top: 1px solid var(--color-text);
    background: var(--color-background-soft);
}

.copyright-policy div{
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding-top: 4rem;
    padding-bottom: 4rem;
}

/* For screen sizes lower than 700px  */
@media screen and (max-width: 43.75em) {
    .footer-navigation {
        gap: 1rem;
    }

    .social-icon {
        max-width: 5rem;
        height: 5rem;
        fill: var(--color-text);
    }

    .footer-logo .logo {
        width: 100%;
        display: inline-block;
        max-width: 10rem;
    }



}

@media screen and (max-width: 37.5em) {
    .social-list {
        grid-template-columns: 1fr 1fr;
    }
}


</style>

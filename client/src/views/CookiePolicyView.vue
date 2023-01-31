<script setup>
import { RouterLink } from 'vue-router'
</script>

<template>
    <div class="container">
        <h1>Cookie Policy</h1>

        <div class="policy-body">
            <p>This website uses cookies to improve user experience and to track usage statistics. By using our website, you consent to our use of cookies in accordance with this policy.</p>
            <p>Cookies are small text files that are placed on your device when you visit a website. They are widely used in order to make websites work, or work more efficiently, as well as to provide information to the website owners. You can control and delete cookies through your browser settings. However, please note that disabling cookies may affect the functionality of this website.</p>
            <p>This website may use third-party services that place their own cookies on your device. These include analytics services such as Google Analytics, and social media buttons, such as those provided by Facebook and Twitter. We do not have control over these third-party cookies and you should check the respective third-party's website for more information about these.</p>
            <p>We reserve the right to make changes to this policy at any time. Any changes made to the policy will be posted on this page. If you have any questions about this policy, please contact us <RouterLink :to="{ name: 'contact'}">here</RouterLink>.</p>
        </div>
    </div>
</template>

<style scoped>
h1 {
    text-align: center;
    margin-top: 4rem;
    margin-bottom: 2rem;
}

.policy-body{
    max-width: 60ch;
    margin: 0 auto;
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
    text-align: justify;
    margin-bottom: 8rem;
    padding: 3rem;
}

</style>
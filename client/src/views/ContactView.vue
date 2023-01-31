<template>
    <div class="container">
        <div class="contact-head">
            <h1 class="heading">Contact</h1>
            <p :class="{ success : success, error: !success}">{{ result }}</p>
        </div>
        <form @submit.prevent="handleSubmit">
            <div class="form-field">
                <input v-model="fullname" type="text" placeholder="Full Name" aria-label="Full Name" required>
            </div>
            <div class="form-field">
                <input v-model="email" type="text" placeholder="Email" aria-label="Email" required>
            </div>
            <div class="form-field">
                <input v-model="subject" type="text" placeholder="Subject" aria-label="Subject" required>
            </div>
            <div class="form-field">
                <textarea v-model="message" placeholder="Message" id="" cols="30" rows="10" aria-label="Message" required></textarea>
            </div>
            <button type="submit" class="btn btn-primary">Send Message</button>
        </form>
    </div>
</template>

<script>
export default {
    data() {
        return {
            fullname: '',
            email: '',
            subject: '',
            message: '',
            response: {},
            success: false,
        }
    },
    computed: {
        body() {
            return {
                fullName: this.fullname,
                email: this.email,
                subject: this.subject,
                message: this.message,
            }
        },
        result() {
            if (this.response["error"] !== undefined){
                this.success = false;
                return this.response["error"];
            } else {
                this.success = true;
                return this.response["result"];
            }
        }
    },
    methods: {
        handleSubmit() {
            fetch('http://localhost:3000/api/contact', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(this.body)
            })
            .then(resp => resp.json())
            .then(data => this.response = data)
            .catch(err => console.log(err))
        }
    }
}
</script>


<style scoped>
textarea {
    resize: none;
}

.container {
    display: grid;
    grid-template-columns: 1fr 1fr;
    min-height: calc(100vh - var(--header-height) - var(--around-padding) * 2);
    align-items: center;
    justify-items: center;
    gap: 4rem;
    margin-bottom: 8rem;
}

.contact-head{
    display: flex;
    flex-direction: column;
    gap: 2rem;
    align-items: center;
    justify-content: center;
}

form {
    display: flex;
    flex-direction: column;
    gap: 2rem;
    max-width: 40rem;
    width: 100%;
}

input, textarea{
    padding: 1.5rem 1rem;
    width: 100%;
    outline: none;
    border: none;
    background: var(--color-background-soft);
    color: var(--color-text);
}

:focus {
    box-shadow: inset 0px -3px 0px 0px var(--color-primary);
}

.success {
    color: green;
}

.error {
    color: red;
}

@media screen and (max-width: 37.5em) {
    .container {
        grid-template-columns: 1fr;
    }
}

</style>
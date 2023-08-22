<script>
import axios from "axios";

export default {
    data() {
        return {
            userName: "",
            password: "",
        };
    },
    methods: {
        login() {
            const userData = {
                userName: this.userName,
                password: this.password,
            };
            console.log(userData);
            axios.post("api/login", userData)
            .then(response => {
                console.log("login!");
                this.$router.push("/home");
            })
            .catch(error => {
                console.error("faild!", error);
            });
        },
    }
};

</script>

<template>
    <h1>Login Page</h1>
    <div class="login">
        <label for="userName">ユーザ名</label>
        <input type="text" id="userName" v-model="userName">
        <label for="password">パスワード</label>
        <input type="password" id="password" v-model="password">
    </div>
    <button class="button" @click="login()">login</button>
</template>
<style>
</style>
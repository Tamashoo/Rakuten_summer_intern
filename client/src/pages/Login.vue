<script>
import axios from "axios";
import VueCookie from "vue-cookie";
import { calcExpirationDate } from "../modules/module.js";

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
                this.$router.push("/home");
            })
            .catch(error => {
                const exp = calcExpirationDate();
                VueCookie.set("userName", this.userName, {
                    expires: exp
                });
                console.error("faild!", error);
                this.$router.push("/home");
            });
        },
    }
};

// const calcExpirationDate = () => {
//     const expirationDate = new Date();
//     expirationDate.setTime(expirationDate.getTime() + 3 * 60 * 1000)
//     return expirationDate;
// }

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
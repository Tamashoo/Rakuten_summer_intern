<script setup>
import axios from "axios";
import VueCookie from "vue-cookie";
import { calcExpirationDate } from "../modules/module.js";
import { useRouter } from "vue-router";
import { ref } from "vue";

const Router = useRouter();

const userName = ref("");
const password = ref("");
const wrongFlag = ref("");

const login = () => {
    const userData = {
        userName: userName.value,
        password: password.value,
    }
    axios.post("/api/login", userData)
        .then(response => {
            if (response === true) {
                const exp = calcExpirationDate();
                VueCookie.set("userName", userName.value, {
                    expires: exp
                });
                Router.push("/home");
            } else {
                wrongFlag.value = true;
            }
        })
        .catch(error => {
            console.error("faild!", error);
        });
    // この下の処理は後に消す
    const exp = calcExpirationDate();
    VueCookie.set("userName", userName.value, {
        expires: exp
    });
    Router.push("/home");
}

const signup = () => {
    Router.push("/signup");
}

</script>

<template>
    <h1 class="text-2xl font-bold mt-4 text-center">Login Page</h1>
    <div class="w-80 h-auto border-2 rounded-lg text-xl login mx-auto mt-4 text-left flex justify-center">
        <div class="w-auto place-content-center h-auto">
            <form @submit.prevent="login">
                <div class="userName mb-3 mx-auto">
                    <label for="userName">ユーザ名</label><br>
                    <input type="text" class="border-2 mt-4" id="userName" placeholder="user name" v-model="userName" required>
                </div>
                <div class="password mb-3 mx-auto">
                    <label for="password">パスワード</label><br>
                    <input type="password" class="border-2 mt-4" placeholder="password" id="password" v-model="password" required>
                </div>
                <p v-if="wrongFlag === true" class="text-red-600 text-xs">ユーザ名またはパスワードが間違っています</p>
                <div class="flex justify-end">
                    <button class="bg-transparent hover:bg-blue-500 text-blue-700 font-semibold hover:text-white py-2 px-4 border border-blue-500 hover:border-transparent rounded-lg mb-4 mt-4 float">login</button>
                </div>
            </form>
        </div>
    </div>
    <div class="flex justify-end mr-4 mt-4">
        <button class="bg-gray-300 hover:bg-gray-400 text-gray-800 font-bold py-2 px-4 rounded-lg" @click="signup()">signup</button>
    </div>
</template>

<style>
</style>

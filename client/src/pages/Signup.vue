<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const Router = useRouter();

const Name = ref('');
const Password1 = ref('');
const Password2 = ref('');
const passwordFlag = ref("");

const register = () => {
    //登録する処理
    if (checkPassword() === true) {
        Router.push('/login');
    } else {
        passwordFlag.value = true;
    }
}

const login = () => {
    Router.push("/login");
}

const checkPassword = () => {
    if (Password1.value === Password2.value) {
        return true;
    }
    return false;
}

</script>
<template>
    <h1 class="text-2xl font-bold mt-4 text-center">Signup Page</h1>
    <div class="w-80 h-auto border-2 rounded-lg text-xl login mx-auto mt-4 text-left flex justify-center">
        <div class="w-auto place-content-center h-auto">
            <form @submit.prevent="register">
                <div class="userName mb-3 mx-auto">
                    <p>ユーザ名</p>
                    <input required class="input border-2 mt-4" v-model="Name" type="text" placeholder="user name" />
                </div>
                <div class="password mb-3 mx-auto">
                    <p>パスワード</p>
                    <input required class="input border-2 mt-4" v-model="Password1" type="password" placeholder="password" />
                </div>
                 <div class="password mb-3 mx-auto">
                    <p>パスワード(再入力)</p>
                    <input required class="input border-2 mt-4" v-model="Password2" type="password" placeholder="password" />
                </div>
                <p v-if="passwordFlag === true" class="text-red-600 text-xs">パスワードが異なります</p>
                <div class="flex justify-end">
                    <button class="button bg-transparent hover:bg-blue-500 text-blue-700 font-semibold hover:text-white py-2 px-4 border border-blue-500 hover:border-transparent rounded-lg mb-4 mt-4 float">登録</button>
                </div>
            </form>
        </div>
    </div>
    <div class="flex justify-end mr-4 mt-4">
        <button class="bg-gray-300 hover:bg-gray-400 text-gray-800 font-bold py-2 px-4 rounded-lg" @click="login()">back to login page</button>
    </div>
</template>

<style>
</style>
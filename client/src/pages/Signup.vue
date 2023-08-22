<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const Router = useRouter();

const Name = ref('');
const Password = ref('');

const register = () => {
    //登録する処理

    Router.push('/home');
}

</script>
<template>
    <h1 class="aiueo">サインアップ</h1>
    <form @submit="register">
        <p>名前</p>
        <input required class="input" v-model="Name" placeholder="名前" />
        <p>パスワード</p>
        <input required class="input" v-model="Password" placeholder="パスワード" />
        <p />
        <button class="button">登録</button>
    </form>
</template>
<style>
.aiueo{
    text-align: center;
}
.button{
    color: blue;
}
</style>
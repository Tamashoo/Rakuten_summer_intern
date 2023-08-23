<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import { useRouter } from 'vue-router';
import { checkCookie } from '../modules/module';

const Router = useRouter();

const result = ref(10);

onMounted(() => {
    checkCookie();
    axios.get('api/receiptresult')
    .then((responce) => {
        result.value = responce;
    })
    .catch(error => {
        console.log(error);
    })
});

</script>
<template>
    <h1>送信完了</h1>
    
    <div class="rounded-3xl w-80 h-80 border-4 border-rakuten text-center mx-auto mt-6 bg-amber-100">
        <h2 class="font-bold text-2xl text-rakuten">獲得Exp</h2>
    <div class="text-9xl text-rakuten font-black m-10">{{ result }}</div>
    </div>
    <button @click="Router.push('/home')" class="HomeBtn">ホーム画面に戻る</button>
</template>
<style>
.Result{
    text-align: center;
    font-size: 100px;
}
.HomeBtn{
    position: absolute;
    bottom: 20px;
}
</style>
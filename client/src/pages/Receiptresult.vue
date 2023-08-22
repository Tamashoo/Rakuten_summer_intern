<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import { useRouter } from 'vue-router';

const Router = useRouter();

const result = ref(10);

onMounted(() => {
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
    <h2>獲得Exp:</h2>
    <div class="Result">{{ result }}</div>
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
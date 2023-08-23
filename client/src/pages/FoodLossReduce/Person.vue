<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import { useRouter } from 'vue-router';

const Router = useRouter();

const ReduceScore = ref(1234);

onMounted(() => {
    axios.get("api/hoodlossreduce/person")
    .then(response => {
        ReduceScore.value = response;
    })
    .catch(error => {
        console.log(error);
    })
});

</script>
<template>
    <h1>あなたのフードロス削減度</h1>
    <div class="ReduceScore">{{ ReduceScore }}</div>
    <button @click="Router.push('/home')" class="HomeBtn">ホーム画面に戻る</button>
</template>
<style>
.ReduceScore{
    text-align: center;
    font-size: 100px;
}
.HomeBtn{
    position: absolute;
    bottom: 20px;
}
</style>
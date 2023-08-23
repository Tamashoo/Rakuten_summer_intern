<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import  axios  from 'axios';
import { checkCookie } from '../../modules/module';

const Router = useRouter();

const AllReduce = ref(23456);

onMounted(() => {
    checkCookie();
    axios.get("api/foodlossreduce/all")
    .then((response) => {
        AllReduce.value = response;
    })
    .catch((error) => {
        console.log(error);
    })
});

</script>
<template>
    <h1>みんなのフードロス削減度</h1>
    <div class="AllReduce">{{ AllReduce }}</div>
    <button @click="Router.push('/home')" class="HomeBtn">ホーム画面に戻る</button>
</template>
<style>
.AllReduce{
    text-align: center;
    font-size: 100px;
}
.HomeBtn{
    position: absolute;
    bottom: 20px;
}
</style>
<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import { useRouter } from 'vue-router';
import { checkCookie } from '../../modules/module';
import HomeBtn  from "@/components/HomeBtn.vue";

const Router = useRouter();

const ReduceScore = ref(1234);

onMounted(() => {
    checkCookie();
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
    <HomeBtn />
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
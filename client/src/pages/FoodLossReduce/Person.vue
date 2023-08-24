<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import { useRouter } from 'vue-router';
import { checkCookie } from '../../modules/module';
import HomeBtn  from "@/components/HomeBtn.vue";

const Router = useRouter();

const ReduceScore = ref(1234567);

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
    <div class="rounded-full border-4 border-rakuten bg-amber-100 w-80 h-80 relative mx-auto my-6">
        <div class="text-7xl font-black text-center text-rakuten mt-28">{{ ReduceScore }}</div>
    </div>
    <HomeBtn />
</template>
<style>
</style>
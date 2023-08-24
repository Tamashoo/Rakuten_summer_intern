<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import  axios  from 'axios';
import { checkCookie } from '../../modules/module';
import HomeBtn  from "@/components/HomeBtn.vue";

const Router = useRouter();

const AllReduce = ref(2345678);

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
    <div class="rounded-full border-4 border-rakuten bg-gradient-to-br from-amber-200 via-amber-50 to-amber-200 w-80 h-80 relative mx-auto mt-6">
        <div class="text-7xl font-black text-center text-rakuten mt-28">{{ AllReduce }}</div>
    </div>
    <HomeBtn />
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
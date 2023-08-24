<script setup>
import { ref, onMounted } from "vue";
import axios from 'axios';
import { useRouter } from 'vue-router';
import { checkCookie } from "../modules/module";
import HomeBtn  from "@/components/HomeBtn.vue";

const Router = useRouter();

const HistoryList = ref([
    { "createdat": 20230820, "getexp": 100 },
    { "createdat": 20230821, "getexp": 120 },
    { "createdat": 20230822, "getexp": 130 },
    { "createdat": 20230823, "getexp": 140 },
])

onMounted(() => {
    checkCookie();
    axios.get("api/history")
    .then((response) => {
        HistoryList.value = response.historylist;
    })
    .catch((error) => {
        console.log(error);
    })
});
</script>

<template>
    <h1>履歴</h1>
    <table class="relative mx-auto border-4 border-rakuten mt-6 table-fixed w-80">
        <thead>
            <tr class="font-bold text-4xl bg-amber-200 text-center">
                <th class="">日付</th>
                <th>獲得Exp</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="(item, index) in HistoryList" :key="index " class="text-xl font-normal">
                <td class="text-center" :class="[(index%2)==0?'bg-orange-100':'bg-amber-100']">{{ item.createdat }}</td>
                <td class="text-right" :class="[(index%2)==0?'bg-orange-50':'bg-amber-50']">{{ item.getexp }}</td>
            </tr>
        </tbody>
    </table>
    <HomeBtn />
</template>

<style>
</style>
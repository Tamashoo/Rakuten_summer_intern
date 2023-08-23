<script setup>
import { ref, onMounted } from "vue";
import axios from 'axios';
import { useRouter } from 'vue-router';

const Router = useRouter();

const HistoryList = ref([
    { "createdat": 20230820, "getexp": 100 },
    { "createdat": 20230821, "getexp": 120 },
    { "createdat": 20230822, "getexp": 130 },
    { "createdat": 20230823, "getexp": 140 },
])

onMounted(() => {
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
    <table>
        <thead>
            <tr>
                <th>日付</th>
                <th>獲得Exp</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="(item, index) in HistoryList" :key="index">
                <td>{{ item.createdat }}</td>
                <td>{{ item.getexp }}</td>
            </tr>
        </tbody>
    </table>
    <button @click="Router.push('/home')" class="HomeBtn">ホーム画面に戻る</button>
</template>

<style>
table,
td {
    border: 1px solid #333;
}

th {
    border: 1px solid #333;
}
.HomeBtn{
    position: absolute;
    bottom: 20px;
}
</style>
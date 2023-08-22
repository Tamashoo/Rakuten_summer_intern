<script setup>
import { ref } from 'vue';
import router from '../router';
import { useRouter } from 'vue-router';

//ダミーデータ
const Character = ref('https://1.bp.blogspot.com/-vsrjLB8Wyng/YEGP_R749_I/AAAAAAABdck/-khfythfWqU5wb6X4FzfzeaJVhIztvByACNcBGAsYHQ/s896/buranko_boy_sad.png');
const Level = ref('3');
const Exp = ref('3200');

const Router = useRouter();

const OpenMenuFlg = ref(false);

const MenuItems = ref([
    { name: 'あなたのフードロス貢献度', id: '/foodlossreduce/person'},
    { name: 'みんなのフードロス貢献度', id: '/foodlossreduce/all'},
    { name: '履歴', id: '/history'},
    { name: 'レシート送信', id: '/receipt'},
]);

const ClickedMenu = (targetIndex) => {
    Router.push(MenuItems.value[targetIndex].id);
}

</script>
<template>
    <h1>HOME</h1>
    <p>{{ 'Lv. ' + Level }}</p>
    <p>{{ 'Exp. ' + Exp }}</p>
    <img :src="Character" />

    <button v-if="OpenMenuFlg" class="menubtn" @click="OpenMenuFlg = !OpenMenuFlg">X</button>
    <button v-else class="menubtn" @click="OpenMenuFlg = !OpenMenuFlg">menu</button>
    <div v-if="OpenMenuFlg" class="menu">
        <div v-for="(item, index) in MenuItems" :key="index">
            <div @click="ClickedMenu(index)">{{ item.name }}</div>
        </div>
    </div>
</template>
<style>
.menubtn{
    position: absolute;
    bottom: 20px;
    right: 20px;
}
.menu{
    position: absolute;
    bottom: 40px;
    right: 20px;
}
</style>
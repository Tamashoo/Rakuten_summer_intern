<script setup>
import { ref, onMounted } from 'vue';
import router from '../router';
import { useRouter } from 'vue-router';
import  axios  from 'axios';
import VueCookie from "vue-cookie";
import { checkCookie } from '../modules/module';

//ダミーデータ
const Character = ref('https://r.r10s.jp/evt/event/okaimonopanda/common/download/wallpaper_201501.png');
const Level = ref('3');
const Exp = ref('3200');

onMounted(() => {
    checkCookie();
    //apiの処理
    axios.get("api/home")
    .then(function(responce) {
        CharacterData.value = responce;
    })
    .catch(error => {
        console.log("faild", error);
    })
});

const Router = useRouter();

const OpenMenuFlg = ref(false);

const MenuItems = ref([
    { name: 'あなたのフードロス貢献度', id: '/foodlossreduce/person'},
    { name: 'みんなのフードロス貢献度', id: '/foodlossreduce/all'},
    { name: '履歴', id: '/history'},
    { name: 'レシート送信', id: '/receipt'},
    { name: 'ログアウト', id: '/login'}
]);

const ClickedMenu = (targetIndex) => {
    if (MenuItems.value[targetIndex].id === "/login") {
        deleteCookie();
    }
    Router.push(MenuItems.value[targetIndex].id);
}

const deleteCookie = () => {
    VueCookie.delete("userName");
}

</script>
<template>
    <h1>HOME</h1>
    <div class="flex justify-center border-4 border-rakuten rounded-full bg-amber-100 mx-2">
        <div class="w-16 h-16 border-4 border-amber-500 bg-white p-2 m-2 rounded-full font-bold drop-shadow text-center text-4xl text-rakuten">{{ Level }}</div>
        <p class="text-4xl m-auto text-center font-bold text-rakuten drop-shadow">{{ 'Exp. ' + Exp }}</p>
    </div>
    <div class="flex">
        <img :src="Character" class="w-80 h-80 mx-auto mt-6" />
    </div>

        <button v-if="OpenMenuFlg==false" class="absolute inset-x-8 bottom-5 bg-rakuten p-3 rounded-full font-bold text-white" @click="OpenMenuFlg = !OpenMenuFlg">メニュー</button>
    <div v-if="OpenMenuFlg" class="absolute z-0 inset-x-8 mx-auto bottom-5 text-2xl bg-white border border-rakuten rounded-xl">
        <div v-for="(item, index) in MenuItems" :key="index">
            <div @click="ClickedMenu(index)" class="text-center border-b">{{ item.name }}</div>
        </div>
        <div @click="OpenMenuFlg=false" class="text-center border-t border-rakuten">閉じる</div>
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
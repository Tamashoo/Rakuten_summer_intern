<script setup>
import axios from "axios";
import { useRouter } from "vue-router";
import { onMounted, ref } from "vue";
import { checkCookie, getCookie } from "../modules/module";
import HomeBtn  from "../components/HomeBtn.vue";

const Router = useRouter();
const imageURL = ref("");
const FileName = ref('');
const submitFlag = ref("");

onMounted(() => {
    checkCookie();
});

const uploadReceipt = (e) => {
    const file = e.target.files[0];
    FileName.value = e.target.files[0].name;
    createImage(file);
}

const submitReceipt = () => {
    if (imageURL.value === "") {
        alert("ファイルを選択してください。");
    } else {
        const data = {
            username: getCookie(),
            receipt: imageURL.value,
        };
        axios.post("http://13.211.209.41:8080/receipt", data)
            .then(response => {
                if (response.data.result === true) {
                    Router.push("/receiptresult")
                } else {
                    submitFlag.value = true;
                } 
            })
            .catch(error => {
                console.error("faild!", error);
            })
    }
};

// inputにアップロードされた画像ファイルを画像化(base64)にする
const createImage = (file) => {
    const reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = () => {
        imageURL.value = reader.result;
    }
};

</script>
<template>
    <h1>レシート送信</h1>
    <div class="upload ">
        <label class="mt-6 rounded-full bg-amber-500 text-white font-medium text-2xl text-center absolute inset-x-8 p-3">
            ファイル選択
            <input class="hidden" ref="file" type="file" id="file" accept=".png, .jpg, jpeg" @change="uploadReceipt">
        </label>
    </div>
    <p v-if="submitFlag === true">レシートを認識できませんでした</p>
    <div class=" relative mt-24 mb-6 mx-auto text-center w-60 h-96 bg-amber-100">
        <h2>プレビュー</h2>
        <div v-if="imageURL" >
        <p>{{ FileName }}</p>
        <img class="" :src="imageURL"  alt="uploadReceiptImage">
    </div>
    </div>
    <button class="rounded-full bg-rakuten text-2xl text-white absolute inset-x-8 font-medium my-5 p-3" @click="submitReceipt()">送信する</button>
    <HomeBtn />
</template>
<style>
</style>
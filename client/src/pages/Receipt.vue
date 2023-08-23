<script setup>
import axios from "axios";
import { useRouter } from "vue-router";
import { onMounted, ref } from "vue";
import { checkCookie } from "../modules/module";

const Router = useRouter();
const imageURL = ref("");
const FileName = ref('');

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
            image: imageURL.value
        }
        axios.post("api/receipt", data)
            .then(response => {
                Router.push("/receiptresult")
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
        <label class="mt-6 rounded-full bg-amber-500 text-white font-medium text-2xl text-center absolute inset-x-8">
            ファイル選択
            <input class="hidden" ref="file" type="file" id="file" accept=".png, .jpg, jpeg" @change="uploadReceipt">
        </label>
    </div>
    <div class="mt-16 mb-6 mx-auto text-center w-60 h-96 bg-amber-100">
        <h2>プレビュー</h2>
        <div v-if="imageURL" >
        <p>{{ FileName }}</p>
        <img class="" :src="imageURL"  alt="uploadReceiptImage">
    </div>
    </div>
    <button class="rounded-full bg-rakuten text-2xl text-white absolute inset-x-8 bottom-8 font-medium" @click="submitReceipt()">送信する</button>
</template>
<style>
</style>
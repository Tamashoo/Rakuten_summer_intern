<script setup>
import axios from "axios";
import { useRouter } from "vue-router";
import { ref } from "vue";

const Router = useRouter();
const imageURL = ref("");

const uploadReceipt = (e) => {
    const file = e.target.files[0];
    createImage(file);
}

const submitReceipt = () => {
    if (imageURL.value === "") {
        console.log("ファイルを選択してください。");
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
    <h1>Receipt</h1>
    <div class="upload">
        <input ref="file" type="file" id="file" accept=".png, .jpg, jpeg" @change="uploadReceipt">
        <button class="button" @click="submitReceipt()">送信</button>
    </div>
    <div class="preview">
        <img :src="imageURL" alt="uploadReceiptImage">
    </div>
</template>
<style>
</style>
<template>
  <h2>Home</h2>
  <p>Welcome to the Note's Home page!</p>
  <el-row>
    <el-col :span="24">
      <el-button style="width: 100%;" type="primary" @click="router.push({ name: 'new' })">Create A New Note</el-button>
    </el-col>
  </el-row>
  <el-divider>OR</el-divider>
  <el-row :gutter="10">
    <el-col :span="24 - 8">
        <el-input v-model="slug" placeholder="Enter a slug"></el-input>
    </el-col>
    <el-col :span="8">
        <el-button style="width: 100%;" type="primary" @click="router.push({ name: 'view', params: { slug } })" :disabled="slug == ''">View Notes</el-button>
    </el-col>
  </el-row>
</template>
<script setup lang="ts">
import { useRouter } from 'vue-router';
import { ref } from 'vue'
const router = useRouter()
const slug = ref('')
</script>

<template>
  <el-form v-model="new_form">
    <el-form-item label="Title">
      <el-input v-model="new_form.title"></el-input>
    </el-form-item>
    <el-form-item label="Is Encrypted">
      <el-switch v-model="new_form.is_encrypted"></el-switch>
    </el-form-item>
    <el-form-item label="Password" v-if="new_form.is_encrypted">
      <el-input v-model="new_form.password" type="password"></el-input>
    </el-form-item>
    <el-form-item label="Content">
      <MdEditor v-model="new_form.content"></MdEditor>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="createNote">Create Note</el-button>
    </el-form-item>
  </el-form>
</template>
<script setup lang="ts">
import { ref } from 'vue'
import { MdEditor } from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';
import { encrypt } from './utils/crypto';
import { generatePasswd, generateSalt } from './utils/gen';
import { useRouter } from 'vue-router';
import { create_note } from './utils/store';
const router = useRouter()

const new_form = ref({
  title: '',
  content: '',
  is_encrypted: false,
  password: ''
})
async function createNote() {
  const { title, content, is_encrypted, password } = new_form.value;
  const salt = is_encrypted ? generateSalt() : null;
  const encrypted_content = is_encrypted ? encrypt(content ?? '', password, salt ?? '') : content;
  const encrypted_title = is_encrypted ? encrypt(title ?? '', password, salt ?? '') : title;
  const encrypted_tag = is_encrypted ? encrypt('tag', password, salt ?? '') : '';
  const encrypted_verification_tag = is_encrypted ? encrypt('verification', password, salt ?? '') : '';

  await create_note({
    'title': encrypted_title,
    'body': encrypted_content,
    'is_encrypted': is_encrypted,
    'encryption_tag': encrypted_tag,
    'encryption_salt': salt,
    'encryption_verification_tag': encrypted_verification_tag,
  }, router);

  ElMessage.success('Note created successfully')
}
</script>

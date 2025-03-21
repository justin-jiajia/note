<template>
  <AskPasswdComponent ref="askpasswd" />
  <el-form v-model="edit_form">
    <el-form-item label="Title">
      <el-input v-model="edit_form.title"></el-input>
    </el-form-item>
    <el-form-item label="Content">
      <MdEditor v-model="edit_form.body"></MdEditor>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="upd" :loading="update_loading">Update Note</el-button>
    </el-form-item>
  </el-form>
</template>
<script setup lang="ts">
import AskPasswdComponent from './AskPasswdComponent.vue';
import { ref, watch } from 'vue'
import { MdEditor } from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';
import { config, XSSPlugin } from 'md-editor-v3';
import { useRoute, useRouter } from 'vue-router';
import { fetch_note, cur_note, updatecurrent } from './utils/store';
const route = useRoute()
const router = useRouter()
const askpasswd = ref(null)
const update_loading = ref(false)

const edit_form = ref({
  title: '',
  body: '',
})

const updShow = async (slug: string) => {
  await fetch_note(slug, false)
  await askpasswd.value.askPasswd()
  edit_form.value.title = cur_note.value.title
  edit_form.value.body = cur_note.value.body
}

const upd = async () => {
  update_loading.value = true
  await updatecurrent(router, edit_form.value)
  update_loading.value = false
}

watch(() => route.params.slug, updShow, { immediate: true });

config({
  markdownItPlugins(plugins) {
    return [
      ...plugins,
      {
        type: 'xss',
        plugin: XSSPlugin,
        options: {},
      },
    ];
  },
});
</script>

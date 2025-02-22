<template>
  <AskPasswdComponent ref="askpasswd" />
  <el-container>
    <el-main id="preview-container">
      <el-row justify="center">
        <el-col :span="6" style="font-size: larger; text-align: center;">
          <strong>{{ cur_note.title }}</strong>
        </el-col>
        <el-col :span="6">
          <span style="font-size: x-small;">Updated At: {{ format_date(cur_note.updated_at) }} <br /> Created At: {{ format_date(cur_note.created_at) }}</span>
        </el-col>
        <el-col :span="12">
          <el-button
            :loading="refresh_loading"
            :icon="Refresh"
            @click="refresh()">
          </el-button>
          <el-button
            :icon="Edit"
            @click="$router.push({ name: 'edit', params: { slug: route.params.slug } })">
          </el-button>
          <el-button
          :icon="Share"
          @click="sharecurruent()">
          </el-button>
          <el-button
            :loading="delete_loading"
            :icon="Delete"
            @click="_delete()">
          </el-button>
        </el-col>
      </el-row>
      <el-divider></el-divider>
      <MdPreview style="max-height: 70vh;" :id="id" :modelValue="cur_note.body" />
    </el-main>
    <el-aside width="10vw">
      <MdCatalog :editorId="id" :scroll-element="scrollElement"/>
    </el-aside>
  </el-container>
</template>

<script setup>
import { Edit, Share, Delete, Refresh } from '@element-plus/icons-vue';
import { ref, watch } from 'vue';
import { MdPreview, MdCatalog } from 'md-editor-v3';
import 'md-editor-v3/lib/preview.css';
import { cur_note, deletecurrent, sharecurruent, fetch_note, decrypt_note } from './utils/store';
import { format_date } from './utils/date';
import { useRoute, useRouter } from 'vue-router';
import AskPasswdComponent from './AskPasswdComponent.vue';
const id = 'preview-only';
const scrollElement = document.getElementById("preview-container");
const route = useRoute();
const router = useRouter();
const askpasswd = ref(null);
const refresh_loading = ref(false);
const delete_loading = ref(false);

async function refresh() {
  refresh_loading.value = true;
  await fetch_note(route.params.slug, true)
  decrypt_note()
  refresh_loading.value = false;
}

async function _delete() {
  delete_loading.value = true;
  await deletecurrent(router)
  delete_loading.value = false;
}

async function viewNote(slug) {
  await fetch_note(slug, false)
  await askpasswd.value.askPasswd()
}

watch(() => route.params.slug, viewNote, { immediate: true });
</script>
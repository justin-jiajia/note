<template>
  <el-dialog v-model="dialogHistoryVisible" title="Histories" width="800">
    <el-table :data="cur_note.histories" style="width: 100%">
      <el-table-column property="created_at_text" label="Date" />
      <el-table-column property="title" label="Note Name" />
      <el-table-column label="Operations">
        <template #default="scope">
          <el-button @click="preview = scope.row.body; previewVisible = true" >
            Preview
          </el-button>
        </template>
      </el-table-column>
    </el-table>
  </el-dialog>
  <el-dialog v-model="previewVisible" title="Preview" width="800">
    <MdPreview :modelValue="preview" />
  </el-dialog>
  <AskPasswdComponent ref="askpasswd" />
  <el-container>
    <el-main id="preview-container">
      <el-row justify="center">
        <el-col style="font-size: larger; text-align: center;">
          <strong>{{ cur_note.title }}</strong>
        </el-col>
      </el-row>
      <el-row>
        <el-col style="font-size: x-small; text-align: end;">
          Updated: {{ format_date(cur_note.updated_at) }} <br /> Created: {{ format_date(cur_note.created_at) }}
        </el-col>
      </el-row>
      <el-divider></el-divider>
      <el-row justify="end">
        <el-col style="text-align: end;">
          <el-button-group>
            <el-button
              :loading="refresh_loading"
              :icon="RefreshIcon"
              @click="refresh()">
            </el-button>
            <el-button
              :icon="ClockIcon"
              @click="dialogHistoryVisible = !dialogHistoryVisible">
            </el-button>
            <el-button
              :icon="EditIcon"
              @click="$router.push({ name: 'edit', params: { slug: route.params.slug } })">
            </el-button>
            <el-button
            :icon="ShareIcon"
            @click="sharecurruent()">
            </el-button>
            <el-button
              :loading="delete_loading"
              :icon="DeleteIcon"
              @click="_delete()">
            </el-button>
            <el-button
              :icon="DocumentCopyIcon"
              @click="copy_source()">
            </el-button>
            <el-button
              :icon="DownloadIcon"
              @click="download_source()">
            </el-button>
          </el-button-group>
        </el-col>
      </el-row>
      <MdPreview style="max-height: 70vh;" :id="id" :modelValue="cur_note.body" />
    </el-main>
    <!--
    <el-aside width="10vw">
      <MdCatalog :editorId="id" :scroll-element="scrollElement"/>
    </el-aside>
    -->
  </el-container>
</template>

<script setup lang="ts">
import { Refresh as RefreshIcon, Edit as EditIcon, Share as ShareIcon, Delete as DeleteIcon, DocumentCopy as DocumentCopyIcon, Download as DownloadIcon, Clock as ClockIcon } from '@element-plus/icons-vue';
import { ref, watch } from 'vue';
import { MdPreview, MdCatalog } from 'md-editor-v3';
import 'md-editor-v3/lib/preview.css';
import { cur_note, deletecurrent, sharecurruent, fetch_note, decrypt_note, copy_source, download_source } from './utils/store.js';
import { format_date } from './utils/date.js';
import { useRoute, useRouter } from 'vue-router';
import AskPasswdComponent from './AskPasswdComponent.vue';
const id = 'preview-only';
const scrollElement = document.getElementById("preview-container");
const route = useRoute();
const router = useRouter();
const preview = ref("");
const previewVisible = ref(false);
const askpasswd = ref<{ askPasswd: () => Promise<void> } | null>(null);
const dialogHistoryVisible = ref(false);
const refresh_loading = ref(false);
const delete_loading = ref(false);

async function refresh() {
  refresh_loading.value = true;
  await fetch_note(route.params.slug as string, true)
  decrypt_note()
  refresh_loading.value = false;
}

async function _delete() {
  delete_loading.value = true;
  await deletecurrent(router)
  delete_loading.value = false;
}

async function viewNote(slug: string) {
  await fetch_note(slug, false)
  if (askpasswd?.value) {
    await askpasswd.value.askPasswd();
  }
}

watch(() => route.params.slug as string | undefined, (slug: string | undefined) => {
  if (slug) {
    viewNote(slug);
  }
}, { immediate: true });
</script>

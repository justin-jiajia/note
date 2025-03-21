<template>
  <el-dialog
    title="Password Required"
    v-model="dialogVisible">
    <el-form :model="pa" :rules="rules" status-icon>
      <el-form-item prop="passwd">
        <el-input
          v-model="pa.passwd"
          placeholder="Enter Password"
          :required="true"
          focus
          id="passwd"
          show-password>
        </el-input>
      </el-form-item>
      <el-form-item>
        <el-switch
          v-model="pa.remember_passwd"
          inline-prompt
          active-text="Remember Password"
          inactive-text="Forget Password">
        </el-switch>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>

<script setup lang="ts">
import { cur_note, verfiy_passwd, decrypt_note } from './utils/store';
import { ref } from 'vue';
import { useRoute } from 'vue-router';

const dialogVisible = ref(false);
const route = useRoute();
const pa = ref({
  passwd: '',
  remember_passwd: true,
});
let resolvePassword = null; // Will hold the Promise resolver

const rules = ref({
  passwd: [
    {
      validator: (_, value, callback) => {
        if (verfiy_passwd(value)) {
          if (pa.value.remember_passwd)
            localStorage.setItem('passwd_' + route.params.slug, pa.value.passwd);
          decrypt_with_passwd();
          if (resolvePassword) { resolvePassword(); resolvePassword = null; }
          callback();
        } else {
          callback(new Error("Wrong Password"));
        }
      },
      trigger: 'change'
    },
    {
      required: true,
      message: 'Please input the password',
      trigger: 'change'
    }
  ],
});

async function askPasswd() {
  if (!cur_note.value.is_encrypted || cur_note.value.decrypted) return;
  
  if (cur_note.value.passwd !== '') {
    pa.value.passwd = cur_note.value.passwd;
    if (!verfiy_passwd(pa.value.passwd)) {
      ElMessage.error('Wrong Password remembered');
      pa.value.passwd = '';
    }
  }
  if (location.hash) {
    pa.value.passwd = location.hash.slice(1);
    if (!verfiy_passwd(pa.value.passwd)) {
      ElMessage.error('Wrong Password on the location hash');
      pa.value.passwd = '';
    }
  } else if (localStorage.getItem('passwd_' + route.params.slug)) {
    pa.value.passwd = localStorage.getItem('passwd_' + route.params.slug);
    if (!verfiy_passwd(pa.value.passwd)) {
      ElMessage.error('Wrong Password stored');
      localStorage.removeItem('passwd_' + route.params.slug);
      pa.value.passwd = '';
    }
  }
  if (pa.value.passwd === '') {
    dialogVisible.value = true;
    return new Promise((resolve) => {
      resolvePassword = resolve;
    });
  } else {
    decrypt_with_passwd();
  }
}

function decrypt_with_passwd() {
  dialogVisible.value = false;
  decrypt_note(pa.value.passwd);
}

defineExpose({ askPasswd });
</script>

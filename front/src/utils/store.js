import { ref } from 'vue'
import { encrypt, decrypt } from './crypto';

export const cur_note = ref({
  title: '',
  body: '',
  created_at: 0,
  updated_at: 0,
  encryption_tag: '',
  encryption_salt: '',
  is_encrypted: false,
  slug: '----',
  passwd: '',
})

export const fetch_note = async (slug, force) => {
  if (slug == cur_note.value.slug && !force) {
    return
  }
  const res = await fetch(import.meta.env.VITE_API_BASE + 'api/v1/notes/' + slug);
  const resjson = await res.json();
  if (!res.ok) {
    ElMessage.error(resjson.error);
    return;
  }
  cur_note.value = { ...cur_note.value, ...resjson }
}

export const verfiy_passwd = (passwd) => {
  if (cur_note.value.encryption_tag != encrypt('tag', passwd, cur_note.value.encryption_salt)) {
    return false
  }
  return true
};

export async function deletecurrent(router) {
  try {
    await ElMessageBox.confirm('This will delete the note, are you sure?', 'Warning', {
      confirmButtonText: 'OK',
      cancelButtonText: 'Cancel',
      type: 'warning'
    });

    const res = await fetch(import.meta.env.VITE_API_BASE + 'api/v1/notes/' + cur_note.value.slug, {
      method: 'DELETE',
      headers: {
        'X-Encryption-Tag': cur_note.value.encryption_verification_tag ? cur_note.value.encryption_verification_tag : '',
      },
    });


    if (!res.ok) {
      ElMessage.error('Failed to delete the note');
      return;
    }

    ElMessage.success('Note deleted successfully');
    cur_note.value = {
      title: '',
      body: '',
      created_at: 0,
      updated_at: 0,
      encryption_tag: '',
      encryption_salt: '',
      is_encrypted: false,
      slug: '----',
      passwd: '',
    }
    router.push({ name: 'home' });
  } catch (error) { }
}

export const sharecurruent = () => {
  const url = window.location.origin + '/' + cur_note.value.slug
  ElMessageBox.confirm(
    'Would you like to share the password?',
    'Confirm',
    {
      distinguishCancelAndClose: true,
      confirmButtonText: 'Yes',
      cancelButtonText: 'No',
    }
  )
    .then(() => {
      navigator.clipboard.writeText(url + '#' + cur_note.value.passwd).then(() => {
        ElMessage.success('Copied to clipboard')
      }).catch(() => {
        ElMessage.error('Failed to copy to clipboard')
      })
    })
    .catch((action) => {
      if ('cancel' != action) return;
      navigator.clipboard.writeText(url).then(() => {
        ElMessage.success('Copied to clipboard')
      }).catch(() => {
        ElMessage.error('Failed to copy to clipboard')
      })
    })
}

export const updatecurrent = async (router, edited) => {
  var edited_encrypted = undefined
  if (cur_note.value.is_encrypted) {
    edited_encrypted = {
      title: encrypt(edited.title, cur_note.value.passwd, cur_note.value.encryption_salt),
      body: encrypt(edited.body, cur_note.value.passwd, cur_note.value.encryption_salt),
    }
  }
  const res = await fetch(import.meta.env.VITE_API_BASE + 'api/v1/notes/' + cur_note.value.slug, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
      'X-Encryption-Tag': cur_note.value.encryption_verification_tag ? cur_note.value.encryption_verification_tag : '',
    },
    body: JSON.stringify(edited_encrypted ? edited_encrypted : edited),
  });

  if (!res.ok) {
    ElMessage.error('Failed to update the note');
    return;
  }

  const resjson = await res.json();

  ElMessage.success('Note updated successfully');
  cur_note.value = { ...cur_note.value, ...resjson }
  cur_note.value.decrypted = false
  router.push({ name: 'view', params: { slug: cur_note.value.slug } })
}

export const create_note = async (note, router) => {
  const res = await fetch(import.meta.env.VITE_API_BASE + 'api/v1/notes', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(note)
  });
    
  if (!res.ok) {
    ElMessage.error('Failed to create the note');
    return;
  }
  cur_note.value = await res.json();
  cur_note.value.passwd = ''
  router.push({ name: 'view', params: { 'slug': cur_note.value.slug } })
  return res;
}

export const decrypt_note = (passwd) => {
  if (passwd)
    cur_note.value.passwd = passwd;
  cur_note.value.decrypted = true;
  cur_note.value.title = decrypt(cur_note.value.title, cur_note.value.passwd, cur_note.value.encryption_salt);
  cur_note.value.body = decrypt(cur_note.value.body, cur_note.value.passwd, cur_note.value.encryption_salt);
  cur_note.value.encryption_verification_tag = encrypt('verification', cur_note.value.passwd, cur_note.value.encryption_salt);
}
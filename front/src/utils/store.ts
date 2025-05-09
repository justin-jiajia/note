import { ref } from 'vue'
import { encrypt, decrypt } from './crypto.js';
import { format_date } from './date.js';

interface Note {
  title: string;
  body: string;
  created_at: number;
  updated_at: number;
  encryption_tag: string;
  encryption_salt: string;
  encryption_verification_tag: string;
  decrypted: boolean;
  is_encrypted: boolean;
  slug: string;
  passwd: string;
  histories: Array<{
    title: string;
    body: string;
    created_at: number;
  }>;
}

const init_note: Note = {
  title: '',
  body: '',
  created_at: 0,
  updated_at: 0,
  encryption_tag: '',
  encryption_salt: '',
  encryption_verification_tag: '',
  decrypted: false,
  is_encrypted: false,
  slug: '----',
  passwd: '',
  histories: [],
}

export const cur_note = ref(init_note)

export const fetch_note = async (slug: string, force: boolean) => {
  if (slug == cur_note.value.slug && !force) {
    return
  }
  const res = await fetch(import.meta.env.VITE_API_BASE + 'api/v1/notes/' + slug);
  const resjson = await res.json();
  for (let cur_history in resjson.histories) {
    console.log(cur_history)
    console.log(resjson.histories[cur_history])
    resjson.histories[cur_history].created_at_text = format_date(resjson.histories[cur_history].created_at);
  }
  if (!res.ok) {
    ElMessage.error(resjson.error);
    return;
  }
  cur_note.value = { ...cur_note.value, ...(resjson as Note) }
}

export const verfiy_passwd = (passwd: string) => {
  if (cur_note.value.encryption_tag != encrypt('tag', passwd, cur_note.value.encryption_salt)) {
    return false
  }
  return true
};

export async function deletecurrent(router: any) {
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
    cur_note.value = init_note
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

export const updatecurrent = async (router: any, edited: any) => {
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
  cur_note.value = { ...cur_note.value, ...(resjson as Note) }
  cur_note.value.decrypted = false
  router.push({ name: 'view', params: { slug: cur_note.value.slug } })
}

export const create_note = async (note: any, router: any) => {
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
  cur_note.value = await res.json() as Note;
  cur_note.value.passwd = ''
  router.push({ name: 'view', params: { 'slug': cur_note.value.slug } })
  return res;
}

export const decrypt_note = (passwd?: string) => {
  if (!cur_note.value.is_encrypted)
    return
  if (passwd)
    cur_note.value.passwd = passwd;
  cur_note.value.decrypted = true;
  cur_note.value.title = decrypt(cur_note.value.title, cur_note.value.passwd, cur_note.value.encryption_salt);
  cur_note.value.body = decrypt(cur_note.value.body, cur_note.value.passwd, cur_note.value.encryption_salt);
  cur_note.value.encryption_verification_tag = encrypt('verification', cur_note.value.passwd, cur_note.value.encryption_salt);
  for (let cur_history in cur_note.value.histories) {
    cur_note.value.histories[cur_history].title = decrypt(cur_note.value.histories[cur_history].title, cur_note.value.passwd, cur_note.value.encryption_salt);
    cur_note.value.histories[cur_history].body = decrypt(cur_note.value.histories[cur_history].body, cur_note.value.passwd, cur_note.value.encryption_salt);
  }
}

export const copy_source = () => {
  navigator.clipboard.writeText(cur_note.value.body).then(() => {
    ElMessage.success('Copied to clipboard')
  }).catch(() => {
    ElMessage.error('Failed to copy to clipboard')
  })
}

export const download_source = () => {
  //export a "title.md" file
  var element = document.createElement('a');
  element.setAttribute('href', 'data:text/plain;charset=utf-8,' + encodeURIComponent(cur_note.value.body));
  element.setAttribute('download', cur_note.value.title + '.md');
  element.style.display = 'none';
  document.body.appendChild(element);
  element.click();
  document.body.removeChild(element);
}

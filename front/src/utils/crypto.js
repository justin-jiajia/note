import pbkdf2 from 'pbkdf2';
import aesjs from 'aes-js';

export const encrypt = (data, key, salt) => {
  const derivedKey = pbkdf2.pbkdf2Sync(key, salt, 1, 256 / 8, 'sha512');
  const aesCtr = new aesjs.ModeOfOperation.ctr(derivedKey, new aesjs.Counter(5));
  const encryptedBytes = aesCtr.encrypt(aesjs.utils.utf8.toBytes(data));
  return aesjs.utils.hex.fromBytes(encryptedBytes);
}

export const decrypt = (data, key, salt) => {
  const derivedKey = pbkdf2.pbkdf2Sync(key, salt, 1, 256 / 8, 'sha512');
  const aesCtr = new aesjs.ModeOfOperation.ctr(derivedKey, new aesjs.Counter(5));
  const decryptedBytes = aesCtr.decrypt(aesjs.utils.hex.toBytes(data));
  return aesjs.utils.utf8.fromBytes(decryptedBytes);
}
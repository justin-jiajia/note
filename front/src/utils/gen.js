// Only a-z, 64bits
export const generateSalt = () => {
  const charset = 'abcdefghijklmnopqrstuvwxyz';
  let salt = '';
  for (let i = 0; i < 8; i++) {
    salt += charset.charAt(Math.floor(Math.random() * charset.length));
  }
  return salt;
}
// Only a-z, A-Z, 0-9
export const generatePasswd = () => {
  const charset = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789';
  let passwd = '';
  for (let i = 0; i < 16; i++) {
    passwd += charset.charAt(Math.floor(Math.random() * charset.length));
  }
  return passwd;
}
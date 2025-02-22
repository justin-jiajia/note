export const format_date = (timestamp) => {
  const date = new Date(timestamp * 1000);
  return date.toLocaleString();
}
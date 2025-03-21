export const format_date = (timestamp: number): string => {
  const date = new Date(timestamp * 1000);
  return date.toLocaleString();
}

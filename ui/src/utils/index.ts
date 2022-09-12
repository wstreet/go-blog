import days from "dayjs";
export function getLocal(key: string) {
  return localStorage[key];
}

export function setLocal(key: string, value: string) {
  localStorage[key] = value;
}

export function formatDate(value: string): string {
  if (!value) {
    return "";
  }
  return days(value).format("YYYY-MM-DD");
}


export function formatTime(value: string): string {
  if (!value) {
    return "";
  }
  return days(value).format("YYYY-MM-DD HH:mm:ss");
}

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

export function getUrlParam(name: string): string {
  const reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)");
  const r = window.location.search.substring(1).match(reg);
  if (r != null) {
    return decodeURIComponent(r[2]);
  }
  return "";
}
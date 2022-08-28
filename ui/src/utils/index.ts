export function getLocal(key: string) {
  return localStorage[key];
}

export function setLocal(key: string, value: string) {
  localStorage[key] = value;
}
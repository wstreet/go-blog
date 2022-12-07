
import Storage from 'vue-ls'

const { ls } = Storage.useStorage({
  namespace: 'blog__', // key prefix
  name: 'ls', // name variable Vue.[ls] or this.[$ls],
  storage: 'local', // storage name session, local, memory
})

export default ls
export default {

  log: false,
  events: {},

  on(name, cb) {
    if (typeof this.events[name] === 'undefined') {
      this.events[name] = []
    }
    this.events[name].push(cb)
  },

  emit(name, data) {
    if (this.log === true) {
      console.info(`'${name}'`, data)
    }
    for (const cb of this.events[name] || []) {
      cb(data)
    }
  }
}

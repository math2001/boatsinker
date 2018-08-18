import em from './em.js'

// this is a very basic componements who let's the client know the state of the
// game => is it their turn for example. It's all based on the events that the
// app is triggering.

export default {

  init() {
    this.box = document.querySelector('#statebox')
    this.name = '<no name thats a bug>'
    this.setMessage(`Please enter your name`)
    em.on('got.name', name => {
      this.name = name
      this.setMessage(`Hello ${name}! Just waiting for an other player to join now...`)
    })
    em.on('connection.msg', msg => {
      if (msg.kind === 'hit') {
        this.setMessage(`Try to hit a boat!`)
      } else if (msg.kind === 'wait') {
        this.setMessage(`Watchout! They're going to hit!`)
      } else {
        this.setMessage(`Not quite sure what's happening right now...`)
      }
    })
    em.on('got.accepted', e => {
      this.setMessage(`Set up your boats!`)
    })
    em.on('boat.allsetup', e => {
      this.setMessage(`Hit send!`)
    })
    em.on('boat.sendsetup', e => {
      this.setMessage(`Alright! Waiting for the other player now...`)
    })
  },

  setMessage(msg) {
    this.box.textContent = msg
  },

}

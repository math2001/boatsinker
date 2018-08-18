import em from './em.js'
import GetName from './getname.js'
import Game from './game.js'
import State from './state.js'

window.addEventListener('load', e => {

  const root = document.getElementById('root')
  GetName.init(root)
  Game.init(root)
  State.init()

  const ws = new WebSocket(`ws://${location.host}/ws`)
  ws.onmessage = msg => {
    em.emit('connection.msg', JSON.parse(msg.data))
  }

  em.on('connection.send', msg => {
    ws.send(JSON.stringify(msg))
  })
  em.on('got.name', name => {
    em.emit('connection.send', {
      kind: 'request',
      name: name
    })
  })

  em.on('connection.msg', msg => {
    if (msg.kind == 'state change' && msg.state == 'setup') {
      em.emit('got.accepted', {
        width: msg.width,
        height: msg.height,
        boatsizes: msg.boatsizes
      })
    }
  })

  GetName.setDOM()
})

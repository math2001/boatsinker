import em from './em.js'
import GetName from './getname.js'
import Toolbar from './toolbar.js'
import Game from './game.js'

GetName.init()
Toolbar.init()
Game.init()

window.addEventListener('load', e => {

  // const ws = new WebSocket(`ws://${location.host}/ws`)
  // ws.onmessage = msg => {
  //   em.emit('connection.msg', JSON.parse(msg.data))
  // }

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
      em.emit('got.setup', { size: msg.size, boatsizes: msg.boat_sizes })
    }
  })

  // debug
  // em.emit('get.name')
  em.emit('got.setup', { width: 10, height: 10, boatsizes: { 5: 1, 4: 1, 3: 2, 2: 1 } })
})

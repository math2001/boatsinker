import em from './em.js'
import GetName from './getname.js'
import Game from './game.js'

GetName.init()
Game.init()

const ws = new WebSocket(`ws://${location.host}/ws`)

em.on('connection.send', msg => {
  ws.send(JSON.stringify(msg))
})

ws.onmessage = msg => {
  em.emit('connection.msg', JSON.parse(msg.data))
}

em.on('got.name', name => {
  em.emit('connection.send', {
    kind: 'request',
    name: name
  })
})

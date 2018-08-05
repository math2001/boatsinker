import Board from './board.js'
import em from './em.js'

export default {

  init() {
    em.on('connection.msg', msg => {
      if (msg.kind === 'state change' && msg.state === 'setup') {
        this.own = new Board()
        this.other = new Board()
      }
    })
  }

}

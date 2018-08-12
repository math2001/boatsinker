import Board from './board.js'
import em from './em.js'

export default {

  init() {
    em.on('got.setup', this.setup.bind(this))
  },

  setup(opts) {
    const boards = document.querySelector('#boards')
    this.own = new Board(boards,   {width: opts.width, height: opts.height, own: true})
    this.other = new Board(boards, {width: opts.width, height: opts.height, own: false})
  }

}

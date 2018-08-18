import Board from './board.js'
import Toolbar from './toolbar.js'
import em from './em.js'

export default {

  init(root) {
    em.on('got.accepted', this.setup.bind(this))
    this.root = root
  },

  setup(opts) {
    Toolbar.init(root, opts.boatsizes)
    const boards = document.createElement('article')
    root.appendChild(boards)
    this.own = new Board(boards,   {width: opts.width, height: opts.height, own: true})
    // this.other = new Board(boards, {width: opts.width, height: opts.height, own: false})
  }

}

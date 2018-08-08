import Board from './board.js'
import em from './em.js'

const boat_sizes = {
  // size: count
  2: 1,
  3: 2,
  4: 1,
  5: 1
}

const toolbar = {
  init() {
    this.toolbar = document.querySelector('#toolbar')
    this.boats = []
    em.on("got.setup", conf => {
      for (const size in conf.boatsizes) {
        for (let i = 0; i < boat_sizes[size]; i++) {
          this.newBoat(size, 0)
        }
      }
    })
  },
  newBoat(size, rotation) {
    const boat = document.createElement('div')
    boat.setAttribute('data-size', size)
    boat.setAttribute('data-rotation', rotation)
    this.toolbar.appendChild(boat)
    boat.addEventListener('click', this.click(boat))
    this.boats.push(boat)
  },

  click(boat) {
    return e => {

    }
  }
}

export default {

  init() {
    toolbar.init()
    em.on('got.setup', this.setup.bind(this))
  },

  setup() {
    const boards = document.querySelector('#boards')
    this.own = new Board(boards)
    this.other = new Board(boards)
  }

}

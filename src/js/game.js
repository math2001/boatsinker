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
    this.currentBoat = null
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
    boat.style.setProperty('--size', size)
    boat.setAttribute('data-size', size)
    boat.textContent = size
    boat.setAttribute('data-rotation', rotation)
    boat.classList.add('boat')
    this.toolbar.appendChild(boat)
    boat.addEventListener('contextmenu', this.rotate.bind(this))
    boat.addEventListener('click', this.select.bind(this))
    this.boats.push(boat)
  },

  rotate(e) {
    e.preventDefault()
    if (e.target.classList.contains('selected')) {
      alert("Please unselect this boat before rotating it.")
      return
    }
    e.target.setAttribute('data-rotation',
      e.target.getAttribute('data-rotation') === '0' ? '1' : '0'
    )
  },

  select(e) {
    if (this.currentBoat) {
      em.emit('boat.unselect', {
        size: parseInt(this.currentBoat.getAttribute('data-size')),
        rotation: parseInt(this.currentBoat.getAttribute('data-rotation'))
      })
      this.currentBoat.classList.remove('selected')
      if (this.currentBoat === e.target) {
        // deselect
        this.currentBoat = null
      } else {
        // clicked on an other boat (e.target)
        this.currentBoat = e.target
      }
    } else {
      // select new boat
      this.currentBoat = e.target
    }

    if (this.currentBoat) {
      this.currentBoat.classList.add('selected')
      em.emit('boat.select', {
        size: parseInt(this.currentBoat.getAttribute('data-size')),
        rotation: parseInt(this.currentBoat.getAttribute('data-rotation'))
      })
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
    this.own = new Board(boards, {own: true})
    this.other = new Board(boards, {own: false})
  }

}

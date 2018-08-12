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
    this.draggedBoat = null
    this.offset = null
    em.on("got.setup", conf => {
      for (const size in conf.boatsizes) {
        for (let i = 0; i < boat_sizes[size]; i++) {
          this.newBoat(size, 0)
        }
      }
    })
    document.addEventListener('mousemove', this.drag.bind(this))
    document.addEventListener('mouseup', this.dragend.bind(this))
  },

  newBoat(size, rotation) {
    const boat = document.createElement('div')
    boat.style.setProperty('--size', size)
    boat.setAttribute('data-size', size)
    boat.textContent = size
    boat.setAttribute('data-rotation', rotation)
    boat.classList.add('boat')
    this.toolbar.appendChild(boat)
    boat.addEventListener('contextmenu', this.click.bind(this))
    boat.addEventListener('mousedown', this.dragstart.bind(this))
    this.boats.push(boat)
  },

  click(e) {
    e.preventDefault()
    e.target.setAttribute('data-rotation',
      e.target.getAttribute('data-rotation') === '0' ? '1' : '0'
    )
  },

  dragstart(e) {
    this.draggedBoat = e.target
    this.draggedBoat.style.position = 'absolute'
    this.draggedBoat.style.opacity = .4
    const r = this.draggedBoat.getBoundingClientRect()
    this.offset = {
      x: e.clientX - r.left,
      y: e.clientY - r.top,
    }
    console.log(this.offset)
  },

  drag(e) {
    if (this.draggedBoat && this.draggedBoat !== null) {
      this.draggedBoat.style.left = e.clientX - this.offset.x + 'px'
      this.draggedBoat.style.top = e.clientY - this.offset.y + 'px'
    }
  },

  dragend(e) {
    if (!this.draggedBoat) {
      return
    }
    this.draggedBoat.style.position = 'static'
    this.draggedBoat.style.opacity = 1
    this.draggedBoat.style.left = null
    this.draggedBoat.style.top = null
    this.draggedBoat = null
    this.offset = null
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

import em from './em.js'

export default {
  init(root, boatsizes) {
    this.elements = { root }

    this.elements.toolbar = document.createElement('article')
    this.elements.toolbar.classList.add('toolbar')
    this.elements.root.appendChild(this.elements.toolbar)

    this.elements.sendsetup = document.createElement('button')
    this.elements.sendsetup.type = 'button'
    this.elements.sendsetup.textContent = 'Send set up'
    this.elements.sendsetup.disabled = true
    this.elements.toolbar.appendChild(this.elements.sendsetup)

    this.elements.sendsetup.addEventListener('click', e => {
      // the toolbar (this componenet) doesn't know where the boats have been
      // placed (it could, but it's not its job). It's the board who listens
      // for this even and then actually send the positions to the server
      em.emit('boat.sendsetup', null)
      e.target.parentElement.removeChild(e.target)
    })

    this.boats = []
    this.currentBoat = null
    for (const size in boatsizes) {
      for (let i = 0; i < boatsizes[size]; i++) {
        this.newBoat(size, 0)
      }
    }

    em.on("boat.placed", this.boatplaced.bind(this))
  },

  newBoat(size, rotation) {
    const boat = document.createElement('div')
    boat.style.setProperty('--size', size)
    boat.setAttribute('data-size', size)
    boat.textContent = size
    boat.setAttribute('data-rotation', rotation)
    boat.classList.add('boat')
    this.elements.toolbar.appendChild(boat)
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
      if (this.currentBoat === e.target) { // deselect
        this.currentBoat = null
      } else { // clicked on an other boat (e.target)
        this.currentBoat = e.target
      }
    } else { // select new boat
      this.currentBoat = e.target
    }

    if (this.currentBoat) {
      this.currentBoat.classList.add('selected')
      em.emit('boat.select', {
        size: parseInt(this.currentBoat.getAttribute('data-size')),
        rotation: parseInt(this.currentBoat.getAttribute('data-rotation'))
      })
    }
  },

  boatplaced(e) {
    if (!this.currentBoat) {
      throw new Error(`How the f*** did a boat got placed? I (toolbar) got nothing!`)
    }
    const size = parseInt(this.currentBoat.getAttribute('data-size'))
    const rotation = parseInt(this.currentBoat.getAttribute('data-rotation'))
    if (size !== e.size || rotation !== e.rotation) {
      throw new Error(`The boat that got place doesn't match up with the boat \
      I (toolbar) got! ${size} !== ${e.size} or ${rotation} !== ${e.rotation}`)
    }
    this.currentBoat.classList.add('placed')
    this.currentBoat = null
    if (this.boats.every(boat => boat.classList.contains('placed'))) {
      this.elements.sendsetup.disabled = false
    }
  }
}


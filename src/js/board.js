import em from './em.js'

function isInsideBoard(boat, board) {
  return (boat.rotation === 0 && boat.x + boat.size <= board.width)
    || (boat.rotation === 1 && boat.y + boat.size <= board.height)
}

function collides(boat, boats) {
  // this functions is absolutely discusting. I need to improve it.
  const occupied_cells = []
  for (const b of boats) {
    for (let i = 0; i < b.size; i++) {
      if (b.rotation == 0) {
        occupied_cells.push([b.x + i, b.y])
      } else {
        occupied_cells.push([b.x, b.y + i])
      }
    }
  }
  const targetcells = []
  for (let i = 0; i < boat.size; i++) {
    if (boat.rotation == 0) {
      targetcells.push([boat.x + i, boat.y])
    } else {
      targetcells.push([boat.x, boat.y + i])
    }
  }
  return occupied_cells.some(c => {
    for (const t of targetcells) {
      if (c[0] === t[0] && c[1] === t[1]) {
        console.log('collision')
        return true
      }
    }
  })
}

export default class Board {

  constructor(parent, opts={}) {
    this.opts = opts
    this.boat = null
    // boats that have been placed. A boat is { size: <int>, rotation: <int>,
    // x: <int>, y: <int> }
    this.boats = []
    this.highlightedcellsindex = []
    if (this.opts.own) {
      em.on('boat.select', boat => this.boat = boat)
      em.on('boat.unselect', boat => this.boat = null)
    }

    this.board = document.createElement('div')
    this.board.classList.add('board')
    this.cells = []
    let cell;
    for (let i = 0; i < this.opts.width * this.opts.height; i++) {
      cell = document.createElement('div')
      cell.classList.add('cell')
      if (this.opts.own) {
        cell.addEventListener('mouseenter', this.mouseenter.bind(this))
        cell.addEventListener('mouseleave', this.mouseleave.bind(this))
      }
      cell.addEventListener('click', this.click.bind(this))
      this.cells.push(cell)
      this.board.appendChild(cell)
    }

    parent.appendChild(this.board)
  }

  cellfromxy(x, y) {
    return this.cells[y * this.opts.height + x]
  }

  celltoxy(cell) {
    const id = this.cells.findIndex(c => c === cell)
    return [id % this.opts.width, (id - id % this.opts.height)/this.opts.height]
  }

  mouseleave(e) {
    for (const coor of this.highlightedcellsindex) {
      const cell = this.cellfromxy(...coor)
      cell.classList.remove('active')
      // the class error can be added if there was a boat already on the cell
      cell.classList.remove('error')
    }
    this.highlightedcellsindex = []
    this.board.classList.remove('error')
  }

  mouseenter(e) {
    if (this.boat) {
      const [x, y] = this.celltoxy(e.target)
      if (!isInsideBoard(Object.assign({ x, y }, this.boat), this.opts)) {
        this.board.classList.add('error')
        return
      }
      for (let i = 0; i < this.boat.size; i++) {
        if (this.boat.rotation === 0) {
          // horizontal
          this.highlightedcellsindex.push([x + i, y])
        } else if (this.boat.rotation === 1) {
          // horizontal
          this.highlightedcellsindex.push([x, y + i])
        }
      }
      for (const coor of this.highlightedcellsindex) {
        const cell = this.cellfromxy(...coor)
        if (cell.classList.contains('boat')) { // there is a boat on that cell
          cell.classList.add('error')
        } else {
          cell.classList.add('active')
        }
      }
    }
  }

  click(e) {
    // checks wether the position (e.target) is valid for the boat (this.boat).
    // If it is, trigger boat.placed with position and information of the boat
    if (!this.boat) return
    const [x, y] = this.celltoxy(e.target)
    const boat = Object.assign({ x, y }, this.boat)
    if (!isInsideBoard(boat, this.opts)) {
      return alert("The boat has to be inside the board.")
    }
    if (collides(boat, this.boats)) {
      return alert("Colliding with a boat! Can't put it there.")
    }
    for (const coor of this.highlightedcellsindex) {
      const cell = this.cellfromxy(...coor)
      cell.classList.remove('active')
      cell.classList.add('boat')
    }
    em.emit("boat.placed", boat)
    this.boats.push(boat)
    this.boat = null
    this.highlightedcellsindex = []
  }
}

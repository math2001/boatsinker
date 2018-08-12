import em from './em.js'

function valid(x, y, boat, board) {
  // note that this doesn't check collision between boats (yet)
  return (boat.rotation === 0 && x + boat.size <= board.width)
    || (boat.rotation === 1 && y + boat.size <= board.height)
}

export default class Board {

  constructor(parent, opts={}) {
    this.opts = opts
    this.boat = null
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
      this.cellfromxy(...coor).classList.remove('active')
    }
    this.highlightedcellsindex = []
    this.board.classList.remove('error')
  }

  mouseenter(e) {
    if (this.boat) {
      const [x, y] = this.celltoxy(e.target)
      if (!valid(x, y, this.boat, this.opts)) {
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
        this.cellfromxy(...coor).classList.add('active')
      }
    }
  }

  click(e) {

  }

}

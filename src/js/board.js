import em from './em.js'

export default class Board {

  constructor(parent, opts={}) {
    this.boat = null
    this.highlightedcellsindex = []
    if (opts.own) {
      em.on('boat.select', boat => this.boat = boat)
      em.on('boat.unselect', boat => this.boat = null)
    }

    let board = document.createElement('div')
    board.classList.add('board')
    this.cells = []
    let cell;
    for (let i = 0; i < 100; i++) {
      cell = document.createElement('div')
      cell.classList.add('cell')
      if (opts.own) {
        cell.addEventListener('mouseenter', this.mouseenter.bind(this))
        cell.addEventListener('mouseleave', this.mouseleave.bind(this))
      }
      cell.addEventListener('click', this.click.bind(this))
      this.cells.push(cell)
      board.appendChild(cell)
    }

    parent.appendChild(board)
  }

  cellfromxy(x, y) {
    return this.cells[y * 10 + x]
  }

  celltoxy(cell) {
    const id = this.cells.findIndex(c => c === cell)
    return [id % 10, (id - id % 10)/10]
  }

  mouseleave(e) {
    for (const coor of this.highlightedcellsindex) {
      this.cellfromxy(...coor).classList.remove('active')
    }
    this.highlightedcellsindex = []
  }

  mouseenter(e) {
    if (this.boat) {
      const [x, y] = this.celltoxy(e.target)
      for (let i = 0; i < this.boat.size; i++) {
        if (this.boat.rotation === 0) {
          // horizontal
          this.highlightedcellsindex.push([x + i, y])
        } else if (this.boat.rotation === 1) {
          // horizontal
          this.highlightedcellsindex.push([x, y + i])
        } else {
          throw new Error(`Invalid boat rotation: should be 0 or 1, got ${this.boat.rotation}`)
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

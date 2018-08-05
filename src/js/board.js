export default class Board {

  constructor() {
    let board = document.createElement('div')
    board.classList.add('board')
    this.cells = []
    let cell;
    for (let i = 0; i < 100; i++) {
      cell = document.createElement('div')
      cell.classList.add('cell')
      this.cells.push(cell)
      board.appendChild(cell)
    }

    document.body.appendChild(board)
  }

  fromxy(x, y) {
    this.cells[y * 10 + x]
  }

  idtoxy(id) {
    return [id - id % 10, id % 10]
  }

}

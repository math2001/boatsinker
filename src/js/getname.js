import em from './em.js'

export default {

  init(root) {
    this.root = root
  },

  setDOM() {
    const form = document.createElement('form')
    const input = document.createElement('input')
    const submit = document.createElement('input')

    input.type = 'text'
    input.placeholder = "Your pseudo"

    input.setAttribute('minlength', 3)
    input.setAttribute('maxlength', 20)

    submit.type = 'submit'
    submit.value = 'Go!'

    form.appendChild(input)
    form.appendChild(submit)
    this.root.appendChild(form)

    input.focus()

    form.addEventListener('submit', e => {
      e.preventDefault()
      em.emit('connection.send', {
        kind: 'request',
        name: input.value
      })
      form.parentElement.removeChild(form)
    })
  }

}

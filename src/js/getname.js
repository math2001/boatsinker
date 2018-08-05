import em from './em.js'

export default {

  init() {
    this.manageDOM()
  },

  manageDOM() {
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
    document.body.appendChild(form)
    form.addEventListener('submit', e => {
      e.preventDefault()
      em.emit("got.name", input.value)
      form.parentElement.removeChild(form)
    })
  }

}

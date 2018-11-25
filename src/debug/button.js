import m from 'mithril'
import Button, { BUTTON_STYLE } from '../component/Button'
import IconButton from '../component/IconButton'

export default {
  view () {
    return m('div', [
      m(Button, {
        title: 'あああ',
        style: BUTTON_STYLE.UNELEVATED,
        dense: true,
        disabled: false,
        icon: 'movie',
        onclick: () => {
          console.log('click')
        }
      }),
      m('br'),
      m(IconButton, {})
    ])
  }
}

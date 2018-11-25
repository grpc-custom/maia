import m from 'mithril'
import Radio from '../component/Radio'

export default {
  view () {
    return m('div', [
      m(Radio, {
        id: 'radio'
      })
    ])
  }
}

import m from 'mithril'
import stream from 'mithril/stream'
import Select from '../component/Select'

const value = stream('')

export default {
  view () {
    return m('div', [
      m(Select, {
        title: 'ABC',
        value,
        options: [
          {label: 'AAA', value: 'aaa'},
          {label: 'BBB', value: 'bbb'},
          {label: 'CCC', value: 'ccc'},
        ]
      }),
      m('button', {
        onclick: () => {
          console.log('click', value())
        }
      }, 'ok')
    ])
  }
}

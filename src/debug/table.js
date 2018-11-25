import m from 'mithril'
import Table from '../component/Table'

export default {
  view() {
    return m('div', [
      m(Table, { data: 'data' })
    ])
  }
}

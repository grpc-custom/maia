import m from 'mithril'
import stream from 'mithril/stream'
import { MDCSelect } from '@material/select'
import './select.scss'

export default {
  oncreate (vnode) {
    const select = new MDCSelect(vnode.dom)
    select.listen('MDCSelect:change', () => {
      vnode.attrs.value(select.value)
    })
  },
  view (vnode) {
    const {options, title} = vnode.attrs || {}
    return m('div', {class: 'mdc-select'}, [
      m('i', {class: 'mdc-select__dropdown-icon'}),
      m('select', {class: 'mdc-select__native-control'}, options.map(({label, value}) => m('option', {value}, label))),
      m('label', {class: 'mdc-floating-label'}, title || ''),
      m('div', {class: 'mdc-line-ripple'})
    ])
  }
}

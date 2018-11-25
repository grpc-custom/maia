import m from 'mithril'
import './radio.scss'

export default {
  view (vnode) {
    const { id } = vnode.attrs || {}
    return m('div', {
      class: 'mdc-form-field'
    }, [
      m('div', { class: 'mdc-radio' }, [
        m('input', {
          class: 'mdc-radio__native-control',
          type: 'radio',
          name: id
        }),
        m('div', { class: 'mdc-radio__background' }, [
          m('div', { class: 'mdc-radio__outer-circle' }),
          m('div', { class: 'mdc-radio__inner-circle' })
        ])
      ]),
      m('label', 'Radio 1')
    ])
  }
}

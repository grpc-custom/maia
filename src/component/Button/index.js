import m from 'mithril'
import { MDCRipple } from '@material/ripple'
import './button.scss'

export const BUTTON_STYLE = {
  DEFAULT: '',
  OUTLINE: 'outline',
  RAISED: 'raised',
  UNELEVATED: 'unelevated'
}

export default {
  oncreate (vnode) {
    MDCRipple.attachTo(vnode.dom)
  },
  view (vnode) {
    const {title, style, dense, disabled, icon, onclick} = vnode.attrs || {}
    let clazz = ['mdc-button']
    switch (style) {
      case BUTTON_STYLE.OUTLINE:
        clazz.push('mdc-button--outlined')
        break
      case BUTTON_STYLE.RAISED:
        clazz.push('mdc-button--raised')
        break
      case BUTTON_STYLE.UNELEVATED:
        clazz.push('mdc-button--unelevated')
        break
    }
    if (dense) {
      clazz.push('mdc-button--dense')
    }
    let iconEl = null
    if (icon) {
      iconEl = m('i', {
        class: 'material-icons mdc-button__icon',
        'aria-hidden': true
      }, icon)
    }
    return m('button', {
      class: clazz.join(' '),
      disabled,
      onclick
    }, iconEl, title)
  }
}

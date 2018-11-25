import m from 'mithril'
import { MDCRipple } from '@material/ripple'
import './icon_button.scss'

export default {
  oncreate (vnode) {
    const icon = new MDCRipple(vnode.dom)
    icon.unbounded = true
  },
  view (vnode) {
    return m('button', {
      class: 'mdc-icon-button material-icons'
    }, 'favorite')
  }
}

import m from 'mithril'
import stream from 'mithril/stream'

import TextField, { TEXT_STYLE } from '../../component/TextField'
import Button, { BUTTON_STYLE } from '../../component/Button'

function getTextFieldType (attrs = {}, type) {
  if (attrs.type) {
    return attrs.type
  }
  type = Array.isArray(type) ? type[0] : type
  return type === 'integer' ? 'number' : 'text'
}

function getTextFieldStyle (attrs = {}) {
  if (attrs.style) {
    return attrs.style
  }
  return TEXT_STYLE.OUTLINE
}

export default {
  oninit (vnode) {
    const { schema } = vnode.attrs || {}
    this.values = {}
    for (let key in schema.properties) {
      this.values[key] = stream()
    }
  },
  view (vnode) {
    const { schema } = vnode.attrs || {}
    const fields = []
    for (let key in schema.properties) {
      const property = schema.properties[key]
      fields.push(m(TextField, {
        title: property.title,
        type: getTextFieldType(property.attrs, property.type),
        value: this.values[key],
        dense: true,
        schema: {
          type: Array.isArray(property.type) ? property.type[0] : property.type
        },
        style: getTextFieldStyle(property.attrs)
      }))
    }

    fields.push(m(Button, {
      title: 'ok',
      style: BUTTON_STYLE.UNELEVATED,
      onclick: () => {
        console.log(JSON.stringify(this.values, null, 2))
      }
    }))
    return m('div', fields)
  }
}

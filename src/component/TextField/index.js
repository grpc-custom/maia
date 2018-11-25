import m from 'mithril'
import { MDCTextField } from '@material/textfield'
import './text_field.scss'

function formatter (schema, value) {
  switch (schema.type.toLowerCase()) {
    case 'integer': {
      const num = parseInt(value)
      return isNaN(num) ? 0 : num
    }
    case 'number': {
      if (/^[0-9]+\.[0]*$/.test(value)) {
        return value
      }
      const num = Number(value)
      return isNaN(num) ? 0 : num
    }
  }
  return value
}

function viewDefaultText ({title, type, value, schema, dense}) {
  const clazz = ['mdc-text-field']
  if (dense) {
    clazz.push('mdc-text-field--dense')
  }
  return m('div', {
    class: clazz.join(' ')
  }, [
    m('input', {
      type: type || 'text',
      class: 'mdc-text-field__input',
      value: value(),
      oninput: m.withAttr('value', (val) => value(formatter(schema, val)))
    }),
    m('label', {class: 'mdc-floating-label'}, title || ''),
    m('div', {class: 'mdc-line-ripple'})
  ])
}

function viewOutlineText ({title, type, value, schema, dense}) {
  const clazz = ['mdc-text-field', 'mdc-text-field--outlined']
  if (dense) {
    clazz.push('mdc-text-field--dense')
  }
  return m('div', {
    class: clazz.join(' ')
  }, [
    m('input', {
      type: type || 'text',
      class: 'mdc-text-field__input',
      value: value(),
      oninput: m.withAttr('value', (val) => value(formatter(schema, val)))
    }),
    m('label', {class: 'mdc-floating-label'}, title || ''),
    m('div', {class: 'mdc-notched-outline'}, [
      m('svg', m('path', {class: 'mdc-notched-outline__path'}))
    ]),
    m('div', {class: 'mdc-notched-outline__idle'})
  ])
}

function viewTextarea ({title, value, dense}) {
  const clazz = ['mdc-text-field', 'mdc-text-field--textarea']
  if (dense) {
    clazz.push('mdc-text-field--dense')
  }
  return m('div', {
    class: clazz.join(' ')
  }, [
    m('textarea', {
      class: 'mdc-text-field__input',
      rows: 3,
      value: value(),
      oninput: m.withAttr('value', value)
    }),
    m('label', {
      class: 'mdc-floating-label'
    }, title)
  ])
}

export const TEXT_STYLE = {
  TEXT: '',
  OUTLINE: 'outline',
  TEXTAREA: 'textarea'
}

export default {
  oncreate (vnode) {
    MDCTextField.attachTo(vnode.dom)
  },
  view (vnode) {
    const {title, type, schema, value, dense, style} = vnode.attrs || {}
    switch (style) {
      case TEXT_STYLE.OUTLINE:
        return viewOutlineText({title, type, schema, value, dense})
      case TEXT_STYLE.TEXTAREA:
        return viewTextarea({title, value, dense})
      default:
        return viewDefaultText({title, type, schema, value, dense})
    }
  }
}

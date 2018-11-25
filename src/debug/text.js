import m from 'mithril'
import stream from 'mithril/stream'
import TextField, { TEXT_STYLE } from '../component/TextField'

const values = {
  string: stream(),
  password: stream(),
  integer: stream(),
  number: stream(),
  boolean: stream()
}

export default {
  view() {
    return m('div', [
      m(TextField, {
        title: 'string',
        type: 'text',
        value: values.string,
        dense: true,
        schema: {
          type: 'string'
        },
        style: TEXT_STYLE.OUTLINE
      }),
      m(TextField, {
        title: 'password',
        type: 'password',
        value: values.password,
        dense: true,
        schema: {
          type: 'string'
        },
        style: TEXT_STYLE.OUTLINE
      }),
      m(TextField, {
        title: 'integer',
        type: 'number',
        value: values.integer,
        dense: true,
        schema: {
          type: 'integer'
        },
        style: TEXT_STYLE.OUTLINE
      }),
      m(TextField, {
        title: 'number',
        value: values.number,
        dense: true,
        schema: {
          type: 'number'
        },
        style: TEXT_STYLE.OUTLINE
      }),
      m('button', {
        onclick: () => {
          console.log('click', JSON.stringify(values, null, 2))
        }
      }, 'ok')
    ])
  }
}

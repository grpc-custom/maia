import m from 'mithril'

class Table {
  constructor () {

  }

  oninit (vnode) {
    const page = m.route.param('page')
    console.log('init ' + vnode.attrs.data + ' ' + page)
  }

  view (vnode) {
    return m('div', 'table ' + vnode.attrs.data)
  }
}

export default new Table()

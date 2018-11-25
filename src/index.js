import m from 'mithril'
import Text from './debug/text'
import Table from './debug/table'
import Select from './debug/select'
import Button from './debug/button'
import form from './debug/form'
import Radio from './debug/radio'

const Home = {
  view() {
    return 'Home!!'
  }
}

m.route(document.body, '/home', {
  '/home': Home,
  '/debug/text': Text,
  '/debug/table': Table,
  '/debug/select': Select,
  '/debug/button': Button,
  '/debug/radio': Radio,
  '/debug/form': form
})

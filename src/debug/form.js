import m from 'mithril'
import DynamicForm from '../module/DynamicForm'

export default {
  view () {
    return m('div', [
      m(DynamicForm, {schema:{title:'A registration form sample',description:'A simple form example.',type:'object',properties:{first_name:{type:'string',title:'First name'},last_name:{type:'string',title:'Last name'},age:{type:['integer','null'],title:'Age'},num:{type:'integer',title:'Num'},bio:{type:['number','null'],title:'Bio'},password:{type:'string',title:'Password',attrs:{type:'password'}},comment:{type:['string','null'],title:'comment',attrs:{style:'textarea'}}},required:['first_name','last_name']}})
    ])
  }
}

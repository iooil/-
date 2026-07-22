const {
  request
 }=require('../../utils/request')
 
 
 Page({
 
 data:{
 
 
 teacher:{},
 
 
 courses:[],
 
 
 reviews:[]
 
 
 },
 
 
 
 onLoad(options){
 
 
 this.teacherId=options.id
 
 
 this.loadTeacherDetail()
 
 
 },
 
 
 
 async loadTeacherDetail(){
 
 
 try{
 
 
 let res=await request({
 
 url:
 `/api/teachers/${this.teacherId}`
 
 })
 
 
 
 this.setData({
 
 teacher:res.data.teacher,
 
 courses:res.data.courses,
 
 reviews:res.data.reviews
 
 })
 
 
 
 }catch(e){
 
 console.log(e)
 
 }
 
 
 },
 
 
 
 evaluateTeacher(){
 
 
 wx.showToast({
 
 title:"评价功能开发中",
 
 icon:"none"
 
 })
 
 
 }
 
 
 })
const {
  request
 }=require('../../utils/request')
 
 
 Page({
 
 data:{
 
 courses:[],
 
 loading:false
 
 },
 
 
 onShow(){
 
 this.loadCourses()
 
 },
 
 
 
 async loadCourses(){
 
 
 this.setData({
 
 loading:true
 
 })
 
 
 try{
 
 
 let res=await request({
 
 url:"/api/my/courses"
 
 })
 
 console.log("我的课程数据:",res.data)
 
 this.setData({
 courses:res.data
 
 })
 
 
 
 }catch(e){
 
 console.log(e)
 
 
 }
 finally{
 
 
 this.setData({
 
 loading:false
 
 })
 
 
 }
 
 
 },
 
 
 
 
 dropCourse(e){
 
 
 let id=e.currentTarget.dataset.id
 
 
 
 wx.showModal({
 
 title:"提示",
 
 content:"确定退选该课程吗？",
 
 
 
 success:(res)=>{
 
 
 if(res.confirm){
 
 
 this.confirmDrop(id)
 
 
 }
 
 
 }
 
 
 })
 
 
 },
 
 
 
 async confirmDrop(id){
 
 
 try{
 
 
 await request({
 
 url:
 `/api/my/courses/${id}/drop`,
 
 method:"POST"
 
 })
 
 
 
 wx.showToast({
 
 title:"退课成功"
 
 })
 
 
 this.loadCourses()
 
 
 
 }catch(e){
 
 
 console.log(e)
 
 
 }
 
 
 
 },
 
 
 evaluateCourse(e){

  let id=e.currentTarget.dataset.id
  
  
  wx.navigateTo({
  
  url:
  `/pages/course-review/course-review?id=${id}`
  
  })
  
  },

  evaluateTeacher(e){

    let teacherId = e.currentTarget.dataset.teacher
  
    let courseId = e.currentTarget.dataset.course
  
  
    wx.navigateTo({
  
      url:
      `/pages/teacher-review/teacher-review?teacherId=${teacherId}&courseId=${courseId}`
  
    })
  
  }
 
 })
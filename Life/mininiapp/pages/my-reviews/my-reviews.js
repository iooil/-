const {
  request
 }=require('../../utils/request')
 
 
 Page({
 
 data:{
 
 courseReviews:[],
 
 teacherReviews:[]
 
 },
 
 
 
 onShow(){
 
 this.loadReviews()
 
 },
 
 
 
 async loadReviews(){
 
 
 try{
 
 
 let res=await request({
 
 url:"/api/my/reviews"
 
 })
 
 console.log("我的评价:",res)
 
 this.setData({
 
 courseReviews:
 res.data.course_reviews,
 
 
 teacherReviews:
 res.data.teacher_reviews
 
 
 })
 
 
 }catch(e){
 
 console.log(e)
 
 }
 
 
 }
 
 
 })
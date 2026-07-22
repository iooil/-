const {
  request
}=require('../../utils/request')


Page({

data:{

teacherId:0,

courseId:0,

score:5,

content:"",

stars:[1,2,3,4,5]

},



onLoad(options){

this.setData({

teacherId:Number(options.teacherId),

courseId:Number(options.courseId)

})


},



chooseScore(e){

this.setData({

score:Number(
e.currentTarget.dataset.score
)

})

},



inputContent(e){

this.setData({

content:e.detail.value

})

},



async submitReview(){


if(!this.data.content){


wx.showToast({

title:"请输入评价内容",

icon:"none"

})

return

}



try{


await request({

url:"/api/reviews/teacher",

method:"POST",

data:{


teacher_id:this.data.teacherId,

course_id:this.data.courseId,

score:this.data.score,

content:this.data.content


}


})


wx.showToast({

title:"评价成功",

icon:"success"

})



setTimeout(()=>{


wx.navigateBack()


},1000)



}catch(e){


console.log(e)


}


}


})
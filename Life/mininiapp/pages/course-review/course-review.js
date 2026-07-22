const {
  request
}=require('../../utils/request')


Page({

data:{


courseId:0,


score:5,


content:"",


stars:[1,2,3,4,5]


},



onLoad(options){


console.log("课程ID:",options.id)


this.setData({

courseId:Number(options.id)

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

url:"/api/reviews/course",

method:"POST",

data:{


course_id:this.data.courseId,

score:this.data.score,

content:this.data.content


}


})



wx.showToast({

title:"评价成功"

})



setTimeout(()=>{

wx.navigateBack()

},1000)



}catch(e){


console.log(e)


}


}


})
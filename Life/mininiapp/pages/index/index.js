Page({
  data: {
    status: '前后端连接测试已完成'
  },

  goLogin() {
    wx.navigateTo({
      url: '/pages/login/login'
    })
  }
})
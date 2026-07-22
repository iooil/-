const { request } = require('../../utils/request')

Page({
  data: {
    user: {
      name: '',
      student_no: '',
      major: ''
    },

    statistics: {
      selected_course_count: 0,
      course_review_count: 0,
      teacher_review_count: 0
    }
  },

  onLoad() {
    const token = wx.getStorageSync('token')

    if (!token) {
      wx.reLaunch({
        url: '/pages/login/login'
      })
      return
    }

    const user = wx.getStorageSync('user')

    if (user) {
      this.setData({
        user
      })
    }

    this.loadProfile()
    this.loadDashboard()
  },

  onShow() {
    if (wx.getStorageSync('token')) {
      this.loadDashboard()
    }
  },

  async loadProfile() {
    try {
      const result = await request({
        url: '/api/user/profile'
      })

      this.setData({
        user: result.data
      })

      wx.setStorageSync('user', result.data)
    } catch (error) {
      console.error('获取用户信息失败：', error)
    }
  },

  async loadDashboard() {
    try {
      const result = await request({
        url: '/api/user/dashboard'
      })

      this.setData({
        statistics: result.data
      })
    } catch (error) {
      console.error('获取首页统计失败：', error)
    }
  },

  goCourses() {
    wx.navigateTo({
      url: '/pages/courses/courses'
    })
  },

  goTeachers() {
    wx.navigateTo({
      url: '/pages/teachers/teachers'
    })
  },

  goMyCourses(){

    wx.navigateTo({
    
    url:
    "/pages/my-courses/my-courses"
    
    })
    
    },
    
    
    goMyReviews(){

      wx.navigateTo({
      
      url:
      "/pages/my-reviews/my-reviews"
      
      })
      
      },

  handleLogout() {
    wx.showModal({
      title: '退出登录',
      content: '确定退出当前账号吗？',

      success(res) {
        if (!res.confirm) {
          return
        }

        wx.removeStorageSync('token')
        wx.removeStorageSync('user')

        wx.reLaunch({
          url: '/pages/login/login'
        })
      }
    })
  }
})
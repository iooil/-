const { request } = require('../../utils/request')

Page({
  data: {
    studentNo: '',
    password: '',
    loading: false
  },

  onStudentNoInput(e) {
    this.setData({
      studentNo: e.detail.value
    })
  },

  onPasswordInput(e) {
    this.setData({
      password: e.detail.value
    })
  },

  async handleLogin() {
    const studentNo = this.data.studentNo.trim()
    const password = this.data.password

    if (!studentNo) {
      wx.showToast({
        title: '请输入学号',
        icon: 'none'
      })
      return
    }

    if (!password) {
      wx.showToast({
        title: '请输入密码',
        icon: 'none'
      })
      return
    }

    if (this.data.loading) {
      return
    }

    this.setData({
      loading: true
    })

    wx.showLoading({
      title: '登录中'
    })

    try {
      const result = await request({
        url: '/api/auth/login',
        method: 'POST',
        data: {
          student_no: studentNo,
          password: password
        }
      })

      wx.setStorageSync('token', result.data.token)
      wx.setStorageSync('user', result.data.user)

      wx.showToast({
        title: '登录成功',
        icon: 'success'
      })

      setTimeout(() => {
        wx.reLaunch({
          url: '/pages/home/home'
        })
      }, 800)
    } catch (error) {
      console.error('登录失败：', error)
    } finally {
      wx.hideLoading()

      this.setData({
        loading: false
      })
    }
  },

  goRegister() {
    wx.navigateTo({
      url: '/pages/register/register'
    })
  }
})
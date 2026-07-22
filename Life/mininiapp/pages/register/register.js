Page({
  data: {
    studentNo: '',
    name: '',
    major: '',
    password: '',
    confirmPassword: ''
  },

  onStudentNoInput(e) {
    this.setData({
      studentNo: e.detail.value
    })
  },

  onNameInput(e) {
    this.setData({
      name: e.detail.value
    })
  },

  onMajorInput(e) {
    this.setData({
      major: e.detail.value
    })
  },

  onPasswordInput(e) {
    this.setData({
      password: e.detail.value
    })
  },

  onConfirmPasswordInput(e) {
    this.setData({
      confirmPassword: e.detail.value
    })
  },

  handleRegister() {
    const {
      studentNo,
      name,
      major,
      password,
      confirmPassword
    } = this.data
  
    if (!studentNo || !name || !major || !password || !confirmPassword) {
      wx.showToast({
        title: '请完整填写注册信息',
        icon: 'none'
      })
      return
    }
  
    if (password.length < 6) {
      wx.showToast({
        title: '密码至少6位',
        icon: 'none'
      })
      return
    }
  
    if (password !== confirmPassword) {
      wx.showToast({
        title: '两次密码不一致',
        icon: 'none'
      })
      return
    }
  
    wx.showLoading({
      title: '注册中'
    })
  
    wx.request({
      url: 'http://127.0.0.1:8080/api/auth/register',
      method: 'POST',
      header: {
        'Content-Type': 'application/json'
      },
      data: {
        student_no: studentNo,
        name: name,
        major: major,
        password: password
      },
  
      success: (res) => {
        if (res.statusCode === 200 && res.data.code === 200) {
          wx.showToast({
            title: '注册成功',
            icon: 'success'
          })
  
          setTimeout(() => {
            wx.navigateBack()
          }, 1200)
        } else {
          wx.showToast({
            title: res.data.message || '注册失败',
            icon: 'none'
          })
        }
      },
  
      fail: (error) => {
        console.error('注册请求失败：', error)
  
        wx.showToast({
          title: '无法连接后端',
          icon: 'none'
        })
      },
  
      complete: () => {
        wx.hideLoading()
      }
    })
  }
})
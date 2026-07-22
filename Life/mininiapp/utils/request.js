const BASE_URL = 'http://10.4.29.225:8080'

function request(options) {
  const token = wx.getStorageSync('token')

  return new Promise((resolve, reject) => {
    wx.request({
      url: BASE_URL + options.url,
      method: options.method || 'GET',
      data: options.data || {},
      header: {
        'Content-Type': 'application/json',
        'Authorization': token ? `Bearer ${token}` : ''
      },

      success(res) {
        if (res.statusCode === 200 && res.data.code === 200) {
          resolve(res.data)
          return
        }

        if (res.statusCode === 401) {
          wx.removeStorageSync('token')
          wx.removeStorageSync('user')

          wx.showToast({
            title: res.data.message || '请重新登录',
            icon: 'none'
          })

          reject(res.data)
          return
        }

        wx.showToast({
          title: res.data.message || '请求失败',
          icon: 'none'
        })

        reject(res.data)
      },

      fail(error) {
        wx.showToast({
          title: '无法连接服务器',
          icon: 'none'
        })

        reject(error)
      }
    })
  })
}

module.exports = {
  request,
  BASE_URL
}
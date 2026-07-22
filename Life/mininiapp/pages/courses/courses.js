const { request } = require('../../utils/request')

Page({
  data: {
    keyword: '',
    courses: [],
    loading: false,

    colleges: ['全部学院'],
    courseTypes: ['全部类型'],

    collegeIndex: 0,
    typeIndex: 0,

    selectedCollege: '',
    selectedCourseType: ''
  },

  onLoad() {
    this.loadFilters()
    this.loadCourses()
  },

  onShow() {
    if (wx.getStorageSync('token')) {
      this.loadCourses()
    }
  },

  onPullDownRefresh() {
    this.loadCourses().finally(() => {
      wx.stopPullDownRefresh()
    })
  },

  onKeywordInput(e) {
    this.setData({
      keyword: e.detail.value
    })
  },

  handleSearch() {
    this.loadCourses()
  },

  handleClear() {
    this.setData({
      keyword: ''
    })

    this.loadCourses()
  },

  onCollegeChange(e) {
    const collegeIndex = Number(e.detail.value)
    const college = this.data.colleges[collegeIndex]

    this.setData({
      collegeIndex,
      selectedCollege: collegeIndex === 0 ? '' : college
    })

    this.loadCourses()
  },

  onTypeChange(e) {
    const typeIndex = Number(e.detail.value)
    const courseType = this.data.courseTypes[typeIndex]

    this.setData({
      typeIndex,
      selectedCourseType: typeIndex === 0 ? '' : courseType
    })

    this.loadCourses()
  },

  async loadFilters() {
    try {
      const result = await request({
        url: '/api/courses/filters'
      })

      this.setData({
        colleges: ['全部学院', ...result.data.colleges],
        courseTypes: ['全部类型', ...result.data.course_types]
      })
    } catch (error) {
      console.error('获取筛选条件失败：', error)
    }
  },

  async loadCourses() {
    if (this.data.loading) {
      return
    }

    this.setData({
      loading: true
    })

    try {
      const keyword = encodeURIComponent(this.data.keyword.trim())
      const college = encodeURIComponent(this.data.selectedCollege)
      const courseType = encodeURIComponent(
        this.data.selectedCourseType
      )

      const result = await request({
        url:
          `/api/courses?keyword=${keyword}` +
          `&college=${college}` +
          `&course_type=${courseType}`
      })

      const courses = result.data.map((item) => {
        const progressPercent =
          item.capacity > 0
            ? Math.min(
                (item.selected_count / item.capacity) * 100,
                100
              )
            : 0

        let scoreText = '暂无'

        if (item.review_count > 0) {
          scoreText = Number(item.average_score).toFixed(1)
        }

        return {
          ...item,
          progress_percent: progressPercent,
          score_text: scoreText
        }
      })

      this.setData({
        courses
      })
    } catch (error) {
      console.error('获取课程列表失败：', error)
    } finally {
      this.setData({
        loading: false
      })
    }
  },

  handleSelect(e) {
    const course = e.currentTarget.dataset.course

    if (course.is_selected) {
      wx.showToast({
        title: '你已经选择该课程',
        icon: 'none'
      })
      return
    }

    if (course.remaining <= 0) {
      wx.showToast({
        title: '该课程人数已满',
        icon: 'none'
      })
      return
    }

    wx.showModal({
      title: '确认选课',
      content: `确定选择《${course.course_name}》吗？`,

      success: async (modalResult) => {
        if (!modalResult.confirm) {
          return
        }

        wx.showLoading({
          title: '选课中'
        })

        try {
          await request({
            url: `/api/courses/${course.id}/select`,
            method: 'POST'
          })

          wx.showToast({
            title: '选课成功',
            icon: 'success'
          })

          this.loadCourses()
        } catch (error) {
          console.error('选课失败：', error)
        } finally {
          wx.hideLoading()
        }
      }
    })
  }
})
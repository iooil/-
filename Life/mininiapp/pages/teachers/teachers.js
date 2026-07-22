const { request } = require('../../utils/request')

Page({
  data: {
    keyword: '',
    teachers: [],
    loading: false,

    colleges: ['全部学院'],
    collegeIndex: 0,
    selectedCollege: ''
  },

  onLoad() {
    this.loadFilters()
    this.loadTeachers()
  },

  onShow() {
    if (wx.getStorageSync('token')) {
      this.loadTeachers()
    }
  },

  onPullDownRefresh() {
    this.loadTeachers().finally(() => {
      wx.stopPullDownRefresh()
    })
  },

  onKeywordInput(e) {
    this.setData({
      keyword: e.detail.value
    })
  },

  handleSearch() {
    this.loadTeachers()
  },

  handleClear() {
    this.setData({
      keyword: ''
    })

    this.loadTeachers()
  },

  onCollegeChange(e) {
    const collegeIndex = Number(e.detail.value)
    const college = this.data.colleges[collegeIndex]

    this.setData({
      collegeIndex,
      selectedCollege: collegeIndex === 0 ? '' : college
    })

    this.loadTeachers()
  },

  async loadFilters() {
    try {
      const result = await request({
        url: '/api/teachers/filters'
      })

      this.setData({
        colleges: ['全部学院', ...result.data.colleges]
      })
    } catch (error) {
      console.error('获取教师筛选条件失败：', error)
    }
  },

  async loadTeachers() {
    if (this.data.loading) {
      return
    }

    this.setData({
      loading: true
    })

    try {
      const keyword = encodeURIComponent(
        this.data.keyword.trim()
      )

      const college = encodeURIComponent(
        this.data.selectedCollege
      )

      const result = await request({
        url:
          `/api/teachers?keyword=${keyword}` +
          `&college=${college}`
      })

      const teachers = result.data.map((item) => {
        const nameFirst = item.name
          ? item.name.substring(0, 1)
          : '师'

        const scoreText =
          item.review_count > 0
            ? Number(item.average_score).toFixed(1)
            : '暂无'

        const reviewText =
          item.review_count > 0
            ? `${item.review_count}条评价`
            : '暂无评价'

        const courseText =
          item.course_count > 0
            ? `${item.course_count}门课程`
            : '暂无课程'

        return {
          ...item,
          name_first: nameFirst,
          score_text: scoreText,
          review_text: reviewText,
          course_text: courseText
        }
      })

      this.setData({
        teachers
      })
    } catch (error) {
      console.error('获取教师列表失败：', error)
    } finally {
      this.setData({
        loading: false
      })
    }
  },

  goTeacherDetail(e){

    let id=e.currentTarget.dataset.id
    
    
    wx.navigateTo({
    
    url:
    `/pages/teacher-detail/teacher-detail?id=${id}`
    
    })
    
    
    }
})
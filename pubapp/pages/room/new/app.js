// pages/room/new/app.js
const app = getApp()
Page({

    /**
     * 页面的初始数据
     */
    data: {
        pickerHidden: true,
        chosen: ''
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad(options) {

    },

    /**
     * 生命周期函数--监听页面初次渲染完成
     */
    onReady() {

    },

    /**
     * 生命周期函数--监听页面显示
     */
    onShow() {

    },

    /**
     * 生命周期函数--监听页面隐藏
     */
    onHide() {

    },

    /**
     * 生命周期函数--监听页面卸载
     */
    onUnload() {

    },

    /**
     * 页面相关事件处理函数--监听用户下拉动作
     */
    onPullDownRefresh() {

    },

    /**
     * 页面上拉触底事件的处理函数
     */
    onReachBottom() {

    },

    /**
     * 用户点击右上角分享
     */
    onShareAppMessage() {
    },

    pickerConfirm(e) {
        this.setData({
          pickerHidden: true
        })
        this.setData({
          chosen: e.detail.value
        })
    },

    pickerCancel() {
        this.setData({
            pickerHidden: true
        })
  },

    pickerShow() {
        this.setData({
            pickerHidden: false
        })
    },

    submit(e) {
        e.detail.value.user_info = app.globalData.userInfo
        console.log('form发生了submit事件，携带数据为：', e.detail.value)
        // TODO 发送请求创建房间

    },

    router(e) {
        wx.navigateTo({
            url: '/pages/room/router/app'
        })
    }

})
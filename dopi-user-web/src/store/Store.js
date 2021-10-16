class Store {

    constructor() {
        this.userInfo = null;
        this.loggedIn = this.userInfo != null
        this.message = null;
    }

    setUserInfo(userInfo) {
        console.log("userinfo ", userInfo)
        if (userInfo && userInfo.username) {
            this.userInfo = userInfo
            this.loggedIn = true
        } else {
            this.userInfo = null
            this.loggedIn = false
        }
    }

    setMessage(message) {
        this.message = message;
    }
}

export default new Store()
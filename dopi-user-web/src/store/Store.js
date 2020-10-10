class Store {

    constructor() {
        this.userInfo = null;
        this.loggedIn = this.userInfo != null
        this.users = null;
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

    setUsers(users) {
        this.users = users;
    }

}

export default new Store()
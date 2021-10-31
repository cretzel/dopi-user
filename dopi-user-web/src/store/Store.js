class Store {

    constructor() {
        this.message = null;
    }

    setUserInfo(userInfo) {
        if (userInfo && userInfo.username) {
            sessionStorage.setItem("userinfo", JSON.stringify(userInfo))
        } else {
            sessionStorage.removeItem("userinfo")
        }
    }

    getUserInfo() {
        let userInfo = sessionStorage.getItem("userinfo")
        if (userInfo != null) {
            return JSON.parse(userInfo);
        }
        return null;
    }

    isLoggedIn() {
        return this.getUserInfo() != null;
    }

    setMessage(message) {
        this.message = message;
    }
}

export default new Store()
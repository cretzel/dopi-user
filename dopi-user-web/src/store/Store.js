class Store {

    constructor() {
        this.message = null;
        this.userInfo = this.loadUserInfo();
        this.loggedIn = this.userInfo != null;
    }

    setUserInfo(userInfo) {
        if (userInfo && userInfo.username) {
            sessionStorage.setItem("userinfo", JSON.stringify(userInfo))
            this.userInfo = userInfo;
            this.loggedIn = true;
        } else {
            this.userInfo = null;
            this.loggedIn = false;
            sessionStorage.removeItem("userinfo")
        }
    
    }

    loadUserInfo() {
        let userInfo = sessionStorage.getItem("userinfo")
        if (userInfo != null) {
            return JSON.parse(userInfo);
        }
        return null;
    }

    setMessage(message) {
        this.message = message;
    }

    resetMessages() {
        this.message = null;
    }
}

export default new Store()
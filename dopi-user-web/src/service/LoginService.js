import store from "../store/Store";

class LoginService {

    constructor (store) {
        this.store = store
    }

    async login(login) {
        const response = await fetch("/api/user/login" , {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(login)
        });
        if (response.ok) {
            let userInfo = await response.json();
            store.setUserInfo(userInfo);
            return Promise.resolve();
        }
        store.setUserInfo(null);
        return Promise.reject("Error logging in");
    }

    async refreshToken() {
        const response = await fetch("/api/user/refresh", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            }
        });

        if (response.ok) {
            let userInfo = await response.json();
            store.setUserInfo(userInfo);
        } else {
            store.setUserInfo(null);
        }
    }

    async startTokenRefresher() {
        this.refreshToken();
        this.tokenRefresher();
    }

    async tokenRefresher() {
        const _this = this;
        window.setTimeout(function() {
            if(store.loggedIn) {
                _this.refreshToken();
            }
            _this.tokenRefresher();
        },  30 * 1000)

    }
}

export default new LoginService()


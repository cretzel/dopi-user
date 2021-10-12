class Store {

    constructor() {
        this.userInfo = null;
        this.loggedIn = this.userInfo != null
        this.users = null;
        this.user = null;
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

    setUser(userDto) {
        if (userDto == null) {
            this.user = {
                username: '',
                roles: ''
            }
            return
        }
        this.user = {
            username: userDto.username,
            roles: userDto.roles.join(', ')
        }
    }

    getUserDto() {
        if (this.user == null) {
            return null;
        }
        let roles = this.user.roles.split(/,\s/);
        console.log("user.roles", this.user.roles)
        console.log("roles", roles)
        return {
            username: this.user.username,
            roles: roles
        }
    }

}

export default new Store()
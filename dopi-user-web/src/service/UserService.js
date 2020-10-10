import store from "../store/Store";

class UserService {

    constructor (store) {
        this.store = store
    }

    async getMe() {
        const response = await fetch("/api/user/users/me" , {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        })

        if (response.ok) {
            let user = await response.json();
            console.log("Fetched user:", user);
            return user;
        } else {
            console.log("Error fetching user");
        }
    }

    async getUsers() {
        const response = await fetch("/api/user/users" , {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        })

        if (response.ok) {
            let users = await response.json();
            console.log("Fetched users:", users);
            store.setUsers(users);
        } else {
            console.log("Error fetching user");
        }
    }

}

export default new UserService()


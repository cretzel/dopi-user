class UserService {

    async getMe() {
        const response = await fetch("/api/user/users/me", {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        })

        if (response.ok) {
            let user = await response.json();
            return user;
        } else {
            console.log("Error fetching user");
        }
    }

    async getUsers() {
        const response = await fetch("/api/user/users", {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        })

        if (response.ok) {
            return response.json();
        } else {
            console.log("Error fetching user");
        }
    }

    async getUser(username) {
        const response = await fetch("/api/user/users/" + username, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        })

        if (response.ok) {
            return response.json();
        }
        return Promise.reject("Error loading user")
    }

    async putUser(user) {
        const response = await fetch("/api/user/users/" + user.username, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(user)
        });
        if (response.ok) {
            return await response.json();
        }
        console.log("Error updating user");
        return Promise.reject("Error updating user")
    }

    async postUser(user) {
        const response = await fetch("/api/user/users/" + user.username, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(user)
        });
        if (response.ok) {
            return await response.json();
        }
        return Promise.reject("Error creating user")
    }

    async deleteUser(username) {
        const response = await fetch("/api/user/users/" + username, {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json'
            }
        });
        if (response.ok) {
            return Promise.resolve();
        }
        return Promise.reject("Error deleting user");
    }

}

export default new UserService()


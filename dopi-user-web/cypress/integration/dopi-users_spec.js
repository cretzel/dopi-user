describe('Dopi Users', () => {

    let username;

    beforeEach(() => {
        cy.visit('http://localhost:8080/users/login')
        cy.get('#username').clear().type("admin")
        cy.get('#password').clear().type("admin")
        cy.get('#login-button').click()
        cy.get('.navbar-item.logged-in-user').should('have.text', "admin")

        const rd = Math.floor(Math.random() * 1000);
        username = `user-${rd}`
    })

    it('Should open user list', () => {
        cy.visit('http://localhost:8080/users')
        cy.get('.user-list')
        cy.get('.user-list table tbody tr:first td:first').should('have.text', 'admin')
        cy.get('.user-list table tbody tr:first td:first a').click()
    })

    it('Should open user details', () => {
        cy.visit('http://localhost:8080/users')
        cy.get('.user-list table tbody tr:first td:first a').click()
        cy.url().should('include', '/users/admin')
    })

    it('Should save user details', () => {
        cy.visit('http://localhost:8080/users/admin')
        cy.get('.user-details')
        cy.get('#roles').clear().type("admin, user")
        cy.get('#save').click()
        cy.visit('http://localhost:8080/users/admin')
        cy.get('#roles').should('have.value', 'admin, user')
    })

    it('Should create & delete user', () => {

        cy.visit('http://localhost:8080/users')
        cy.get('#create-user-button').click()
        cy.get('.new-user')
        cy.get('#username').clear().type(username)
        cy.get('#roles').clear().type("user, bar")
        cy.get('#password').clear().type("Secret1/")
        cy.get('#save').click()

        cy.get('.user-list')
        cy.get('.user-list tr[data-user = "' + username + '"] td:first').should('have.text', username)

        cy.visit('http://localhost:8080/users/' + username)
        cy.get('#delete').click();
        cy.get('#doDelete').click();

        cy.get('.user-list')
        cy.get('.user-list tr[data-user = "' + username + '"] td:first').should('not.exist')

    })

    it('Should validate new user details', () => {
        cy.visit('http://localhost:8080/users/new')
        cy.get('.new-user')
        cy.get('#username').clear().type('xxx')
        cy.get('#roles').clear()
        cy.get('#password').clear()
        cy.get('#save').click()

        cy.get('.user-list').should("not.exist")
        cy.get('.message.is-danger').should("exist")
    })


    it('New user should login', () => {
        newUser(username, "user");
        newUserLogin(username);
    })

    it('New user should not access user list', () => {
        newUser(username);
        newUserLogin(username);
        cy.visit('http://localhost:8080/users');
        cy.get('.user-list').should("not.exist");
    })

})

function newUser(username, roles) {
    cy.visit('http://localhost:8080/users/new');
    cy.get('.new-user');
    cy.get('#username').clear().type(username);
    cy.get('#roles').clear().type(roles ? roles : "user");
    cy.get('#password').clear().type("Secret1/");
    cy.get('#save').click();
}

function newUserLogin(username) {
    cy.visit('http://localhost:8080/users/login');
    cy.get('#username').clear().type(username);
    cy.get('#password').clear().type("Secret1/");
    cy.get('#login-button').click();
    cy.get('.navbar-item.logged-in-user').should('have.text', username)
}



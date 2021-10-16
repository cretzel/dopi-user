describe('Dopi Users', () => {

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
        const rd = Math.floor(Math.random() * 1000);
        const username = `user-${rd}`

        cy.visit('http://localhost:8080/users')
        cy.get('#create-user-button').click()
        cy.get('.new-user')
        cy.get('#username').clear().type(username)
        cy.get('#roles').clear().type("user, bar")
        cy.get('#password').clear().type("secret")
        cy.get('#save').click()

        cy.get('.user-list')
        cy.get('.user-list tr[data-user = "' + username + '"] td:first').should('have.text', username)

        cy.visit('http://localhost:8080/users/' + username)
        cy.get('#delete').click();
        cy.get('#doDelete').click();

        cy.get('.user-list')
        cy.get('.user-list tr[data-user = "' + username + '"] td:first').should('not.exist')

    })

})
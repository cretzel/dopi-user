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
        cy.get('#roles').should('have.value', 'admin')
        cy.get('#roles').type(" user")
        cy.get('#save').click()
        cy.visit('http://localhost:8080/users/admin')
        cy.get('#roles').should('have.value', 'admin, user')

    })
})
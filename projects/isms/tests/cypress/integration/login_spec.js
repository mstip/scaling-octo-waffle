describe('Login', () => {
    it('successfully loads', () => {
        cy.visit('');
    })

    it('login successfully', () => {
        const username = 'admin';
        const password = 'admin'
        cy.visit('');
        cy.get('input[name=userName]').type(username)
        cy.get('input[name=password]').type(password)
        cy.get('button[type=submit]').click()
        cy.url().should('include', '/hello')
        cy.get('h1').should('contain', 'Platzhalter')
    })

    it('should login failed', () => {
        const username = 'wrong';
        const password = 'wrong';
        cy.visit('');
        cy.get('input[name=userName]').type(username)
        cy.get('input[name=password]').type(password)
        cy.get('button[type=submit]').click()
        cy.url().should('not.include', '/hello')
        cy.get('.alert.alert-danger').should('exist');
        // cy.get('h1').should('contain', 'Platzhalter')
    })
})
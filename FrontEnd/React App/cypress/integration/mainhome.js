describe('My first test', () => {
    beforeEach(() => {
        cy.visit('http://localhost:3000')
        cy.get('.email').type('lucas')
        cy.get('.password').type('123')
        cy.contains('LoginConfirm').click()
    })
    it('Visit mainhome', () => {
        cy.contains('Home').click()
    })
    it('Chat', () => {
        cy.contains('Chat').click()
    })
    it('Hot Topic', () => {
        cy.contains('Hot Topic').click()
    })
    it('Username', () => {
        cy.contains('lucas').click()
    })
    it('New meeting', () => {
        cy.contains('New Meeting').click()
    })
    it('Join', () => {
        cy.contains('Join Meeting').click()
    })
    it('Friend', () => {
        cy.contains('Friend').click()
    })
    it('Contain-username', () => {
        cy.get('#icon').click().get('#username').contains('lucas')
    })
    it('Logout', () => {
        cy.get('#icon').click().get('#logout').contains('Sign out').click()
    })
})
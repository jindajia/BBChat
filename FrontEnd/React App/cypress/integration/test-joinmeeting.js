describe('joinmeeting page test', () => {
    beforeEach(() => {
        cy.visit('http://localhost:3000')
        cy.get('.email').type('lucas')
        cy.get('.password').type('123')
        cy.contains('LoginConfirm').click()
        cy.contains('Join Meeting').click()
    })
    it('Visit mainhome', () => {
        cy.contains('Home').click()
    })
    it('Chat', () => {
        cy.contains('Chat').click()
    })
    it('Username', () => {
        cy.contains('lucas').click()
    })
    it('Room number input', () => {
        cy.get('#roomnumber').type('happy').should('have.value','happy')
    })
    it('Room password input', () => {
        cy.get('#password').type('hello').should('have.value','hello')
    })
    it('Create button', () => {
        cy.contains('JOIN').click()
    })
})
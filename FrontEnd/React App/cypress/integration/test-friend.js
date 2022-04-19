describe('friend page test', () => {
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
    it('Username', () => {
        cy.contains('lucas').click()
    })
    it('Add friend', () => {
        cy.contains('Add Friend').click()
    })
    it('Random', () => {
        cy.contains('Random').click()
    })
})
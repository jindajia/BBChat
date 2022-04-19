describe('createmeeting page test', () => {
    beforeEach(() => {
        cy.visit('http://localhost:3000')
        cy.get('.email').type('lucas')
        cy.get('.password').type('123')
        cy.contains('LoginConfirm').click()
        cy.contains('New Meeting').click()
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
    it('Room name input', () => {
        cy.get('#roomname').type('happy').should('have.value','happy')
    })
    it('Room password input', () => {
        cy.get('#roompassword').type('hello').should('have.value','hello')
    })
    it('Create button', () => {
        cy.contains('CreateRoom').click()
    })
    it('Random password input', () => {
        cy.get('#randompassword').click()
    })
    it('Try create with password', () =>{
        cy.get('#roomname').type('happy').should('have.value','happy')
        cy.get('#roompassword').type('hello').should('have.value','hello')
        cy.contains('CreateRoom').click()
	    cy.on('window:confirm', () => true);
	})
    it('Try create with random password', () =>{
        cy.get('#roomname').type('happy').should('have.value','happy')
        cy.get('#randompassword').click()
        cy.contains('CreateRoom').click()
	    cy.on('window:confirm', () => true);
	})
})
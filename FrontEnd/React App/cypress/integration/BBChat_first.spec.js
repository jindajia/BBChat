// BBChat_first.spec.js created with Cypress
//
// Start writing your Cypress tests below!
// If you're unfamiliar with how Cypress works,
// check out the link below and learn how to write your first test:
// https://on.cypress.io/writing-first-test
describe('My first test',() =>{
	it('Visit BBChat website',()=>{
		cy.visit('http://localhost:3000')
		cy.contains('Login').click()
	})
	it('Try Registration',() =>{
		cy.visit('http://localhost:3000')
		cy.contains('Registration').click()
		cy.get('.email').type('jinda').should('have.value','jinda')
		cy.get('.password').type('jiajinda').should('have.value','jiajinda')
		cy.contains('RegistrationConfirm').click()
	})
})
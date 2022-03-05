// BBChat_first.spec.js created with Cypress
//
// Start writing your Cypress tests below!
// If you're unfamiliar with how Cypress works,
// check out the link below and learn how to write your first test:
// https://on.cypress.io/writing-first-test
describe('My first test',() =>{
  	beforeEach(() => {
		cy.visit('http://localhost:3000')
	})
	it('Visit BBChat website',() =>{
		cy.contains('Login').click()
	})
	it('Try Registration',() =>{
		cy.contains('Registration').click()
		cy.get('.email').type('jinda').should('have.value','jinda')
		cy.get('.password').type('jiajinda').should('have.value','jiajinda')
		cy.contains('RegistrationConfirm').click()
	})
    it('check the Sign up button', function(){
		cy.contains('Registration').click()
		cy.get('.email').type('jinda').should('have.value','jinda')
		cy.get('.password').type('jiajinda').should('have.value','jiajinda')
        cy.contains('Username is not available.')
        cy.on('window:alert',(txt)=> {
            //Mocha assertions
            expect(txt).to.contains('The user already exists');
        })
    })
	it('Try Login', () =>{
		cy.contains('Login').click()
		cy.get('.email').type('jinda').should('have.value','jinda')
		cy.get('.password').type('jiajinda').should('have.value','jiajinda')
		cy.contains('LoginConfirm').click()
	    cy.on('window:confirm', () => true);
	})
	it('Test Profile', ()=>{
		cy.contains('Login').click()
		cy.get('.email').type('jinda').should('have.value','jinda')
		cy.get('.password').type('jiajinda').should('have.value','jiajinda')
		cy.contains('LoginConfirm').click()
		cy.contains('jinda')
	})
	it('Test Logout', ()=>{
		cy.contains('Login').click()
		cy.get('.email').type('jinda').should('have.value','jinda')
		cy.get('.password').type('jiajinda').should('have.value','jiajinda')
		cy.contains('LoginConfirm').click()
		cy.contains('Logout').click()
	})
})
describe('My first test',() =>{
    beforeEach(() => {
      cy.visit('http://localhost:3000')
  })
  it('Visit BBChat website',() =>{
      cy.contains('Login').click()
  })
  it('Try Registration',() =>{
      cy.contains('Registration').click()
      cy.get('.email').type('lucas').should('have.value','lucas')
      cy.get('.password').type('123').should('have.value','123')
      cy.contains('RegistrationConfirm').click()
  })
  it('Try Login', () =>{
      cy.contains('Login').click()
      cy.get('.email').type('lucas').should('have.value','lucas')
      cy.get('.password').type('123').should('have.value','123')
      cy.contains('LoginConfirm').click()
      cy.on('window:confirm', () => true);
  })
})
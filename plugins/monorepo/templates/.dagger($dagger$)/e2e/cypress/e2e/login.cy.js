describe('Login', () => {
    it('Visits the home page and tries to log in', () => {
      // Open the website
      cy.visit('/');
  
      // Fill out the form
      cy.get('input[name=email]').should('be.visible').type('admin@admin.com');
      cy.get('input[name=password]').should('be.visible').type('Admin@123!');

      cy.get('button[type=submit]').should('be.visible').click({timeout: 10000});

      cy.url().should('include', '/user');
    });
  });
  
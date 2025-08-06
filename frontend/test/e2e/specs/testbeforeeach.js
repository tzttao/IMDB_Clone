describe('My First Test', () => {
  before(() => {
    cy.login()
    cy.log('整个describe运行前运行一次，做一些准备工作')
  })
  beforeEach(() => {
    cy.log('每个it之前都会执行，做一些准备工作')
  })
  it('Visits the app root url', () => {
    cy.login('admin', 'admin')
  })

  afterEach(() => {
    cy.log('每个it之后都会执行，做一些清理工作')
  })
  after(() => {
    cy.login()
    cy.log('整个describe运行完成后运行一次，做一些清理工作')
  })
})

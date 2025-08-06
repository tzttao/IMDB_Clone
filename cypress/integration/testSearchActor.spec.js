// testSearchActor.spec.js created with Cypress
//
// Start writing your Cypress tests below!
// If you're unfamiliar with how Cypress works,
// check out the link below and learn how to write your first test:
// https://on.cypress.io/writing-first-test
describe('actor', () => {
    it('actor', () => {
        const actorName = 'tom'
        cy.visit('http://localhost:8080/#/UserContainer/Homepage')
        // 查找search里包含登录字符串的元素，如果查找不到则失败

        cy.get('input[placeholder="Search"]') //获取input框，写法与jQuery的selector一致
            .clear()
            .type(actorName)  // input框里面输入用户名
            .should('have.value', actorName) // 断言 input的value=username
        cy.get('#search').type('{enter}')// 提交表单
        // 查找按钮，然后点击
        cy.get('table')
            .contains('Hanks').click()
    })
})

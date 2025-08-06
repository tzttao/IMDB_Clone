// testHomepage.spec.js created with Cypress
//
// Start writing your Cypress tests below!
// If you're unfamiliar with how Cypress works,
// check out the link below and learn how to write your first test:
// https://on.cypress.io/writing-first-test
describe('homepage', () => {
    it('test the homepage', () => {
        const searchactor = 'tom'
        cy.visit('http://localhost:8080/#/UserContainer/Homepage')

        cy.get('header')
            .get('i[class="el-icon-s-home"]')
            .get('#menu').click()
            .get('.el-drawer')
                .contains('IMDb Clone')
            .get('.el-drawer')
                .contains('span','Welcome to IMDb!')
            .get('span')

        cy.get('main')
        cy.get('footer')


        cy.get('div[class="block"]')
            .get('span[class="demonstration"]').contains('Feature Today')
            .get('.el-carousel').get('.el-image')
            .get('span[class="demonstration"]').contains('Top Picks')
            .get('.el-carousel').get('.el-image')



        // 查找button.button-login里包含登录字符串的元素，如果查找不到则失败

        // cy.get('input[placeholder="Please enter your username"]') //获取input框，写法与jQuery的selector一致
        //     .clear()
        //     .type(username)  // input框里面输入用户名
        //     .should('have.value', username) // 断言 input的value=username
        // // 输入密码
        // cy.get('input[placeholder="Please enter your password"]')//获取input框，写法与jQuery的selector一致
        //     .clear()
        //     .type(password)
        //     .should('have.value', password)
        // // 提交表单
        // cy.get('button').click() //查找按钮，然后点击
        //
        // cy.visit('#/UserContainer/HomePage')  //校验是否登录成功
    })
})

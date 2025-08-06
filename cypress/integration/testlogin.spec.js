describe('login', () => {
    it('test login', () => {
        const username = 'blessedroll'
        const password = '246810'
        cy.visit('http://localhost:8080/#/UserLogIn')
        cy.contains('button', 'Sign In')  // 查找button.button-login里包含登录字符串的元素，如果查找不到则失败

        cy.get('input[placeholder="Please enter your username"]') //获取input框，写法与jQuery的selector一致
            .clear()
            .type(username)  // input框里面输入用户名
            .should('have.value', username) // 断言 input的value=username
        // 输入密码
        cy.get('input[placeholder="Please enter your password"]')//获取input框，写法与jQuery的selector一致
            .clear()
            .type(password)
            .should('have.value', password)
        // 提交表单
        cy.get('button').click() //查找按钮，然后点击


        cy.contains('blessedroll').click()
        cy.contains('My profile').click()
    })
})

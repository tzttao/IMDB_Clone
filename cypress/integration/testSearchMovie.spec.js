describe('actor', () => {
    it('actor', () => {
        const movieName = 'forrest'
        cy.visit('http://localhost:8080/#/UserContainer/Homepage')
        // 查找search里包含登录字符串的元素，如果查找不到则失败

        cy.get('input[placeholder="Search"]') //获取input框，写法与jQuery的selector一致
            .clear()
            .type(movieName)  // input框里面输入用户名
            .should('have.value', movieName) // 断言 input的value=username
        cy.get('#search').type('{enter}')// 提交表单
        // 查找按钮，然后点击
        // cy.visit('#/UserContainer/HomePage')  //校验是否登录成功

        cy.get('table')
            .contains('Forrest').click()

    })


})

// https://docs.cypress.io/api/table-of-contents

describe('登录', () => {
  it('测试登录', () => {
    const username = 'admin'
    const password = 'admin'
    cy.visit('/#/login')
    cy.contains('button.button-login', '登录')  // 查找button.button-login里包含登录字符串的元素，如果查找不到则失败

    cy.get('input[placeholder="用户名"]') //获取input框，写法与jQuery的selector一致
      .clear()
      .type(username)  // input框里面输入用户名
      .should('have.value', username) // 断言 input的value=username
    // 输入密码
    cy.get('input[placeholder="密码"]')//获取input框，写法与jQuery的selector一致
      .clear()
      .type(password)
      .should('have.value', password)
    // 提交表单
    cy.get('button.button-login').click() //查找按钮，然后点击

    cy.contains('首页')  //校验是否登录成功
  })
})

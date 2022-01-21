<template>
    <div>
        <button @click="testAxios">TestAxios</button>
    </div>
    <!--App -->
</template>

<script>
    export default {
        methods: {
            testAxios() {
                // 由于 main.js 里全局定义的 axios,此处直接使用 $axios 即可。
                // 由于 main.js 里定义了每个请求前缀，此处的 / 即为 /api/， 
                // 经过 vue.config.js 配置文件的代理设置，会自动转为 https://www.baidu.com/，从而解决跨域问题
                this.$axios.get('/').then(response => {
                    if (response.data) {
                        console.log(response.data)
                    }
                }).catch(err => {
                    alert('请求失败')
                })
            }
        }
    }
</script>

<style>

</style>
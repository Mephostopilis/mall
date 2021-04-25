module.exports = {
    // baseUrl: './',
    // assetsDir: 'static',
    // productionSourceMap: false,
    // devServer: {
    //     proxy: {
    //         '/api':{
    //             target:'http://jsonplaceholder.typicode.com',
    //             changeOrigin:true,
    //             pathRewrite:{
    //                 '/api':''
    //             }
    //         }
    //     }
    // }

    transpileDependencies: ['luch-request'],
    css: {
        requireModuleExtension: false,
        loaderOptions: {
            less: {
                // javascriptEnabled: true,
                // modifyVars: {
                //     'hack': `true; @import "${require('path').resolve(process.cwd(), './src').replace(/\\/g, '\\\\')}/variables.less";`
                // }
            }
        }
    }
}

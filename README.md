# baidu-translation-workflow

1. 下载 `baidu-translation.alfredworkflow` 并拖进 alfred 中完成安装。
2. 在百度的[翻译开放平台](https://api.fanyi.baidu.com/api/trans/product/desktop?req=developer)申请创建应用，申请后会得到一个 appId 和秘钥。
![百度开放平台页面.jpg](http://ww1.sinaimg.cn/large/8c9b2ef7gy1gk0ei1f1jfj21y0184jy3.jpg)
3. 在 alfred 中该 workflow 的页面里点击右上角的图标添加环境变量，如下图所示：
![跳转添加环境变量页面.jpg](http://ww1.sinaimg.cn/large/8c9b2ef7gy1gk0ejikvp5j21b80mqwm2.jpg)
4. 在添加环境变量的页面添加之前在百度开放平台页面申请的 appId 和秘钥。其中 name 那一列固定指定是 **`appId`** 和 **`securityKey`** 这两个值。
![填写环境变量的页面.jpg](http://ww1.sinaimg.cn/large/8c9b2ef7gy1gk0eio1py6j21nk0f6jyp.jpg)
5. 使用时通过 **`fy`** 后面跟要翻译的内容
![使用界面.jpg](http://ww1.sinaimg.cn/large/8c9b2ef7gy1gk0ej76k8vj20wi06un0w.jpg)

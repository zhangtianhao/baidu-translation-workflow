# baidu-translation-workflow

1. 下载 `baidu-translation.alfredworkflow` 的[最新版本](https://github.com/zhangtianhao/baidu-translation-workflow/releases)并拖进 alfred 中完成安装。
2. 在百度的[翻译开放平台](https://api.fanyi.baidu.com/api/trans/product/desktop?req=developer)申请创建应用，申请后会得到一个 appId 和秘钥。
![百度开放平台页面.jpg](https://github.com/zhangtianhao/baidu-translation-workflow/blob/main/image/%E7%99%BE%E5%BA%A6%E5%BC%80%E6%94%BE%E5%B9%B3%E5%8F%B0%E9%A1%B5%E9%9D%A2.jpg)
3. 在 alfred 中该 workflow 的页面里点击右上角的图标添加环境变量，如下图所示：
![跳转添加环境变量页面.jpg](https://github.com/zhangtianhao/baidu-translation-workflow/blob/main/image/%E8%B7%B3%E8%BD%AC%E6%B7%BB%E5%8A%A0%E7%8E%AF%E5%A2%83%E5%8F%98%E9%87%8F%E9%A1%B5%E9%9D%A2.jpg)
4. 在添加环境变量的页面添加之前在百度开放平台页面申请的 appId 和秘钥。其中 name 那一列固定指定是 **`appId`** 和 **`securityKey`** 这两个值。
![填写环境变量的页面.jpg](https://github.com/zhangtianhao/baidu-translation-workflow/blob/main/image/%E5%A1%AB%E5%86%99%E7%8E%AF%E5%A2%83%E5%8F%98%E9%87%8F%E7%9A%84%E9%A1%B5%E9%9D%A2.jpg)
5. 使用时通过 **`fy`** 后面跟要翻译的内容
![使用界面.jpg](https://github.com/zhangtianhao/baidu-translation-workflow/blob/main/image/%E4%BD%BF%E7%94%A8%E7%95%8C%E9%9D%A2.jpg)

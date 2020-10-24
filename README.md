# baidu-translation-workflow

1. 下载 `baidu-translation.alfredworkflow` 并拖进 alfred 中完成安装。
2. 在百度的[翻译开放平台](https://api.fanyi.baidu.com/api/trans/product/desktop?req=developer)申请创建应用，申请后会得到一个 appId 和秘钥。
![翻译开放平台页面](https://github.com/zhangtianhao/baidu-translation-workflow/blob/main/image/%E7%99%BE%E5%BA%A6%E5%BC%80%E6%94%BE%E5%B9%B3%E5%8F%B0%E9%A1%B5%E9%9D%A2.jpg)
3. 在 alfred 中该 workflow 的页面里点击右上角的图标添加环境变量，如下图所示：
4. 在添加环境变量的页面添加之前在百度开放平台页面申请的 appId 和秘钥。其中 name 那一列固定指定是 **`appId`** 和 **`securityKey`** 这两个值。
5. 使用时通过 **fy** 后面跟要翻译的内容

# yuque_webhook

yuque webhook

### 使用腾讯云 serverless 部署语雀飞书机器人

> 1. 语雀很好用，但是没有App，无法通知
> 2. 飞书群自定义机器人很方便，但是不知道用来干啥
> 3. 腾讯云函数计算有免费额度，是不是可以利用一下？
> 4. 那是不是可以把语雀的webhook 使用飞书通知，部署在腾讯云上呢？
> 5. 好主意


### 使用方式

1. 在飞书群中添加自定义机器人，复制webhook 地址，拿到其中的 hook_id
2. 执行代码build 中 scf_build.sh 脚本，得到 api.zip 包
3. 在腾讯云云函数中创建云函数，
    * 选择golang
    * 选择本地zip 包
    * 选择上一步生成的 api.zip
    * 超时时间为3S
    * 创建云函数
4. 新建API网关服务
    * 新建API
    * 路径为 /api/webhook
    * 请求方法为POST
    * 免鉴权
    * 后端配置，后端类型为云函数SCF
    * 选择上一步创建的云函数
    * 超时时间为3S
    * 启用响应集成
    * 响应结果返回类型为JSON
    * 完成选择发布
    * 从网关基础配置中拿到公网访问地址，host
5. 打开想要添加webhook的语雀知识库
    * 知识库设置，开发者，添加webhook，命名为飞书机器人
    * URL为 host/api/webhook（这里host 为上一步得到的host地址，包含80 或者 443)
    
DONE

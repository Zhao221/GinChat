# GinChat

# 项目名称
IM（Gin+WebSocket）

# 系统架构
四层：前端，接入层，逻辑层，持久层

# 技术栈
webSocket，channel/goroutine，gin，gorm，swagger，logrus auth，sql，nosql，mq...

# 核心功能
1.发送和接收消息，文字 表情 图片 音频
2.访客，点对点，群聊
3.广播，快捷回复，撤回，心跳检测...

# 消息发送流程
A>登录>鉴权>(游客)>消息类型>群/广播>B 
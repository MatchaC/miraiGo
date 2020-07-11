# miraiGo

**miraiGo**是基于[**mirai-http-api**](https\://github.com/project-mirai/mirai-api-http)的Go封装  
- **origin**下的API是基于原始API的一次封装  
- **origin_test**是针对**origin**中API的一系列测试  
- **根目录下的API**(尚未公开)是基于本人使用习惯对**origin**进行了二次封装  
- **root_test**(尚未公开)是针对根目录中API的一系列测试

# 开始使用
`go get -u github.com/MatchaC/miraiGo`
# API(origin)列表

    func About(remoteAddr string) (RespAbout, error)
    
    func Auth(remoteAddr string, authKey string) (RespAuth, error)
    
    func Verify(remoteAddr string, sessionKey string, qq int) (RespVerify, error)
    
    func Release(remoteAddr string, sessionKey string, qq int) (RespRelease, error)
    
    func SendFriendMessage(remoteAddr string, sessionKey string, targetId int, messageChain []Message) (RespSendMessage, error)
    
    func SendGroupMessage(remoteAddr string, sessionKey string, targetGroupId int, messageChain []Message) (RespSendMessage, error)
    
    func SendTempMessage(remoteAddr string, sessionKey string, targetGroupId int, targetId int, messageChain []Message) (RespSendMessage, error)

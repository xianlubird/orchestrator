# orchestrator
A simple orchestrator for docker just for personal learning.

#目标

##1.能够动态部署和增加减少tomcat实例

为啥那么热爱tomcat。因为我是搞java出身，和tomcat搞了好多年基，喜欢他没商量。这个项目，就先拿它练手。

目标就是能够在后台控制启动一个或者关闭一个tomcat。然后代码部署的是一套代码。


##2. 能够跨host部署

能在一台机器上运行deamon，然后部署另外一台机器的tomcat，让他重启运行，然后还共享一套代码。

##3.宏伟目标
能有服务发现机制，处理log，调度，健康监控等等。能做多少做多少吧。

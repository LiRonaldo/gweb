# gweb
go语言手撸web框架
<h4>一
<h5>base1 最基本的接受请求。响应。
<h5>base2 重写ServeHTTP,将所有的请求都落到ServeHTTP方法中，具体分发和处理在ServeHTTP方法中
<h5>base3 进一步优化，构造一个engine对象。并添加get post方法，run方法，重写ServeHTTP 更加复杂的逻辑处理。将用户设置的url存储起来，根据url去调用handler，设计比较像框架了

<h4>二
<h5>新增context包，包下新增context，router，engine,是并在优化，使其更像一个框架
<h4>三
<h5>新增路有树控制。更灵活。
<h4>四
<h5>新增日志中间件
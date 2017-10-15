# website-scanner
一个使用Golang开发的高性能WEB目录扫描器

# Usage
- url 输入URL，可以为任意URL，程序运行时将会自动进行URL修正，自动加上http前缀，自动去除path中的文件名以及querystring以及hash等等。 (example:http://127.0.0.1/test)
- method HTTP提交方法，推荐使用性能更高流量更小的HEAD，但是特殊情况下如果有防火墙禁止了HEAD方式提交，可以使用GET等方式 (example:HEAD)
- header 头部文件的文件名。如果某些特殊路径的访问需要带上特定的头部或者cookies，则可以在header中设置相应的header。 (example:header.txt)
- thread 线程数量。推荐与CPU的线程数量一致，如果服务器出现过高线程扫死的情况可以适当调低线程。 (example:10)
- path 字典路径。程序在初始化的时候将会自动将字典文件读入内存，并且进行去重操作，所以请务必在程序启动之前正确放置好字典文件，而且无须担心字典文件被重复放置的问题。 (example:字典)
- extname 字典文件扩展名。 (example:txt)

# TODO
- 使用代理进行扫描。
- 递归地将扫描结果再次进行扫描，实现类爬虫的多级扫描。

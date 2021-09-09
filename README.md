

- go-redis： https://github.com/go-redis/redis
- viper： https://github.com/spf13/viper
- zap： https://github.com/uber-go/zap
- rotatelogs： https://github.com/lestrrat-go/file-rotatelogs
- jwt  https://github.com/golang-jwt/jwt
- swaggo/gin-swagger https://github.com/swaggo/gin-swagger


## Redis Example

```
ctx := context.Background()
val, err := global.REDIS.Get(ctx,"name").Result()
if err != nil {
    panic(err)
}
fmt.Println("My Name is ", val)
```



##跨平台

```
注意点一、

go build 的时候会选择性地编译以系统名结尾的文件(linux、darwin、windows、freebsd)。例如Linux(Unix)系统下编译只会选择array_linux.go文件，其它系统命名后缀文件全部忽略。


注意点二、

在xxx.go文件的文件头上添加 // + build !windows (tags)，可以选择在windows系统下面不编译 

// +build !windows

package main
```

### [Httplib(beego)](https://beego.me/docs/module/httplib.md)

#### 支持的方法对象

- Get(url string)
- Post(url string)
- Put(url string)
- Delete(url string)
- Head(url string)
# SSL认证

### 1、生成私钥
openssl genrsa -out server.key 2048

### 2、生成证书 全部回车即可，可以不填
openssl req -new -x509 -key server.key -out server.crt -days 36500

### 国家名称
Country Name(2 letter code) [AU]:CN

### 省名称
State or Province Name (full name) [Some-State]:GuangDong

### 城市名称
Locality Name (eg, city) []:Meizhou

### 公司组织名称
Organization Name (eg, company) [Internet Widgits Pty Ltd]:Xuexiangban

### 部门名称
Organizational Unit Name (eg, section) []:go

### 服务器or网站名称
Common Name (e.g. server FQDN or YOUR name) []:kuangstudy

### 邮件
Email Address []:24736743@qq.com

### 3、生成 csr
openssl req -new -key server.key -out server.csr

### 更改openssl.cnf (Linux是openssl.cfg)
### 1.复制一份你安装的openssl的bin目录里面的openssl.cnf文件到你项目所在的目录
### 2.找到[ CA_default ],打开copy_extensions = copy (就是把前面的#去掉)
### 3.找到[ req ],打开req_extensions = v3_req # The extensions to add to a certificate request
### 4.找到[ v3_req ],添加 subjectAltName = @alt_names
### 5.添加新的标签 [ alt_names ] ,和标签字段
DNS.1 = *.kuangstudy.com

### 生成证书私钥test.key
openssl genpkey -algorithm RSA -out test.key

### 通过私钥test.key生成证书请求文件test.csr(注意cfg和cnf)
openssl req -new -nodes -key test.key -out test.csr -days 3650 -subj "/C=cn/OU=myorg/O=mycomp/CN=myname)" -config
./openssl.cnf -extensions v3_req
### test.csr是上面生成的证书请求文件。ca.crt/server.key是CA证书文件和key,用来对test.csr进行签名认证。这两个文件在第一部分生成。

### 生成SAN证书 pem
openssl x509 -req -days 365 -in test.csr -out test.pem -CA server.crt -CAkey server.key -CAcreateserial -extfile
./openssl.cnf -extensions v3_req

# Token认证
我们先看一个grpc提供我们的一个接口，这个接口中有两个方法，接口位于credentials包下，这个接口需要客户端实现
 type PerRPCCredentials interface {
    GetRequestMetadata(ctx context.Context , uri ...string)(map[string]string,error)
    RequireTransportSecurity() bool
}
第一个方法作用是获取元数据信息，也就是客户端提供的key，value对，context用于控制超时和取消，uri是请求入口处的uri
第二个方法的作用是否需要基于TLS认证进行安全传输，如果返回值是true，则必须加上TLS认证，返回值是false则不用

gRPC将各种认证方式浓缩统一到一个凭证（credentials）上，可以单独使用一种凭证，比如只使用TLS凭证或者只使用自定义凭证，也可以多种
凭证组合，gRPC提供统一的API验证机制，是研发人员使用方便，这也是gRPC设计的巧妙之处
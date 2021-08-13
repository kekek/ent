


### 引入ent opts.proto

```

wget -O ./third-party/ent/options/opts.proto https://raw.githubusercontent.com/ent/contrib/master/entproto/cmd/protoc-gen-ent/options/ent/opts.proto

```

### 

```
protoc -I=./third-party/ent/ --ent_out=. --ent_opt=schemadir=./ent/schema/ api/hello/hello.proto

```

### 参考文档

- https://golang.com.cn/entgo.io/contrib/entproto/cmd/protoc-gen-ent

- 
- 说明

user 表定义edge

card 定义关联键 user_id, edge 指向field user_id

```
syntax = "proto3";

import "third-party/ent/options/opts.proto";

option go_package = "ent-todo/api/user;user";

message User {
    option (ent.schema).gen = true;

    int64 id = 1;
    string account = 2;
    string password = 3;
    string name = 4;
    //    google.protobuf.Timestamp created_at = 5;
    //    google.protobuf.Timestamp updated_at = 6;
    repeated Card card = 7 [(ent.edge) = {}];
}

message Card {
    option (ent.schema).gen = true;

    int64 id = 1;
    string imei = 2;
    string os_version = 3;
    string os_type = 4;
    int64 user_id = 5;
    User user = 6 [(ent.edge) = {ref: "card", unique: true, field: "user_id",required: true}];
}

```

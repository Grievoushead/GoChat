// Copyright 2013 Beego Samples authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package controllers

import (
	"samples/WebIM/db"
	"samples/WebIM/redis"
	"time"
)

// WebSocketController handles WebSocket requests.
type TestController struct {
	baseController
}

func (this *TestController) Test() {
	dbRes := db.GetResult()
	this.Data["MongoResult"] = dbRes
	redisRes := redis.GetResult()
	this.Data["RedisResult"] = redisRes


	this.Data["Divide"] = redisRes / dbRes;
	this.TplNames = "test.html"
}

func (this *TestController) Test2() {
	dbRes := db.GetResult()
	this.Data["MongoResult"] = dbRes
	redisRes := redis.GetResult()
	this.Data["RedisResult"] = redisRes

	time.Sleep(10000 * time.Millisecond)


	this.Data["Divide"] = redisRes / dbRes;
	this.TplNames = "test.html"
}

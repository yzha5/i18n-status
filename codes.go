/**
 * Copyright (c) yangzhao 2022/3/25
 *
 * File name: codes.go
 * Created at: 2022/3/25 9:44 AM
 * Author: yangzhao yzha5@163.com
 *
 * Description: Define status codes
 */

package status

type Code uint32

const (
	// OK 一般是请求成功后返回的状态
	OK Code = iota

	// Unknown 无法判断错误的出处
	Unknown

	// Gateway 网关
	Gateway

	// Server 服务端通用状态码
	// 表现为无法监听端口、启动服务失败等
	Server

	// Plugins 插件表现的状态码
	Plugins

	// Package 第三方包
	// 例如加密字符串出现错误，但又无法准确判断错误类型
	Package

	// Database 数据库
	// 例如插入数据时出现的错误，例如数据库离线等。
	// 但有些错误可能是用户请求（Request）数据上的错误，例如重复的用户名
	Database

	// Cache 缓存
	// 例如Redis离线、键值对不存在等
	Cache

	// Permission 用户权限
	// 例如需要用户登录后操作、用户权限等级不足等
	Permission

	// Redirect 重定向
	Redirect //重定向

	// Develop 开发者代码
	// 例如缺少参数等
	Develop //代码方面的错误

	// Request 用户请求
	// 用户请求数据上出现错误，比如要求是数字类型的参数却传了字母
	Request //用户请求的错误
)

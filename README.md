# goyht

Golang SDK for [YunHeTong](http://sdk.yunhetong.com/)，详细API文档，请阅读[Documentation](https://github.com/leesper/goyht/blob/master/Documentation.md)

# 一. 用户相关接口

## 1. 导入用户信息

```go
func (*Client) AddUser
```

对接第一步导入用户信息。

## 2. 获取用户登录凭证token

```go
func (*Client) UserToken
```

获取用户登录凭证，除了导入用户和获取用户登录凭证这个两个接口外，其他所有的接口地址后都要跟token参数。

## 3. 修改用户手机号

```go
func (*Client) ModifyPhoneNumber
```

用户信息变更后，可以调用该接口修改导入到云合同SDK系统中的数据；地址中的token是获取token接口返回的内容。

## 4. 修改用户名

```go
func (*Client) ModifyUserName
```

用户信息变更后，可以调用该接口修改导入到云合同SDK系统中的数据；地址中的token是获取token接口返回的内容。

# 二. 合同接口

## 1. 根据模板生成合同

```go
func (*Client) CreateTemplateContract
```

根据模版生成合同，功能有：将固定格式字符串替换成指定内容、将固定字符串替换成图片、简单表格操作。

## 2. 上传文件生成合同

```go
func (*Client) CreateFileContract
```

上传文件生成合同。

## 3. 添加参与者

```go
func (*Client) AddPartner
```

合同创建人才能添加参与者且创建人必须是参与者。

## 4. 合同自动签署

```go
func (*Client) SignContract
```

对接平台方有时需要对某些合同实现自动签署，使用此接口完成自动签署功能，调用者token必须为平台用户。

## 5. 合同作废

```go
func (*Client) InvalidateContract
```

合同参与者才能作废。

## 6. 合同列表

```go
func (*Client) ListContracts
```

平台用户才能获取合同列表，返回该应用下已完成、已作废的合同。

## 7. 合同签署状态详情

```go
func (*Client) LookupContractDetail
```

平台及参与用户均可查看。

## 8. 合同下载

```go
func (*Client) DownloadContract
```

平台及参与用户均可下载。

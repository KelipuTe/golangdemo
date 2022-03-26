package main

import (
  "database/sql"
  "errors"
  "fmt"

  _ "github.com/go-sql-driver/mysql"
  pkg_errors "github.com/pkg/errors"
)

// 第二周作业
// 我们在做数据库操作的时候，假设在 dao 层中遇到一个 sql.ErrNoRows ，是否应该
// Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

// sql.ErrNoRows，不应该包装。其他的 error ，比如 sql 错误，需要包装。
// sql.ErrNoRows 属于可以预见的 error，调用方需要处理而且必须立刻处理。
// 像 sql 错误这样的，属于不可预见或者说不应该出现，调用方无法处理也不应该处理，应该抛给上层。

// 实际使用时，一般出现 sql.ErrNoRows 的地方，大部分都是类似 First 这样的定向查询。
// 这意味着找得到或者找不到数据库记录，调用方都必须根据查询结果立刻做出判断。无论调用方是有补救措施，还是继续往上层抛 error。
// 但是一般的程序架构，应用层和数据层是严格分开的。这就意味着，一个查询调用，可能返回 sql.ErrNoRows，也可能返回 sql 错误。
// 为了避免两层之间因为错误类型出现强依赖耦合，数据层需要提供方法供应用层判断查询结果到底是哪一种：查到、没查到的 sql.ErrNoRows、还是其他 error。

func main() {
  Service()
}

// service 业务层
func Service() bool {
  p1dbModel, err := Repository(2)
  if nil != err {
    if !IsNoRows(err) {
      // 其他的 error，比如 sql 错误，从 api 请求层面来说是失败的。
      // 一般是进入报警环节，并给 api 调用者返回错误提示。
      fmt.Printf("query failed with error, %+v\r\n", err)
      return false
    }
    // 可预见的错误的业务流程，从 api 请求层面来说，可以认为是成功的（成功的预判了错误）。
    // 一般是给 api 调用者返回某种引导性提示。
    fmt.Printf("query failed with no rows, %+v\r\n", err)
    return true
  }
  // 正常的业务流程，给 api 调用者返回成功。
  fmt.Printf("query success, id=%d\r\n", p1dbModel.Id)
  return true
}

// ## 数据仓库层 ##

// dbModel 假装有一个数据模型
type dbModel struct {
  Id int
}

// repository 数据仓库层提供的接口
func Repository(queryId int) (*dbModel, error) {
  p1dbModel := &dbModel{}
  err := QueryFirst(p1dbModel, queryId) // 调用查询构造器
  if nil != err {
    return nil, handelQueryErr(err)
  }
  return p1dbModel, nil
}

// handelQueryErr 判断错误是否需要包装
func handelQueryErr(err error) error {
  if errors.Is(err, sql.ErrNoRows) {
    return err
  }
  return pkg_errors.WithStack(err)
}

// isNoRows 判断错误是不是可预见的错误，这里指 sql.ErrNoRows
// 封装成方法是为了解耦，业务层只需要知道是没查到数据，不需要知道是怎么知道的
func IsNoRows(err error) bool {
  return errors.Is(err, sql.ErrNoRows)
}

// #### 数据仓库层 ####

// ## 查询构造器 ##

// queryFirst 查询构造器只执行数据库相关的操作不判断错误类型，判断错误类型是数据仓库的事。
// 这里的 error 至少有三种混在一起：数据库连不上、sql 错误、sql.ErrNoRows。
// 其中 sql.ErrNoRows 属于可预见的错误，其他两种属于不可预见。
func QueryFirst(p1dbModel *dbModel, queryId int) error {
  p1db, err := sql.Open("mysql", makeConn())
  if nil != err {
    return err
  }
  defer p1db.Close()

  err = p1db.QueryRow("SELECT id FROM test where id = ?", queryId).Scan(&p1dbModel.Id)
  return err
}

func makeConn() string {
  var host, port, user, pass, name, charset string
  host = "127.0.0.1"
  port = "3306"
  user = "root"
  pass = "root"
  name = "test"
  charset = "utf8"
  return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", user, pass, host, port, name, charset)
}

// #### 查询构造器 ####

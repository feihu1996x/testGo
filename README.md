testGo
===

> Go语言编写单元测试

```bash
# example: run tests
go test main_test.go main.go -v -cover

# example: run test function in tests file
go test main_test.go main.go -v -cover -test.run TestAdd
go test main_test.go main.go mock_dao.go -v -cover -test.run TestMockSuccessService_Login
```

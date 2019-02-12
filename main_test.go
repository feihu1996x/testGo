package main

import (
    "testing"
)

func TestAdd(t *testing.T) {
    sum := Add(1, 2)
    if sum != 3 {
        t.Error("Add 1 and 2 result isn't 3")
    }
}

func TestMultiAdd(t *testing.T) {
    var tests = []struct {
        Arg1 int
        Arg2 int
        Sum int
    } {
        {1, 1, 2},
        {-1, -1, -2},
        {-1, 1, 0},
        {0, 1, 1},
    }
    for _, test := range tests {
        sum := Add(test.Arg1, test.Arg2)
        if sum != test.Sum {
            t.Errorf("Add %v and %v result isn't %v", test.Arg1, test.Arg2, test.Sum)
        }
    }
}

type StubDao struct {

}

var (
    rst []Data
    d = StubDao{}
)

func (d StubDao) ReadAll() []Data {
    return rst
}

func (d StubDao) SaveData(data *Data) {

}

/*
func TestService_Login(t *testing.T) {
    d := StubDao{}
    srv := NewService(d)
    rst := srv.Login("abc")
    if !rst {
        t.Error("Login error")
    }
}
*/

func TestSuccessService_Login(t *testing.T) {
    rst = []Data{{"abc"}}
    srv := NewService(d)
    rst := srv.Login("abc")
    if !rst {
        t.Error("Login Error")
    }
}

func TestFailService_login(t *testing.T) {
    rst = []Data{}
    srv := NewService(d)
    rst := srv.Login("abc")
    if rst {
        t.Error("Login Success")
    }
}

func TestSuccessReply(t *testing.T) {
    originLogin := Login
    defer func() { Login = originLogin }()
    Login = func (username, password string) bool {
        return true
    }
    if !Reply("a", "a", "aa") {
        t.Errorf("Reply false for login success")
    }
}

func TestFailReply(t *testing.T) {
    originLogin := Login
    defer func() { Login = originLogin }()
    Login = func(username, password string) bool {
        return false
    }
    if Reply("a", "a", "aa") {
        t.Errorf("Reply true for login fail")
    }
}

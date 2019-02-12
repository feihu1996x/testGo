package main

import "fmt"

type Data struct {
    Field string
}

type DAO interface {
     ReadAll() []Data
     SaveData(d *Data)
}

type MongoDao struct {

}

func (d MongoDao) ReadAll() []Data {
    return []Data{}
}

func (d MongoDao) SaveData(data *Data) {
    // ...
}

type Service struct {
    Dao *DAO
}

func (s *Service) Login(username string) bool {
    users := (*s.Dao).ReadAll()
    for _, user := range users {
        if username == user.Field {
            return true
        }
    }
    return false
}

func NewService(d DAO) *Service {
    srv := Service{ &d }
    return &srv
}

func Add(x, y int) int {
    return x + y
}

var Login = func (username, password string) bool {
    if username == password {
        return false
    }
    return true
}

func Reply(username, password, msg string) bool {
    if Login(username, password) {
        fmt.Println(msg)
        return true
    }
    return false
}

func main() {
    fmt.Println(Add(1, 2))

    d := MongoDao{}
    srv := NewService(d)
    fmt.Println(srv.Login("abc"))

    Reply("a", "a", "aa")
    Reply("a", "b", "ab")
}

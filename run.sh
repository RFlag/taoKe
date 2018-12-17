#!/bin/bash
##FTGO_SQL='root:123456@tcp(192.168.0.110:3306)/dataoke?charset=utf8&parseTime=True&loc=Local' go run -race main.go port.go pprof.go
FTGO_SQL='root:123456@tcp(mysql.pri:3306)/ranbb?charset=utf8&parseTime=True&loc=Local' go run -race main.go port.go pprof.go

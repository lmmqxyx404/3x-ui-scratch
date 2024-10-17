# 3x-ui-scratch

# used commands
`go mod init x-ui-scratch`
`go  run x-ui-scratch`

# go skills
## pay attention to initialize variable
`=` and `:=` are different.

## `init`
init 函数是很不一样的
第一次使用是在`web/session/session.go`中，注意缺少了这个初始化函数，就没法在session中调用更复杂的函数
package main

type INotify interface {
	genRandomCode(int) string                                        // 生成验证码
	sendVerifyCode(account, msg string, option ...interface{}) error // 发送验证码
	getCache(string) string                                          // 校验缓存
	setCache(string) error                                           // 缓存
	publicHooks()                                                    // 公共hook
}

// type Notify struct {
// 	url string
// }
//
// func NewNotify(url string) INotify {
// 	return &Notify{url}
// }
//
// func (n *Notify) genRandomCode(i int) string {
// 	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(int32(i)))
// 	return code
// }
//
// func (n *Notify) sendVerifyCode(account, msg string, option ...interface{}) error {
// 	panic("implement me")
// }
//
// func (n *Notify) getCache(s string) string {
// 	panic("implement me")
// }
//
// func (n *Notify) setCache(s string) error {
// 	panic("implement me")
// }
//
// func (n *Notify) publicHooks() {
// 	panic("implement me")
// }

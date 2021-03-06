package impl

import (
	"go-gin/internal/app/bll"
	"go-gin/internal/app/bll/impl/internal"
	"go.uber.org/dig"
)

// Inject 注入bll实现
// 使用方式：
//   container := dig.New()
//   Inject(container)
//   container.Invoke(func(foo IDemo) {
//   })
func Inject(container *dig.Container) error {
	{
		container.Provide(internal.NewTrans, dig.As(new(bll.ITrans)))
		container.Provide(internal.NewDemo, dig.As(new(bll.IDemo)))
	}
	return nil
}

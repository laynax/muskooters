package framework

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"muskooters/services/config"
	"muskooters/services/initializer"
)

var (
	all []Routes
)

type Routes interface {
	Routes(*gin.Engine)
}

type initer struct {
}

func (i *initer) Initialize() func() {
	port := config.MustString("PORT")
	e := gin.New()

	for i := range all {
		all[i].Routes(e)
	}

	go func() {
		err := e.Run(port)
		logrus.Errorln("[framework]", err)
	}()

	return nil
}

// Register a new controller class
func Register(c ...Routes) {
	all = append(all, c...)
}

func init() {
	initializer.Register(&initer{})
}
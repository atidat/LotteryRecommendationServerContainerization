package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"net/http"
)




func main() {
	r := gin.Default()
	r.GET("/api/v1/pods", getPods())
	rus := logrus.New()
	rus.Infof("starting container server...")
	_ = r.Run()
}

func getPods() gin.HandlerFunc {
	return func(con *gin.Context) {
		k8sCfg := &rest.Config{}
		cliS, err := kubernetes.NewForConfig(k8sCfg)
		if err != nil {
			errorHandle(http.StatusInternalServerError, "rest config init failed", con)
		}
		pods, err := cliS.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})
		if err != nil {
			errorHandle(http.StatusInternalServerError, "list pods failed", con)
		}
		logrus.Info(pods)
		con.JSON(http.StatusOK, "")
	}
}

func errorHandle(status int, msg string, con *gin.Context) {
	logrus.Error(msg)
	con.JSON(status, msg)
}
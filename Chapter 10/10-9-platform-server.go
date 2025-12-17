package chapter10

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	kubevirtcorev1 "kubevirt.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

type HTTPServer struct {
	Router     *gin.Engine
	dispatcher *Dispatcher
	client     client.Client
}

type Dispatcher struct {
	Client client.Client
	// other clients or useful libraries etc.
}

func (d *Dispatcher) GetVMs(c *gin.Context) {
	vms := &kubevirtcorev1.VirtualMachineList{}
	project := c.Param("project")

	err := d.Client.List(c.Request.Context(), vms, &client.ListOptions{
		Namespace: project,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "error listing VMs",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"vms": vms.Items,
	})
}

func NewDispatcher(k8s client.Client, router *gin.Engine) *Dispatcher {
	d := &Dispatcher{
		Client: k8s,
		// other clients or useful libraries etc.
	}

	projv1 := router.Group("/api/v1")
	// projv1.POST("/orgs/:orgSlug/projects/:projectSlug/vms", d.CreateVM)
	projv1.GET("/orgs/:orgSlug/projects/:projectSlug/vms", d.GetVMs)

	return d
}

func NewHTTPServer(scheme *runtime.Scheme) *HTTPServer {
	r := gin.Default()
	c, err := client.New(config.GetConfigOrDie(), client.Options{
		Scheme: scheme,
	})
	if err != nil {
		panic(err)
	}
	s := &HTTPServer{
		Router:     r,
		dispatcher: NewDispatcher(c, r),
		client:     c,
	}
	return s
}

// launch the HTTP server
func Run(addr string) error {
	fmt.Println("starting webserver")
	s := NewHTTPServer(scheme)
	err := s.Router.Run("0.0.0.0:3000")
	if err != nil {
		klog.Error(err, "unable to start server")
		panic(err)
	}
	return nil
}

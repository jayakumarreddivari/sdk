package test


import (
	"fmt"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/avinetworks/sdk/go/models"
	"github.com/avinetworks/sdk/go/session"
	"testing"
)

func TestCreateHealthmonitor(t *testing.T) {
	aviClient, err := clients.NewAviClient("10.10.28.91", "admin",
		session.SetPassword("avi123"),
		session.SetTenant("webapp"),
		session.SetVersion("17.2.8"),
		session.SetInsecure)
	if err != nil {
		fmt.Println("Couldn't create session: ", err)
		t.Fail()
	}
	cv, err := aviClient.AviSession.GetControllerVersion()
	fmt.Printf("Avi Controller Version: %v:%v\n", cv, err)

	// Create health monitor in webapp tenant
	hmobj := models.HealthMonitor{}
	hmobj.Name = "Test-Hm"
	hmobj.Type = "HEALTH_MONITOR_HTTP"
	hmobj.ReceiveTimeout = 2
	hmobj.SendInterval = 3
	hmobj.SuccessfulChecks = 10
	hmobj.TenantRef = "/api/tenant?name=webapp"
	httpmonitor := models.HealthMonitorHTTP{}
	httpmonitor.ExactHTTPRequest = false
	httpmonitor.HTTPRequest = "HEAD / HTTP/1.0"
	httpmonitor.HTTPResponseCode = append(httpmonitor.HTTPResponseCode, "HTTP_3XX")
	hmobj.HTTPMonitor = &httpmonitor
	nvsobj, err := aviClient.HealthMonitor.Create(&hmobj)
	if err != nil {
		fmt.Println("\n Healthmonitor creation failed: ", err)
		t.Fail()
	}
	fmt.Printf("\n Healthmonitor obj: %+v", *nvsobj)

	// Update healthmonitor
	profobj:= models.HealthMonitor{}
	err = aviClient.AviSession.GetObjectByName("healthmonitor", "Test-Hm", &profobj)
	if err == nil {
		profobj.Name = "Test-Healthmonitor"
		profobj.ReceiveTimeout = 3
		profobj.SendInterval = 4
		profobj.SuccessfulChecks = 10
		profobj.Type = "HEALTH_MONITOR_HTTP"

		upObj , err := aviClient.HealthMonitor.Update(&profobj)
		if err != nil {
			fmt.Println("\n Healthmonitor Updation failed: ", err)
			t.Fail()
		}
		fmt.Printf("\n Healthmonitor obj: %+v", *upObj)
	}

}
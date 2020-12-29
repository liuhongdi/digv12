package global

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"log"
	"os"
)

var (
	Enforcer *casbin.Enforcer
)

//创建casbin的enforcer
func SetupCasbinEnforcer() (error) {

	a, err := gormadapter.NewAdapterByDB(DBLink)
	if (err != nil) {
		log.Fatalf("gormadapter.NewAdapterByDB err: %v", err)
	}
	dir, _ := os.Getwd()
	modelPath := dir + "/config/rbac_model.conf"
	//csvPath := dir + "/config/rbac2.csv"
	fmt.Println("modelPath:"+modelPath);
	//fmt.Println("csvPath:"+csvPath);
	var errC error
	Enforcer, errC = casbin.NewEnforcer(modelPath, a)
	//fmt.Printf("RBAC test start\n") // output for debug
	if (errC != nil) {
		//fmt.Println(errC)
		log.Fatalf("SetupCasbinEnforcer err: %v", errC)
		return errC
	} else {
		Enforcer.LoadPolicy()
		Enforcer.EnableLog(true)
		return nil
	}
}

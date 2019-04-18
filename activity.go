package readredisvar_flogo

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/go-redis/redis"
)
var log = logger.GetLogger("readredisvar_log")
// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error)  {

	host := context.GetInput("host").(string)
	port := context.GetInput("port").(string)
	dbNo := context.GetInput("dbNo").(int)
	variable := context.GetInput("variable").(string)
	log.Infof("Connecting to Redis: [%s]:[%s]", host, port)
	log.Infof("DB no: [%s]",dbNo)
	log.Infof("Variable: [%s]",variable)

	client := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
	val, err := client.Get(variable).Result()
	if err != nil {
		panic(err)
	}
	log.Infof("Result: [%s]",val)
	context.SetOutput("value", val)
	client.Close()
	return true, nil
}

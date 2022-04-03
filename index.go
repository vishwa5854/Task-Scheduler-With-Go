// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"strconv"
// 	"time"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/kamva/mgm/v3"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// func initDataBaseConnection() {
// 	fmt.Println("Initializing database connection")
// 	err := mgm.SetDefaultConfig(nil, "kiotapp", options.Client().ApplyURI(""))

// 	if err == nil {
// 		fmt.Println("Database connection initialized")
// 	} else {
// 		fmt.Println("Database connection failed")
// 	}
// }

// type Job struct {
// 	mgm.DefaultModel `bson:",inline"`
// 	Name             string `json:"name" bson:"name"`
// 	Time             int64  `json:"time" bson:"time"`
// 	Status           bool   `json:"status" bson:"status"`
// }

// func NewJob(name string, time int64) *Job {
// 	return &Job{
// 		Name:   name,
// 		Time:   time,
// 		Status: false,
// 	}
// }

// func main() {
// 	initDataBaseConnection()

// 	app := fiber.New()

// 	/* Create new task with a given name and time. */
// 	app.Post("/create", func(c *fiber.Ctx) error {
// 		name := c.Query("name")
// 		time, err := strconv.ParseInt(c.Query("time"), 10, 64)

// 		fmt.Println(time)

// 		if name == "" || err != nil || time == 0 {
// 			return c.Status(400).JSON(fiber.Map{
// 				"message": "Missing name or time",
// 			})
// 		}

// 		job := NewJob(name, time)
// 		err = mgm.Coll(job).Create(job)

// 		if err != nil {
// 			return c.Status(500).JSON(fiber.Map{
// 				"message": "Error creating job",
// 			})
// 		}

// 		// fmt.Println(reflect.TypeOf(job.id))

// 		return c.JSON(fiber.Map{
// 			"message": "Task created",
// 			"data":    job,
// 		})
// 	})

// 	app.Get("/job", func(c *fiber.Ctx) error {
// 		/* If nothing is passed then we return all tasks active and inactive. */
// 		id, _ := strconv.ParseInt(c.Query("id"), 10, 64)

// 		fmt.Println(id)
// 		fmt.Println(c.Query("id"))

// 		var jobs any
// 		job := &Job{}
// 		coll := mgm.Coll(job)

// 		jobs = coll.FindByID(c.Query("id"), job)

// 		return c.JSON(fiber.Map{
// 			"jobs": jobs,
// 		})
// 	})

// 	jobs := []string{}

// 	app.Post("/", func(c *fiber.Ctx) error {

// 	})

// 	app.Listen(":3000")
// }

// func convertMilliSecondsToSecondsFromNow(time int64) int64 {
// 	return time / 1000
// }

// func getCurrentTimeInMilliSeconds() int64 {

// }

// /*
// 	1. Mongo DB connection
// 	2. Expose create, read, update, delete apis basic ones with no functionality yet
// 	3. On server start fetch all the latest events from the database that are not done yet
// 	4. Mongo doc has _id which will be used to identify the mongo routine, once executed the mongo doc
// 		will be updated with the status of the routine
// 	5. Edit the time at which it should be ran in future
// 		1. First we will fetch the mongo db and see the status of it, if it is already executed we will spin
// 		up another sub routine with the new time
// 		2. If it is not yet executed we will send the _id - SIGTERM in the shared channel and that will kill the existing
// 		routine and then we will create a new routine with the time
// 	6. Delete will be same as edit only we will delete the routine
// 	7. Retrieve will get all the currently running routines from the mongo db
// */

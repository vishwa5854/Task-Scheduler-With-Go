package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Task struct {
	Name   string
	Time   int64
	Status bool
}

func createNewTask(name string, time int64, status bool) *Task {
	return &Task{
		Name:   name,
		Time:   time,
		Status: false,
	}
}

var tasks = make(map[string]*Task)
var mainChannel = make(chan map[string]*Task)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("Hello World")
		return c.JSON(fiber.Map{
			"tasks": tasks,
		})
	})

	app.Post("/", func(c *fiber.Ctx) error {
		taskName := c.Query("name")

		if tasks[taskName] != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Task already exists",
			})
		}

		taskTime, _ := strconv.ParseInt(c.Query("time"), 10, 64)

		task := createNewTask(taskName, taskTime, false)

		tasks[taskName] = task

		go taskExecutor(task)

		return c.SendString("Task successfully created bruh")
	})

	app.Patch("/", func(c *fiber.Ctx) error {
		taskName := c.Query("name")
		taskTime, _ := strconv.ParseInt(c.Query("time"), 10, 64)

		tasks[taskName].Time = taskTime

		go taskExecutor(tasks[taskName])

		return c.SendString("Task successfully updated bruh")
	})

	app.Delete("/", func(c *fiber.Ctx) error {
		taskName := c.Query("name")

		tasks[taskName].Time = 0

		return c.SendString("Task successfully deleted bruh")
	})

	app.Listen(":3000")
}

func taskExecutor(task *Task) {
	fmt.Println("Initiated task : " + task.Name)

	currentTime := time.Now().UnixMilli()

	originalTask := task

	if currentTime < task.Time {
		time.Sleep(time.Duration(task.Time-currentTime) * time.Millisecond)

		if originalTask.Time == tasks[originalTask.Name].Time {
			fmt.Println("Running task : " + task.Name)
			CallGoogle()
			tasks[task.Name].Status = true
			fmt.Println("Task Done : " + task.Name)
		} else {
			return
		}
	} else {
		fmt.Println("Task already timed out")
	}
}

func CallGoogle() {
	url := "https://www.google.com"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		errorHandler(err)
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		errorHandler(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
		fmt.Println(body)
		errorHandler(err)
	}
}

func errorHandler(err any) {
	fmt.Println("Error in calling google.com")
	fmt.Println(err)
}

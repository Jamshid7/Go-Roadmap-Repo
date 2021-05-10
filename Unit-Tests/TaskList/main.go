package TaskList
import "fmt"
type TaskList struct {
    TaskName string
    TaskDate string
}

func (t TaskList) GetTask(singleTask, singleDate string) string{
	t.TaskName = singleTask
	t.TaskDate = singleDate
	return t.TaskName + t.TaskDate	  
}

//just to check it here also
func main() {
	t := new(TaskList)
	res := t.GetTask("Submit Lab5 at", " 23.55")
	fmt.Println(res)
}

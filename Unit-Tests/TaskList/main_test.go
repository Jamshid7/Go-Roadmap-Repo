package TaskList
import ("testing")

func TestTaskList(t *testing.T){
	task := new(TaskList)
	res := task.GetTask("Submit Lab5 at", " 23.55")
	res1 := task.GetTask("Submit Quiz before", " 08.00")
	res2 := task.GetTask("Submit HW before", " 23.59")
	if(res != "Submit Lab4 at 23.55") {
		t.Error("Expected Submit Lab5 at 23.55 but got error")
	}
	if(res1 != "Submit Quiz before 08.00") {
		t.Error("Expected Submit Quiz before 08.00 but got error")
	}
	if(res2 != "Submit HW before 23.59") {
		t.Error("Expected Submit HW before 23.59 but got error")
	}
}

package model

import (
	"encoding/json"
	"fmt"
)

type stuMgr struct {
	students map[string]student
}

func NewstuMgr() stuMgr {
	return stuMgr{
		students: make(map[string]student),
	}
}

func (sm *stuMgr) AddStudent(name string, s student) {
	sm.students[name] = s
}

func (sm *stuMgr) DelStuent(name string) {

	delete(sm.students, name)
}

func (sm stuMgr) ShowStuent() {
	var output string
	for name, student := range sm.students {
		s, err := json.Marshal(student)
		if err != nil {
			fmt.Printf("%s\t%v", name, err)
		}

		output += fmt.Sprintf("%s\t%s\n", name, string(s))
	}
	fmt.Printf(output)
}

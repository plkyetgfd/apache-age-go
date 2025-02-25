package age

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type VPerson struct {
	Name   string
	Age    int64
	Weight float64
}

type EWorkWith struct {
	Weight int64
}

func TestPathMapping(t *testing.T) {
	rstStr1 := `[{"id": 2251799813685425, "label": "Person", "properties": {"name": "Smith"}}::vertex, 
	{"id": 2533274790396576, "label": "workWith", "end_id": 2251799813685425, "start_id": 2251799813685424, "properties": {"weight": 3}}::edge, 
	{"id": 2251799813685424, "label": "Person", "properties": {"name": "Joe"}}::vertex]::path`

	rstStr2 := `[{"id": 2251799813685424, "label": "Person", "properties": {"name": "Joe"}}::vertex, 
	{"id": 2533274790396576, "label": "workWith", "end_id": 2251799813685425, "start_id": 2251799813685424, "properties": {"weight": 3}}::edge, 
	{"id": 2251799813685425, "label": "Person", "properties": {"name": "Smith"}}::vertex]::path`

	rstStr3 := `[{"id": 2251799813685424, "label": "Person", "properties": {"name": "Joe"}}::vertex, 
	{"id": 2533274790396579, "label": "workWith", "end_id": 2251799813685426, "start_id": 2251799813685424, "properties": {"weight": 5}}::edge, 
	{"id": 2251799813685426, "label": "Person", "properties": {"name": "Jack"}}::vertex]::path`

	mapper := NewAGMapper(nil)
	mapper.PutType("Person", reflect.TypeOf(VPerson{}))
	mapper.PutType("workWith", reflect.TypeOf(EWorkWith{}))

	entity1, _ := mapper.unmarshal(rstStr1)
	entity2, _ := mapper.unmarshal(rstStr2)
	entity3, _ := mapper.unmarshal(rstStr3)

	fmt.Println(" **** ", entity1)
	fmt.Println(" **** ", entity2)
	fmt.Println(" **** ", entity3)

	assert.Equal(t, entity1.GType(), entity2.GType(), "Type Check")
	p1 := entity1.(*MapPath)
	p2 := entity2.(*MapPath)
	p3 := entity3.(*MapPath)

	assert.Equal(t, p1.Get(2).(VPerson).Name, p2.Get(0).(VPerson).Name)
	assert.Equal(t, p2.Get(0).(VPerson).Name, p3.Get(0).(VPerson).Name)

	assert.Equal(t, p1.Get(1).(EWorkWith).Weight, int64(3))
	assert.Equal(t, p2.Get(1).(EWorkWith).Weight, int64(3))
	assert.Equal(t, p3.Get(1).(EWorkWith).Weight, int64(5))
}

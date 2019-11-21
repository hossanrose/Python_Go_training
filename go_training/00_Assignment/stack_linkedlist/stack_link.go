package main
import (
 "fmt"
)

type person struct {
  name    string
  age       int
  counter   int
  next      *person
}

type myStack interface {
    removeNodeEnd() *person
    addNodeEnd(*person) *person
  //  setSize (int)
}



func main() {
  //NONparallel
  mike := &person{"mike",33, 1, nil}
  personList := mike
  greg := &person{"greg", 443, 2, nil}
  //james := &person{45, 554, 3, nil}
  //marshall := &person{56, 123, 4, nil}
  personList = stack_push(greg, personList)
  //personList = stack_push(james, personList)
  //personList = stack_push(marshall, personList)
  printList(personList)
  personList = stack_pop(personList)
  printList(personList)
  //Parallel

  ch :=make(chan *person)
  num :=5

  for i:=0;i<num;i++{
    go producer_linkedlist(ch,i)
    //time.Sleep(1*time.Second)
    //go consumer_linkedlist(ch,i)
    //time.Sleep(1 * time.Second)
 }


}

func producer_linkedlist(ch chan *person, id int){

}
func printList(personList *person) {
  for p := personList; p != nil; p = p.next {
   fmt.Println(p)
  }
}

func stack_push(newPerson *person, personList myStack ) *person{
  return personList.addNodeEnd(newPerson)
}
func stack_pop(personList myStack ) *person{
  return personList.removeNodeEnd()
}

func (personList *person) addNodeEnd(newPerson *person) *person {
  if personList == nil {
    return personList
  }
  for p := personList; p != nil; p = p.next {
    if p.next == nil {
      p.next = newPerson
      return personList
    }
   }
  return personList
}
func (personList *person) removeNodeEnd() *person {
  if personList == nil {
    return personList
  }
  for p := personList; p != nil; p = p.next {
    if p.next==nil{
      // Need to implement removal of headnode
      return personList
    }
    if p.next.next == nil {
      p.next= nil
      return personList
    }
   }
  return personList
}

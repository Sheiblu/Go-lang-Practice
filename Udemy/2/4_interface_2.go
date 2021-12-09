package main

import ("fmt")

/* define an interface */
type People interface {
   getName() string
   getAge() int
}

/* define a circle */
type Man struct {
  Name string
  Age int
}

/* define a rectangle */
type Woman struct {
   Name string
   Age int
}

// for man
func(m Man) getName() string {
   return m.Name
}

func(m Man) getAge() int {
   return m.Age
}

// for woman
func(w Woman) getName() string {
   return w.Name
}

// func(w Woman) getAge() string {
//    return w.Age
// }

func Get_Man (name string, age int) People {
   return Man {name, age}
}

func Get_Woman (name string, age int) People {
   return Woman {name, age}
}



func main() {

   /*
   // General time
   sheiblu := Man {"Sheiblu", 25}
   israt := Woman {"israt", 25}
   
   fmt.Println("Man : ", sheiblu.getName())
   fmt.Println("Woman : ", israt.getName())
   fmt.Println("Man Age: ", sheiblu.getAge())
   fmt.Println("Woman Age: ", israt.getAge())  // Get a error.  for this reason we use interface
   */

   sheiblu := Get_Man("Sheiblu", 25)
   fmt.Println("Man : ", sheiblu.getName())

   israt := Get_Woman("israt", 25)  //  this give a error that u r not implement  interface  getAge()
   fmt.Println("Woman : ", israt.getName())  


}
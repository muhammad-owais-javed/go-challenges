package sprint

func FactorialIterative(n int) int {

 if n <= 1 && n >= 0  {
  return 1
 }
 
 if n < 0 {
  return 0
}
 return n * FactorialIterative(n)

}

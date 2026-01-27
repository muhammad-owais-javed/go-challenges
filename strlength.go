package sprint

func StrLength(s string) []int {

 var counter int = 0
 res := make([]int, 0)

 for _, r := range s {
  counter += int(r-r)+1
 }

 res = append(res, counter)
 res = append(res, len(str))

 return res
  //fmt.Println(counter)
}

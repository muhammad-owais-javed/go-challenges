package sprint

func Doop(a int, op string, b int) int {

  switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "/":
	 if b == 0 || a == 0{
		return 0
		}
		return a / b
	case "*":
		return a * b
	case "%":
	 if b == 0 {
		return 0
	}
			
	 var modulo = a % b
	 	return modulo

	default:
		return 0
	}
 }
}

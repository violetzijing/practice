package practice

func main() {

}

func gcd(p int, q int) int {
  if q == 0 {
    return p
  }
  r = math.Mod(p, q)
  
  return gcd(q, r)
}

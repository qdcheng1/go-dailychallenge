func climbStairs(n int) int {
    if n < 3 {
        return n
    }
    var a, b int = 1,2
    
    for i := 3; i <= n; i++ {
        var c int = a + b
        a = b
        b = c
    }
    return b
}

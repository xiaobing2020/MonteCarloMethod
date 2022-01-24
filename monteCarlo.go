// 蒙特卡罗算法go语言实现，及两个特例
// 特例1: 通过圆方程计算pi
// 特例2：计算定积分 y = x^2  在[0,1]区间的面积(积分表达式没法复制过来)
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func circle(x, y float64) float64 {
    return (x - 0.5) * (x - 0.5) + (y - 0.5) * (y - 0.5)
}

func square(x, y float64) float64 {
    return x * x - y
}

// 蒙特卡罗算法go语言实现
// 参数说明: left: 函数fn的左区间
//          right: 函数fn的右区间
//          fn: 曲线方程 f(x,y) 例如: 1)  f(x, y) = (x - 0.5)^2 + (y - 0.5)^2   2) f(x, y) = x^2 - y
//          upper: 函数fn的取值上界 普通函数该值传0, 圆或曲线
func monteCarlo(left, right float64, fn func(x, y float64) float64, upper float64) (int, int) {
    var x, y float64
    var m, n int
    rand.Seed(time.Now().Unix())
    for i := 0; i < 10000000; i++ { // 样本容量100万个 为了精确度可以更大一点
        x = rand.Float64() * (right - left)
        y = rand.Float64() * (right - left)
        if fn(x, y) < upper {
            m++
        } else {
            n++
        }
    }

    return m, n
}

func main() {
    m, n := monteCarlo(0, 1, circle, 0.25)
    fmt.Println("m = ",m, "n = ", n)
    pi := (4 * float64(m)) / float64(m + n)
    fmt.Println("PI is: ",pi)
    m, n = monteCarlo(0, 1, square, 0.0)
    fmt.Println("m = ",m, "n = ", n)
    s := float64(n) / float64(m + n)
    fmt.Println("y = x^2 在 0，1的面积是: ", s)

    //output:
    //m =  7854396 n =  2145604
    //PI is:  3.1417584
    //m =  6667086 n =  3332914
    //y = x^2 在 0，1的面积是:  0.3332914
}

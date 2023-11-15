package main
import (
    "testing"
    "time"
)
func BenchmarkLocation(b *testing.B) {
    for n := 0; n < b.N; n++ {
        loc, _ := time.LoadLocation("Asia/Kolkata")
        time.Now().In(loc)
    }
}
func BenchmarkLocation2(b *testing.B) {
    loc, _ := time.LoadLocation("Asia/Kolkata")
    for n := 0; n < b.N; n++ {
        time.Now().In(loc)
    }
}
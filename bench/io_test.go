package bench

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func BenchmarkWriteString(b *testing.B){
	for i := 0 ; i < b.N ; i++ {
		io.WriteString(ioutil.Discard,"hogehoge")
	}
}

func BenchmarkFmtFprint(b *testing.B){
	for i := 0 ; i < b.N ; i++ {
		fmt.Fprint(ioutil.Discard,"hogehoge")
	}
}

func BenchmarkWriteWithStr2Byte(b *testing.B){
	for i := 0 ; i < b.N ; i++ {
		ioutil.Discard.Write( []byte("hogehoge") )
	}
}

func BenchmarkWrite(b *testing.B){
	hogehoge := []byte("hogehoge")
	for i := 0 ; i < b.N ; i++ {
		ioutil.Discard.Write( hogehoge )
	}
}

func BenchmarkWriteStringReal(b *testing.B){
	fd,_ := os.Create("nul")
	fd.Close()
	for i := 0 ; i < b.N ; i++ {
		io.WriteString(fd,"hogehoge")
	}
}

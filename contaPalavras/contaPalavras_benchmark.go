package main

import (
	"testing"
)

func BenchmarkContaPalavras(b *testing.B) {
	for i := 0; i < b.N; i++ {
		contaPalavras(`Lorem ipsum dolor sit amet, consectetur adipiscing elit, 
					   sed do eiusmod tempor incididunt ut labore et dolore magna 
					   aliqua. Ac turpis egestas sed tempus urna et pharetra. Duis 
					   at consectetur lorem donec massa sapien faucibus. Viverra ipsum 
					   nunc aliquet bibendum enim. Dui accumsan sit amet nulla facilisi 
					   morbi tempus iaculis. Blandit volutpat maecenas volutpat blandit 
					   aliquam etiam erat. Augue ut lectus arcu bibendum at. Pharetra diam 
					   sit amet nisl suscipit adipiscing bibendum. Pharetra diam sit amet 
					   nisl suscipit adipiscing bibendum est ultricies. Dolor sit amet consectetur 
					   adipiscing elit pellentesque.`)
	}
}

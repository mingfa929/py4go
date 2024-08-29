package python

// See:
//  https://docs.python.org/3/c-api/
//  https://github.com/golang/go/wiki/cgo
//  https://www.datadoghq.com/blog/engineering/cgo-and-python/
//  https://github.com/sbinet/go-python

// #cgo pkg-config: python3-embed
// #cgo CFLAGS: -I/Users/bytedance/apps/miniforge3/envs/py312/include/python3.12
// #cgo LDFLAGS: -L/Users/bytedance/apps/miniforge3/envs/py312/lib -Wl,-rpath,/Users/bytedance/apps/miniforge3/envs/py312/lib -lpython3.12 
import "C"

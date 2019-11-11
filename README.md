# ortoolslp
Go package to use the 'linear_solver' component of Google OR-Tools

It enables you to solve linear programming models with GLOP, CLP and CBC solvers via the Google OR-Tools API on Windows. It was tested on Windows 10.
The compilation process is quite complex because it involves C++, SWIG, C and many tools more. To simploy the use of the library this repository only contains the minimum you need to run and develop applications with ortoolslp.

The original code stems from github.com/google/or-tools and was forked to github.com/baobabsoluciones/or-tools to work on an API that can be used with Go.

Requirements

	- Go
	- MinGW64 C++ compiler(this package was tested with TDM GCC, an installer can be found on http://tdm-gcc.tdragon.net/)

Install it by running

	go get -u github.com/baobabsoluciones/ortoolslp

This command will install ortoolslp in your go path and you can use it in your Go program by writing

	import "github.com/baobabsoluciones/ortoolslp"

You can compile any program using the package now. Running it requires one more step. In the repository now in the go path there is a DLL that the program needs to start. On your system you can probably find it under

	C:\Users\YourUserName\go\src\github.com\baobabsoluciones\ortoolslp\_gowraplp.dll

You need to put this DLL in a place where Windows can find it. Windows searches your system in all directories given in the Path variable and in the directory where you placed your programs EXE file. It is up to you where you copy the DLL, it doesn't matter. If you install your compiled program on a customers system, you can put the EXE and the DLL in the same directory. On your system, maybe you have multiple programs that need the DLL, then you only want to copy it into one place. A good place is

	C:\Users\YourUserName\go\bin

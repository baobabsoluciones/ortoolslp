# ortoolslp
Go package to use the 'linear_solver' component of Google OR-Tools on Windows 10.

It enables you to solve linear programming models with GLOP, CLP, CBC, SCIP, GLPK, CPLEX and Gurobi solvers via the Google OR-Tools API on Windows. It was tested on Windows 10 with x86_64 architecture.
The compilation process is quite complex because it involves C++, SWIG, C and many tools more. To simplify the use of the library, this repository only contains the minimum binaries and Go-files you need to run and develop applications with ortoolslp.

The original code stems from [github.com/google/or-tools](https://github.com/google/or-tools) and was forked to [github.com/baobabsoluciones/or-tools](https://github.com/baobabsoluciones/or-tools) to work on an API that can be used with Go.

## Requirements

	- Go
	- MinGW64 C and C++ compiler

Go is not compatible with Microsoft compilers. If you don't have a MinGW64 C and C++ compiler installed already, we recommend the "TDM64 bundle". Download the "TDM64 bundle" [here](https://sourceforge.net/projects/tdm-gcc/files/latest/download) or via the [official website](http://tdm-gcc.tdragon.net/).

## Installation

Install the ortoolslp package by running

	go get -u github.com/baobabsoluciones/ortoolslp

This command will install ortoolslp in your GOPATH and you can use it in your Go program by writing

	import "github.com/baobabsoluciones/ortoolslp"

Now you can compile any program that uses the package. Running it requires one more step. In the repository, now in your GOPATH, there is a DLL that the program needs to start. On your system you can probably find it under

	C:\Users\YourUserName\go\src\github.com\baobabsoluciones\ortoolslp\_gowraplp.dll
	
This DLL contains all the code to solve models with GLOP, CBC and CLP.
Additionally in the same directory you will find a DLL named glpk*.dll where the star is a version number. You need this DLL aswell if you intend to use the GLPK solver. The necessary DLLs for CPLEX, Gurobi and SCIP are not included because those are proprietary and require a license.
This version of ortoolslp was linked to cplex1290.dll, gurobi81.dll, scip.dll and glpk_4_65.dll. Since they are only delay-loaded you don't need these DLLs as long as you don't try to use the respective solvers.

You need to put at least \_gowraplp.dll in a place where Windows can find it. Windows searches your system in all directories given in the Path environment variable and in the directory where you placed your programs EXE file. It is up to you where you copy the DLLs you need, it doesn't matter. The installers of Gurobi, CPLEX and SCIP update your Path automatically.
If you install your compiled program on a customers system, you can put the EXE and the DLLs in the same directory. This will work with all solvers except Gurobi, it has a more complicated licensing mechanism and it is recommended to use th installer on every system you need Gurobi. On your own development system, maybe you have multiple programs that need these DLLs, then you only want to copy it into one place that is in your Path. A good place is

	C:\Users\YourUserName\go\bin

Editing system environment variables on Windows is explained [here](https://docs.oracle.com/en/database/oracle/r-enterprise/1.5.1/oread/creating-and-modifying-environment-variables-on-windows.html).

## Troubleshooting

\_gowraplp.dll, cplex1290.dll, scip.dll and gurobi81.dll depend on DLLs that are commonly included on most Windows installations today.

If you get errors that mention ucrtbase.dll or MSVCP140.dll you should either update your Windows system or install ["Microsoft Visual C++ 2015 Redistributable Update 3 RC"](https://www.microsoft.com/en-us/download/details.aspx?id=52685).

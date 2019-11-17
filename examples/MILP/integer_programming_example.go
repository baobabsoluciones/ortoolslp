// Copyright 2010-2018 Google LLC
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This is a Go replica of https://github.com/google/or-tools/blob/master/ortools/linear_solver/samples/integer_programming_example.py
// Small example to illustrate solving a MIP problem.
package main

import (
	"fmt"

	"github.com/baobabsoluciones/ortoolslp"
)

func main() {
	// Integer programming sample.
	// [START solver]
	// Create the mip solver with the desired backend
	 solver := ortoolslp.NewSolver("IntegerProgrammingExample", ortoolslp.SolverCBC_MIXED_INTEGER_PROGRAMMING)
	// solver := ortoolslp.NewSolver("IntegerProgrammingExample", ortoolslp.SolverCLP_LINEAR_PROGRAMMING)
	// solver := ortoolslp.NewSolver("IntegerProgrammingExample", ortoolslp.SolverGLOP_LINEAR_PROGRAMMING)
	// solver := ortoolslp.NewSolver("IntegerProgrammingExample", ortoolslp.SolverGLPK_LINEAR_PROGRAMMING)
	// solver := ortoolslp.NewSolver("IntegerProgrammingExample", ortoolslp.SolverSCIP_MIXED_INTEGER_PROGRAMMING)
	// solver := ortoolslp.NewSolver("IntegerProgrammingExample", ortoolslp.SolverCPLEX_LINEAR_PROGRAMMING)
	// solver := ortoolslp.NewSolver("IntegerProgrammingExample", ortoolslp.SolverGUROBI_LINEAR_PROGRAMMING)
	// [END solver]

	// [START variables]
	// x, y, and z are non-negative integer variables.
	x := solver.IntVar(0.0, ortoolslp.SolverInfinity(), "x")
	y := solver.IntVar(0.0, ortoolslp.SolverInfinity(), "y")
	z := solver.IntVar(0.0, ortoolslp.SolverInfinity(), "z")
	// [END variables]

	// [START constraints]
	// 2*x + 7*y + 3*z <= 50
	constraint0 := solver.Constraint(-ortoolslp.SolverInfinity(), float64(50))
	constraint0.SetCoefficient(x, 2)
	constraint0.SetCoefficient(y, 7)
	constraint0.SetCoefficient(z, 3)

	// 3*x - 5*y + 7*z <= 45
	constraint1 := solver.Constraint(-ortoolslp.SolverInfinity(), float64(45))
	constraint1.SetCoefficient(x, 3)
	constraint1.SetCoefficient(y, -5)
	constraint1.SetCoefficient(z, 7)

	// 5*x + 2*y - 6*z <= 37
	constraint2 := solver.Constraint(-ortoolslp.SolverInfinity(), float64(37))
	constraint2.SetCoefficient(x, 5)
	constraint2.SetCoefficient(y, 2)
	constraint2.SetCoefficient(z, -6)
	// [END constraints]

	// [START objective]
	// Maximize 2*x + 2*y + 3*z
	objective := solver.Objective()
	objective.SetCoefficient(x, 2)
	objective.SetCoefficient(y, 2)
	objective.SetCoefficient(z, 3)
	objective.SetMaximization()
	// [END objective]

	// Solve the problem and print the solution.
	// [START print_solution]
	solver.Solve()
	// Print the objective value of the solution.
	fmt.Printf("Maximum objective function value = %f\n", solver.Objective().Value())
	fmt.Println()
	// Print the value of each variable in the solution.
	fmt.Printf("%s = %f\n", x.Name(), x.Solution_value())
	fmt.Printf("%s = %f\n", y.Name(), y.Solution_value())
	fmt.Printf("%s = %f\n", z.Name(), z.Solution_value())
	// [END print_solution]
}

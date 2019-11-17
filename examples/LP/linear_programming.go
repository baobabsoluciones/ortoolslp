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

// This example is a Go replica of https://github.com/google/or-tools/blob/master/ortools/linear_solver/samples/linear_programming_example.py
//Linear optimization example
package main

import (
	"fmt"
	"os"

	"github.com/baobabsoluciones/ortoolslp"
)

func main() {
	//Entry point of the program
	//Instantiate a solver, naming it LinearExample.
	// solver := ortoolslp.NewSolver("LinearExample", ortoolslp.SolverCBC_MIXED_INTEGER_PROGRAMMING)
	 solver := ortoolslp.NewSolver("LinearExample", ortoolslp.SolverCLP_LINEAR_PROGRAMMING)
	// solver := ortoolslp.NewSolver("LinearExample", ortoolslp.SolverGLOP_LINEAR_PROGRAMMING)
	// solver := ortoolslp.NewSolver("LinearExample", ortoolslp.SolverGLPK_LINEAR_PROGRAMMING)
	// solver := ortoolslp.NewSolver("LinearExample", ortoolslp.SolverSCIP_MIXED_INTEGER_PROGRAMMING)
	// solver := ortoolslp.NewSolver("LinearExample", ortoolslp.SolverCPLEX_LINEAR_PROGRAMMING)
	// solver := ortoolslp.NewSolver("LinearExample", ortoolslp.SolverGUROBI_LINEAR_PROGRAMMING)

	//Create the two variables and let them take on any value.
	x := solver.NumVar(0, ortoolslp.SolverInfinity(), "x")
	y := solver.NumVar(0, ortoolslp.SolverInfinity(), "y")

	//Objective function: Maximize 3x + 4y.
	objective := solver.Objective()
	objective.SetCoefficient(x, 3)
	objective.SetCoefficient(y, 4)
	objective.SetMaximization()

	//Constraint 0: x + 2y <= 14.
	constraint0 := solver.Constraint(-ortoolslp.SolverInfinity(), float64(14))
	constraint0.SetCoefficient(x, 1)
	constraint0.SetCoefficient(y, 2)

	//Constraint 1: 3x - y >= 0.
	constraint1 := solver.Constraint(float64(0), ortoolslp.SolverInfinity())
	constraint1.SetCoefficient(x, 3)
	constraint1.SetCoefficient(y, -1)

	//Constraint 2: x - y <= 2.
	constraint2 := solver.Constraint(-ortoolslp.SolverInfinity(), float64(2))
	constraint2.SetCoefficient(x, 1)
	constraint2.SetCoefficient(y, -1)

	fmt.Println("Number of variables =", solver.NumVariables())
	fmt.Println("Number of constraints =", solver.NumConstraints())

	//Solve the system.
	status := solver.Solve()
	//Check that the problem has an optimal solution.
	if status != ortoolslp.SolverOPTIMAL {
		fmt.Println("The problem does not have an optimal solution!")
		os.Exit(1)
	}

	fmt.Println("Solution:")
	fmt.Println("x =", x.Solution_value())
	fmt.Println("y =", y.Solution_value())
	fmt.Println("Optimal objective value =", objective.Value())
	fmt.Println("")
	fmt.Println("Advanced usage:")
	fmt.Println("Problem solved in ", solver.Wall_time(), " milliseconds")
	fmt.Println("Problem solved in ", solver.Iterations(), " iterations")
	fmt.Println("x: reduced cost =", x.Reduced_cost())
	fmt.Println("y: reduced cost =", y.Reduced_cost())
	activities := solver.ComputeConstraintActivities()
	fmt.Println("constraint0: dual value =",
		constraint0.Dual_value())
	fmt.Println(" activities =",
		activities.Get(constraint0.Index()))
	fmt.Println("constraint1: dual value =",
		constraint1.Dual_value())
	fmt.Println(" activities =",
		activities.Get(constraint1.Index()))
	fmt.Println("constraint2: dual value =",
		constraint2.Dual_value())
	fmt.Println(" activities =",
		activities.Get(constraint2.Index()))
}

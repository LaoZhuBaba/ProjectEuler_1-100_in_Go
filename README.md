# ProjectEuler_1-100_in_Go
My solutions for Project Euler challenges from 1 - 100.  As a Go language learning exercise

Most of the solutions solve the problem from beginning to end.
In a few cases I partly solved the problem in my head so the starting point for the code may not be totally obvious.
In a at least one case I ran the code several times with different ranges of numbers before finding the solution.  As a result the inputs may seem arbitrary.

To compile or run a particular challenge you need to specify build tags.  So to run challenge 1 the command would be:

    go run -tags c1,challenge main.go

To compile challenge 86 to executable file challenge86:

    go build -tags c86,challenge -o challenge86

In most cases code which is common between challenge has been put in a separate "shared" package.ß
Application Outline for Port Scammer CLI Tool

I want to build a CLI application that can run continuously and listed for possible port scans and alert the user if a potential port scan is detected. The application should also provide a way to log the scans and display them in a user-friendly manner.

* the application should be built using Go
* it should be run on the command line
* it should listen for incoming TCP connections via a port allocation scheme
* if a port scan is detected a message should be displayed to the user
* the application should log the scans to a file
* the application should have a terminal user-interface built with the Bubbletea framework

Generate instructions in this order

1. Structure the application based on the existing code following go best practices
2. The command line options should be handled using the `cobra` package
3. Create a `cmd` directory to store the command line commands
4. Update the existing `main.go` file that initializes the Cobra command and starts the application
5. Create a `portscammer` package that contains the core logic of the application
    * This package should handle the TCP connections, detect port scans, log the scans, and provide a way to display them
    * The package should also include a function to start listening on a specified port
    * Use the `logrus` package for logging functionality
6. Create a `ui` package that contains the terminal user interface built with Bubbles
7. Create a `config` package to handle configuration settings for the application
8. Create a `models` package to define the data structures used in the application
9. Create a `utils` package for utility functions that can be reused across the application
10. Update the `README.md` file that provides an overview of the application, how to install it, and how to use it
11. The command should take one parameter for the port to listen on, defaulting to `8080` if not provided
12. Ensure that the application is well-documented with comments and has a clear structure for maintainability
13. Add unit tests for the core functionality of the application
14. Ensure that the application can be built and run using standard Go commands

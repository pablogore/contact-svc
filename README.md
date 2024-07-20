# Contact Duplicate Finder

## Description

Contact Duplicate Finder is an application that identifies duplicate contacts in a CSV file based on a weighted scoring system. The application compares various fields of the contacts and determines the similarity, assigning an accuracy (`High`, `Low`, or ignored) based on the score obtained.

## Architecture

The application is structured into several packages to maintain a modular and organized codebase:

- **main.go**: The entry point of the application. It reads contacts from a CSV file and finds duplicates.
- **model**: Defines data structures and interfaces.
    - `contact.go`: Contains the definition of the `Contact` structure and the `ContactComparer` interface.
- **handler**: Implements the business logic.
    - `contact_handler.go`: Contains the `FindDuplicates` function that identifies duplicate contacts.
- **utils**: Provides utility functions.
    - `file_csv.go`: Contains functions for reading CSV files.
[]
## Requirements

- Go 1.16 or higher

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/your-username/contact-svc.git
    cd contact-svc
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

## Usage

1. Make sure you have a `contacts.csv` file in the same directory as `main.go` with the following format:
    ```csv
    ID,FirstName,LastName,Email,ZipCode,Address
    1001,C,F,mollis.lectus.pede@outlook.net,449-6990,Tellus. Rd.
    1002,C,French,mollis.lectus.pede@outlook.net,39746,449-6990 Tellus. Rd.
    1003,Ciara,F,non.lacinia.at@zoho.ca,39746,1234 Main St
    ```

2. Run the application:
    ```sh
    go run main.go
    ```

## Testing

The project includes unit tests to ensure all key functionalities work correctly.

1. Run all tests:
    ```sh
    go test -v ./...
    ```

## Example Output

The output of the application will be similar to:

# QScrapper

QScrapper is a customizable web scraping tool designed to fetch and process data efficiently from various web sources. It leverages Go's powerful concurrency model and supports configurable proxies, delay between requests, and dynamic JSON path parsing for targeted data extraction.

## Project Structure

```
/qscrapper
    /cmd
        main.go          # Application entry point.
    /config
        config.go        # Manages loading of configuration settings.
    /scraper
        scraper.go       # Implements the scraping logic.
    /proxy
        proxy.go         # Handles proxy server rotation.
    /logger
        logger.go        # Provides logging functionality.
    /storage
        storage.go       # Manages data storage.
    /parser
        parser.go        # Parses JSON data based on configurable paths.
    config.json          # Stores configuration settings such as JSON paths, delays, and proxies.
    Makefile             # Simplifies build and run processes.
    README.md            # Documentation.
```

## Configuration

Edit `config.json` to specify your scraping parameters:

```json
{
    "Path": "entities.#.content.entity",
    "Delay": 5,
    "Proxies": []
}
```

- **Path**: JSON path for targeted data extraction.
- **Delay**: Time (in seconds) to wait between each request.
- **Proxies**: List of proxy servers to use for requests.

## Setup

1. Clone the repository:

```bash
git clone https://github.com/HritikR/QScrapper
cd qscrapper
```

2. Ensure Go is installed on your system and dependencies are set:

```bash
go mod tidy
```

## Building and Running

Use the provided Makefile for building and running the application:

- **Build the application**:

```bash
make build
```

- **Run the application**:

```bash
make run
```

- **Clean build artifacts**:

```bash
make clean
```

## Makefile Commands

- `build`: Compiles the application and places the binary in the `./build` directory.
- `run`: Builds (if necessary) and runs the compiled application.
- `clean`: Removes the `./build` directory and cleans up build artifacts.

## Usage

QScrapper is designed to be flexible, allowing you to specify various parameters directly from the command line to tailor the scraping process to your needs. Here's how to use the available flags:

- **`start`**: Specifies the starting page number for the scraping process. Defaults to `1` if not provided.
  
- **`end`**: Defines the ending page number for the scraping. Defaults to `1`, allowing for a single-page scrape if not overridden.
  
- **`url`**: The base URL to scrape, with a placeholder for the page number. This parameter is required and does not have a default value.
  
- **`out`**: Sets the path for the output file where the scraped data will be stored. Defaults to `output.json` if not specified.

### Running QScrapper

To run QScrapper with custom parameters, navigate to the project directory and execute the following command, adjusting the flags as needed:

```bash
go run cmd/main.go --start=1 --end=5 --url="http://example.com/pages?page={page}" --out="myData.json"
```

This example command will scrape pages 1 through 5 of `http://example.com/pages?page={page}`, replacing `{page}` with the actual page number, and save the results to `myData.json`.

### Running the Compiled Binary

```bash
./qscrapper --start=1 --end=5 --url="http://example.com/pages?page={page}" --out="myData.json"
```

## Customization

Adjust the `config.json` for different scraping needs. The application supports dynamic changes to the scraping path, request delays, and proxy configurations without code modifications.

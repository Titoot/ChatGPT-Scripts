~readme by ChatGPT~

This Go program reads a list of URLs from a file and checks the status of each site. It then writes the results to files based on the status code group (200, 300, or 400).

## Usage
To use this program, run the following command:

```bash
go run main.go <sites_file> <num_threads>
```
where `<sites_file>` is the path to the file containing the list of URLs and `<num_threads>` is the number of goroutines to use. The program will output the status code and content length for each site, and write the results to files named `200.txt`, `300.txt`, and `400.txt` for the corresponding status code groups.

## Example
Suppose you have a file named sites.txt containing the following URLs:

```
google.com
example.com
invalid.url
```
You can run the program with the following command:

```bash
go run main.go sites.txt 10
```

This will use 10 goroutines to check the status of each site in the file. The output will look something like this:

```
200 7424
200 7424
404 7424
```

The program will also create the files `200.txt`, `300.txt`, and `400.txt`, and write the results to the appropriate file. For example, the file `200.txt` will contain the following:

```
google.com [7424]
example.com [7424]
```


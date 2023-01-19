# waybackurls

This program uses the Wayback Machine's CDX API to fetch the URLs. It constructs the URL by appending the domain name entered by the user to the base URL of the API and makes a GET request to that URL. It then reads the response body and converts it to a string. The response body contains URLs separated by newline characters. The program splits the string into lines and prints out each line.


# Use
```
go run waybackurls.go
```

Then give your desired domain as input. For example, ```0xshahriar.xyz``` . You will find a file created with ```{your domain}.txt```. 

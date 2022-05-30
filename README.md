# Top pro-Russian sources in Slovakia in JSON

To get the [gerulata][1] data in JSON and enriched with URL, IP addresses,
Autonomous System and ICMP ping statitics:

```
go run . gerulata_top_pro_russian_sources.txt | \
jq -r '.[] | select(.threat!="low" and .type=="Web")'
```

I extracted text from PDF like this:

``` 
# had to manually fix page breaks and shorter IDs
pdftotext gerulata_top_pro_russian_sources.pdf
```

[1]: https://blog.gerulata.com/russian-propaganda-network-in-slovakia/

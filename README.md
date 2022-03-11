# Top pro-Russian sources in Slovakia

See https://blog.gerulata.com/russian-propaganda-network-in-slovakia/

* [x] Parse PDF into JSON
* [x] Enrich data with URL path, IP addresses and AS description
* [ ] Think how to use the data

# Parse PDF into JSON with additional data

```
pdftotext gerulata_top_pro_russian_sources.pdf # I had to manually fix some data

go run parse.go gerulata_top_pro_russian_sources.txt | \
jq -r '.[] | select(.threat!="low" and .type=="Web")'
```

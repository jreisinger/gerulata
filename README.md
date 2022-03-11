# Top pro-Russian sources in Slovakia in JSON

I took the data in PDF format from [gerulata][1] and converted it to text:

```
pdftotext gerulata_top_pro_russian_sources.pdf # had to manually fix some data
```

To get the data in JSON and enriched with URL path, IP addresses and AS
description:

```
go run parse.go gerulata_top_pro_russian_sources.txt | \
jq -r '.[] | select(.threat!="low" and .type=="Web")'
```

[1]: https://blog.gerulata.com/russian-propaganda-network-in-slovakia/
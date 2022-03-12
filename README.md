# Top pro-Russian sources in Slovakia in JSON

I took the data in PDF format from [gerulata][1] and converted it to text:

```
pdftotext gerulata_top_pro_russian_sources.pdf # had to manually fix some data
```

To get the data in JSON and enriched with URL path, IP addresses and AS
description:

```
go run . gerulata_top_pro_russian_sources.txt | \
jq -r '.[] | select(.threat!="low" and .type=="Web")'
```

```
{
  "id": 8799273843,
  "title": "Extraplus (extraplus.sk)",
  "type": "Web",
  "activity": "extreme",
  "impact": "n/a",
  "influence": "extreme",
  "threat": "extreme",
  "url": "extraplus.sk",
  "ip_addresses": [
    "195.181.248.13"
  ],
  "as": "WEBGLOBE-SK-AS",
  "ping": "0% packet loss, sent 5, recv 5, avg round-trip 13 ms"
}
```

[1]: https://blog.gerulata.com/russian-propaganda-network-in-slovakia/

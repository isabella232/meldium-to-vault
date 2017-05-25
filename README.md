# Meldium to Vault

1. Export meldium data to CSV - [How-to](http://support.meldium.com/knowledgebase/articles/656755-export-meldium-data-to-a-spreadsheet)

2. Rename downloaded file to `meldium.csv` and move it to root of this repo

3. Configure vault
```
export VAULT_ADDR=https://example.net:8200
export VAULT_TOKEN=***
```

4. Run `go run main.go`

This will try to create secrets in `meldium/*` backend, before running ensure it exists or change it in `main.go`

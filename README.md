# dndgen

Define general DnD rules and gameplay as an [ent](https://entgo.io/docs/getting-started) schema, then do fun stuff with it.

## usage

### dependencies

```
go install github.com/amonks/run/cmd/run@v1.0.0-beta.30
go get github.com/99designs/gqlgen@v0.17.30
go get github.com/hedwigz/entviz
```


Execute `run dev` to generate the schema, gqlgen files, and start the server.

## schema visualization screenshots
<img width="1400" alt="grid" src="assets/grid.png">
<hr>
<img width="1400" alt="loose" src="assets/loose.png">
<hr>
<img width="1400" alt="heir" src="assets/heirarchy.png">
<hr>
<img width="1400" alt="big" src="assets/big.png">
<hr>
<img width="1400" alt="zoom" src="assets/zoomhov.png">

## data

Data is seeded from the same JSON files that back the [5e-srd-api](https://github.com/5e-bits/5e-srd-api). (with some small adjustments here and there) 





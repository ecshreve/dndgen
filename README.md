# dndgen

Model Dungeons and Dragons rules and gameplay as an [ent](https://entgo.io/docs/getting-started) schema, then do fun stuff with it. Loose plans and ideas tracked in the [TODO](./TODO) file.
<hr>

## usage

devcontainer / local dev
```
- clone the repo
- open in devcontainer (recommended)
- or run locally

# These two are run during devcontainer setup after the container is created
> go install github.com/amonks/run/cmd/run@v1.0.0-beta.30
> go mod tidy

# Build and run the app
> run dndgen

# or

# Run the app in development mode
> run dev
```

then visit `localhost:8087` in a browser for the graphql playground, add `/viz` to the url to see the `entviz` visualization.

## screenshots
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

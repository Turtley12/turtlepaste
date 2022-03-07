#/bin/bash
rm turtlepaste.wasm
env GOOS=js GOARCH=wasm go build -o turtlepaste.wasm . 
echo Built!
cp turtlepaste.wasm docs/
cd docs/
echo Starting Server!
turtleweb

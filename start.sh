#/bin/bash
rm turtlepaste.wasm
env GOOS=js GOARCH=wasm go build -o turtlepaste.wasm . 
echo Built!
cp turtlepaste.wasm web/
cd web/

echo Starting Server!
turtleweb

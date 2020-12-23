rm -rf dist/
mkdir dist/

dotnet warp -r osx-x64 -o dist/powbot-launcher-osx
dotnet warp -r win10-x64 -o dist/powbot-launcher.exe
dotnet warp -r linux-x64 -o dist/powbot-launcher-linux

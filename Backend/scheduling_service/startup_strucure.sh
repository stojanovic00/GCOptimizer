mkdir api config core persistence proto util
touch main.go 

cd api 
mkdir app client handler middleware
cd ..

cd core 
mkdir domain repo service
cd ..

cd persistence
mkdir client repo
cd ..

echo "Project structure created"

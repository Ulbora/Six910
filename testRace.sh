cd prodsort
go test -race > raceTest.out
sleep 15
cd ..
cd handlers
go test -race > raceTest.out
sleep 15
cd ..
cd managers
go test -race > raceTest.out
cd ..


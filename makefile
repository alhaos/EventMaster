build:
	go build -o EventMaster.exe -buildvcs=false

run:
	go build -o EventMaster.exe -buildvcs=false
	"start EventMaster.exe"
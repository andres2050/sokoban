# Examples:
# ./run.sh profundidad < nivel2.txt
# ./run.sh anchura < nivel2.txt
# ./run.sh iterativa < nivel2.txt

go run main.go $1 <&0
#./main $1 <&0
go run main.go "hello" | cat -e
go run main.go "HELLO" | cat -e
go run main.go "HeLlo HuMaN" | cat -e
go run main.go "1Hello 2There" | cat -e
go run main.go "Hello\nThere" | cat -e
go run main.go "{hello & There #}" | cat -e
go run main.go "hello There 1 to 2!" | cat -e
go run main.go "MaD3IrA&LiSboN" | cat -e
go run main.go "1a\"#FdwHywR&/()=" | cat -e
go run main.go "{|}~" | cat -e
go run main.go "[\]^_ 'a" | cat -e
go run main.go "RGB" | cat -e
go run main.go ":;<=>?@" | cat -e
go run main.go "\!\" #$%&'()*+,-./" | cat -e
go run main.go "ABCDEFGHIJKLMNOPQRSTUVWXYZ" | cat -e
go run main.go "abcdefghijklmnopqrstuvwxyz" | cat -e